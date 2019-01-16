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
	Depth           int

	IsPlane bool

	limit chan struct{}
	wg    sync.WaitGroup
}

func (k *Ki) Traverse(path string) (*Entry, error) {
	if profile {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("%v msec elapsed to traverse\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
		}()
	}

	k.limit = make(chan struct{}, k.ConcurrentNum)

	rootpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	rootdir := &Entry{
		path:  rootpath,
		isDir: true,
	}

	var depth int

	k.wg.Add(1)
	go func() {
		k.traverse(rootdir, depth+1)
		k.wg.Done()
	}()
	k.wg.Wait()

	return rootdir, nil
}

func (k *Ki) traverse(e *Entry, depth int) {
	if depth > k.Depth {
		return
	}

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
			select {
			case k.limit <- struct{}{}:
				k.wg.Add(1)
				go func(i int) {
					k.traverse(e.entries[i], depth+1)
					<-k.limit
					k.wg.Done()
				}(i)
			default:
				k.traverse(e.entries[i], depth+1)
			}
		}
	}

	// trim entries according to number of ignored directories
	e.entries = e.entries[:len(e.entries)-ignored]
}
