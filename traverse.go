package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

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
