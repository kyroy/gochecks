package main

import (
	"fmt"
	"github.com/kyroy/gochecks/pkg/gotest"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	gotests := path.Join("gotests")
	files, err := ioutil.ReadDir(gotests)
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 {
		fmtFiles := files
		files = nil
	Outer:
		for _, f := range fmtFiles {
			for _, f2 := range os.Args[1:] {
				if f.Name() == f2 {
					files = append(files, f)
					continue Outer
				}
			}
		}
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("Creating logs for", f.Name())
			testResult, _ := gotest.Run(path.Join(gotests, f.Name()))
			if err := ioutil.WriteFile(path.Join("testresults", f.Name()+".log"), testResult, os.ModePerm); err != nil {
				panic(err)
			}
		}
	}
}
