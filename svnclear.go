package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	wd, _     = os.Getwd()
	targetDir = flag.String("dir", wd, "The root dir")
	count     = 0
)

func main() {
	flag.Parse()
	filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == ".svn" {
			fmt.Printf(">try to delete \"%v\"\n", path)
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				count++
			}
			return filepath.SkipDir
		}
		return nil
	})
	fmt.Printf("summary: delete %v .svn dir.\n", count)
}
