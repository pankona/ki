package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	dirList := flag.Args()[0:]

	for _, v := range dirList {
		rootdir, err := traverse(v)
		if err != nil {
			fmt.Println(err)
		}
		Render(rootdir)
	}
}
