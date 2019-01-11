package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	flag.Parse()
	dirList := flag.Args()[0:]
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("%v msec elapsed in total\n", (end.Sub(start)).Nanoseconds()/int64(time.Millisecond))
	}()

	for _, v := range dirList {
		//rootdir, err := ConcurrentTraverse(v)
		rootdir, err := Traverse(v)
		if err != nil {
			fmt.Println(err)
		}
		Render(rootdir)
	}
}
