package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/pankona/ki"
)

var (
	con      = flag.Int("c", runtime.NumCPU(), "specify concurrent num")
	level    = flag.Int("l", math.MaxInt32, "specify limit of tree depth")
	all      = flag.Bool("a", false, "specify to include hidden directory (default: false)")
	onlyDir  = flag.Bool("d", false, "specify to include only directories (default: false)")
	onlyFile = flag.Bool("f", false, "specify to include only files (default: false)")
	plane    = flag.Bool("p", false, "specify to enable plane rendering (default: false)")
	profile  = flag.Bool("with-profile", false, "specify to enable profiling (default: false)")
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
		IncludeFileOnly: *onlyFile,
		Depth:           *level,
		IsPlane: func() bool {
			if *onlyFile {
				return true
			}
			return *plane
		}(),
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
		rootdir, err := k.Traverse(v)
		if err != nil {
			fmt.Printf("failed to traverse: %v\n", err)
			os.Exit(1)
		}
		k.Render(rootdir)
	}
}
