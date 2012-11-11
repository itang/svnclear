package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	wd, _     = os.Getwd()
	targetDir = flag.String("dir", wd, "svn root dir")
	count     = 0
)

func main() {
  flag.Parse()
	filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path, info.Name())
		if info.IsDir() && info.Name() == ".svn" /*&& exists(path)*/ {
			fmt.Printf("\"%v\" is .svn dir, try to delete it...\n", path)
			os.RemoveAll(path)
			count++
			fmt.Printf("delete \"%v\"!\n", path)
			return filepath.SkipDir
		}
		return nil
	})
	fmt.Printf("delete %v .svn dir.\n", count)
}

// Exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
