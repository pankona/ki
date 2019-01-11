package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

func ConcurrentTraverse(path string) (*entry, error) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("%v msec elapsed to ConcurrentTraverse\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
	}()
	return ctraverse(path)
}

func ctraverse(path string) (*entry, error) {
	var emap = sync.Map{}

	curpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(curpath)
	if err != nil {
		return nil, err
	}

	rootdir := &entry{
		path:    curpath,
		isDir:   true,
		entries: make([]*entry, len(files)),
	}
	emap.Store(curpath, rootdir)

	var wg sync.WaitGroup
	for i, v := range files {
		p, err := filepath.Abs(filepath.Join(curpath, v.Name()))
		if err != nil {
			return nil, err
		}

		e := &entry{
			path:  p,
			isDir: v.IsDir(),
		}

		parent, ok := emap.Load(filepath.Join(p, ".."))
		if !ok {
			return nil, fmt.Errorf("failed to register [%s]: [%s] is not registered yet", e.path, parent)
		}

		if v.IsDir() {
			emap.Store(p, e)
			wg.Add(1)
			go func(i int) {
				en, err := ctraverse(p)
				if err != nil {
					panic(err)
				}
				parent.(*entry).entries[i] = en
				wg.Done()
			}(i)
		} else {
			parent.(*entry).entries[i] = e
		}
	}
	wg.Wait()

	return rootdir, nil
}
