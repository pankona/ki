package ki

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"
)

var output = &bytes.Buffer{}

func (k *Ki) Render(e *Entry) {
	if profile {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("%v msec elapsed to render\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
		}()
	}

	if k.IsPlane {
		fmt.Fprintf(output, "%s", e.path)
	} else {
		fmt.Fprintf(output, "%s/", filepath.Base(e.path))
	}
	fmt.Fprintf(output, "\n")
	for i, v := range e.entries {
		k.render(v, 0, []bool{}, i < len(e.entries)-1)
	}

	fmt.Println(output)
}

func (k *Ki) render(e *Entry, depth int, parentHasChild []bool, hasNext bool) {
	if k.IsPlane {
		fmt.Fprintf(output, "%s", e.path)
	} else {
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
	}

	fmt.Fprintf(output, "\n")

	hasChild := append(parentHasChild, hasNext)
	for i, v := range e.entries {
		k.render(v, depth+1, hasChild, i < len(e.entries)-1)
	}
}
