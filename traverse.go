package main

import (
	//	"fmt"
	"os"
	"path/filepath"
)

func traverse(path string) (*entry, error) {
	rootpath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	rootdir := &entry{path: rootpath, isDir: true}

	emap := make(map[string]*entry)
	emap[rootpath] = rootdir

	err = filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
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
			emap[p] = e
		}

		parent, ok := emap[filepath.Join(p, "..")]
		if !ok {
			return fmt.Errorf("failed to register [%s]: [%s] is not registered yet", e.path, parent)
		}

		parent.entries = append(parent.entries, e)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return rootdir, nil
}
