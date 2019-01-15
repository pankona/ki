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
	onlyDir = flag.Bool("d", false, "specify to include only directories [default: false]")
	plane   = flag.Bool("p", false, "specify to enable plane rendering [default: false]")
	profile = flag.Bool("with-profile", false, "specify to enable profiling [default: false]")
)

func main() {
	flag.Parse()

	dirList := flag.Args()[0:]
	if len(dirList) == 0 {
		dirList = []string{"."}
	}

	k := ki.Ki{
		ConcurrentNum:   *con,
		IgnoreHiddenDir: !*all,
		IncludeDirOnly:  *onlyDir,
		IsPlane:         *plane,
	}

	if *profile {
		start := time.Now()
		defer func() {
			end := time.Now()
			fmt.Printf("%v msec elapsed in total\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
		}()

		ki.EnableProfile()
		fmt.Printf("Worker num: %d\n", k.ConcurrentNum)
	}

	for _, v := range dirList {
		rootdir, err := k.Traverse(v)
		if err != nil {
			fmt.Printf("failed to traverse: %v\n", err)
			os.Exit(1)
		}
		k.Render(rootdir)
	}
}
