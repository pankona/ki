package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pankona/ki"
)

var (
	con     = flag.Int("c", 0, "specify concurrent num [default: 0]")
	all     = flag.Bool("a", false, "specify to include hidden directory [default: false]")
	plane   = flag.Bool("p", false, "specify to enable plane rendering [default: false]")
	profile = flag.Bool("with-profile", false, "specify to enable profiling [default: false]")
)

func main() {
	flag.Parse()

	dirList := flag.Args()[0:]
	if len(dirList) == 0 {
		dirList = []string{"."}
	}

	t := ki.Ki{
		ConcurrentNum:   *con,
		IgnoreHiddenDir: !*all,
	}

	if *profile {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("%v msec elapsed in total\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
		}()

		ki.EnableProfile()
	}

	for _, v := range dirList {
		rootdir, err := t.Traverse(v)
		if err != nil {
			fmt.Printf("failed to traverse: %v\n", err)
			os.Exit(1)
		}
		t.Render(rootdir)
	}
}
