package main

import (
	"bytes"
	"fmt"
	//"os"
	"path/filepath"
	"time"
)

//var output = ioutil.Discard
//var output = os.Stdout
var output = &bytes.Buffer{}

func Render(e *entry) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("%v msec elapsed to render\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
	}()

	fmt.Fprintf(output, "./\n")
	for i, v := range e.entries {
		render(v, 0, []bool{}, i < len(e.entries)-1)
	}

	fmt.Println(output)
}

func render(e *entry, depth int, parentHasChild []bool, hasNext bool) {
	for i := 0; i < depth; i++ {
		if parentHasChild[i] {
			fmt.Fprintf(output, "|   ")
		} else {
			fmt.Fprintf(output, "    ")
		}
	}

	if hasNext {
		fmt.Fprintf(output, "|-- ")
	} else {
		fmt.Fprintf(output, "`-- ")
	}

	if e.isDir {
		fmt.Fprintf(output, "%s/", filepath.Base(e.path))
	} else {
		fmt.Fprintf(output, "%s", filepath.Base(e.path))
	}

	fmt.Fprintf(output, "\n")

	hasChild := append(parentHasChild, hasNext)
	for i, v := range e.entries {
		render(v, depth+1, hasChild, i < len(e.entries)-1)
	}
}
