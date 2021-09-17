package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func countFileLine(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		n++
	}
	n++
	if verbose {
		rel, err := filepath.Rel(wd, filePath)
		if err != nil {
			panic(err)
		}
		if humanRead {
			fmt.Printf("%s\t%s\n", toHuman(n), rel)
		} else {
			fmt.Printf("%d\t%s\n", n, rel)
		}
	}
	return n
}

func countDirLine(dirPath string) int {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	size := len(files)
	ch := make(chan int, size)
	for _, f := range files {
		go func(file os.FileInfo) {
			if file.IsDir() {
				ch <- countDirLine(path.Join(dirPath, file.Name()))
			} else {
				if isSourceFile(file.Name()) {
					ch <- countFileLine(path.Join(dirPath, file.Name()))
				} else {
					ch <- 0
				}
			}
		}(f)
	}
	n := 0
	for i := 0; i < size; i++ {
		n += <-ch
	}
	return n
}
