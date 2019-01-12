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
	profile = flag.Bool("with-profile", false, "specify to enable profiling [default: false]")
)

func main() {
	flag.Parse()

	dirList := flag.Args()[0:]

	t := ki.New(*con)

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
		ki.Render(rootdir)
	}
}
