package ki

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

type Ki struct {
	ConcurrentNum   int
	IgnoreHiddenDir bool
	IncludeDirOnly  bool

	IsPlane bool

	finishCh chan struct{}

	mu         sync.RWMutex
	workingNum int
}

func (k *Ki) Traverse(path string) (*Entry, error) {
	if profile {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("%v msec elapsed to traverse\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
		}()
	}

	k.finishCh = make(chan struct{})

	rootpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	rootdir := &Entry{
		path:  rootpath,
		isDir: true,
	}

	k.workingNum++
	go func() {
		k.traverse(rootdir)

		k.mu.Lock()
		k.workingNum--
		k.mu.Unlock()

		k.mu.RLock()
		if k.workingNum == 0 {
			k.finishCh <- struct{}{}
		}
		k.mu.RUnlock()
	}()
	<-k.finishCh

	return rootdir, nil
}

func (k *Ki) traverse(e *Entry) {
	files, err := ioutil.ReadDir(e.path)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.entries = make([]*Entry, len(files))
	var ignored int

	for i, v := range files {
		if v.Name()[0] == '.' && k.IgnoreHiddenDir {
			// ignore hidden directory
			ignored++
			continue
		}
		if !v.IsDir() && k.IncludeDirOnly {
			ignored++
			continue
		}

		i = i - ignored

		fullpath, err := filepath.Abs(filepath.Join(e.path, v.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}

		e.entries[i] = &Entry{
			path:  fullpath,
			isDir: v.IsDir(),
		}

		if v.IsDir() {
			k.mu.RLock()
			if k.workingNum < k.ConcurrentNum {
				k.mu.RUnlock()

				k.mu.Lock()
				k.workingNum++
				k.mu.Unlock()

				go func(i int) {
					k.traverse(e.entries[i])

					k.mu.Lock()
					k.workingNum--
					if k.workingNum == 0 {
						k.finishCh <- struct{}{}
					}
					k.mu.Unlock()
				}(i)
			} else {
				k.mu.RUnlock()
				k.traverse(e.entries[i])
			}
		}
	}

	// trim entries according to number of ignored directory
	e.entries = e.entries[:len(e.entries)-ignored]
}
