package main

import (
	"fmt"
	"path/filepath"
)

func Render(e *entry) {
	fmt.Printf("./\n")
	for i, v := range e.entries {
		render(v, 0, []bool{}, i < len(e.entries)-1)
	}
}

func render(e *entry, depth int, parentHasChild []bool, hasNext bool) {
	for i := 0; i < depth; i++ {
		if parentHasChild[i] {
			fmt.Printf("|   ")
		} else {
			fmt.Printf("    ")
		}
	}

	if hasNext {
		fmt.Printf("|-- ")
	} else {
		fmt.Printf("`-- ")
	}

	if e.isDir {
		fmt.Printf("%s/", filepath.Base(e.path))
	} else {
		fmt.Printf("%s", filepath.Base(e.path))
	}

	fmt.Printf("\n")

	hasChild := append(parentHasChild, hasNext)
	for i, v := range e.entries {
		render(v, depth+1, hasChild, i < len(e.entries)-1)
	}
}
