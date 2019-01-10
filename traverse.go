package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func Traverse(path string) (*entry, error) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("%v msec elapsed to traverse\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
	}()
	return traverse(path)
}

func traverse(path string) (*entry, error) {
	rootpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	rootdir := &entry{path: rootpath, isDir: true}

	emap := sync.Map{}
	emap.Store(rootpath, rootdir)

	walkfn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path == rootpath {
			return nil
		}

		p, err := filepath.Abs(filepath.Join(path, "."))
		if err != nil {
			return err
		}

		e := &entry{path: p, isDir: info.IsDir()}

		if info.IsDir() {
			emap.Store(p, e)
		}

		parent, ok := emap.Load(filepath.Join(p, ".."))
		if !ok {
			return fmt.Errorf("failed to register [%s]: [%s] is not registered yet", e.path, parent)
		}
		parent.(*entry).entries = append(parent.(*entry).entries, e)

		return nil
	}

	err = filepath.Walk(rootpath, walkfn)
	if err != nil {
		return nil, err
	}

	return rootdir, nil
}
