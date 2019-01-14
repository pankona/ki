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

	limit chan struct{}
	wg    sync.WaitGroup
}

func New(concurrent int) *Ki {
	return &Ki{
		ConcurrentNum: concurrent,
	}
}

func (k *Ki) Traverse(path string) (*entry, error) {
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

	rootdir := &entry{
		path:  rootpath,
		isDir: true,
	}

	k.wg.Add(1)
	go func() {
		k.traverse(rootdir)
		k.wg.Done()
	}()
	k.wg.Wait()

	return rootdir, nil
}

func (k *Ki) traverse(e *entry) {
	files, err := ioutil.ReadDir(e.path)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.entries = make([]*entry, len(files))
	var ignored int

	for i, v := range files {
		if v.Name()[0] == '.' && k.IgnoreHiddenDir {
			// ignore hidden directory
			ignored++
			continue
		}

		i = i - ignored

		fullpath, err := filepath.Abs(filepath.Join(e.path, v.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}

		e.entries[i] = &entry{
			path:  fullpath,
			isDir: v.IsDir(),
		}

		if v.IsDir() {
			select {
			case k.limit <- struct{}{}:
				k.wg.Add(1)
				go func(i int) {
					k.traverse(e.entries[i])
					<-k.limit
					k.wg.Done()
				}(i)
			default:
				k.traverse(e.entries[i])
			}
		}
	}

	// trim entries according to number of ignored directory
	e.entries = e.entries[:len(e.entries)-ignored]
}
