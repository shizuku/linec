package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var verbose bool
var lang string
var humanRead bool

func init() {
	flag.BoolVar(&verbose, "v", false, "show line count of each file.")
	flag.StringVar(&lang, "lang", "", "language to scan, default to all.\nusing ':' split languages.\ne.g: -lang=c:cpp:js:ts")
	flag.BoolVar(&humanRead, "h", false, "show count in human-readable form.")
}

var wd string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	wd = dir
}

func main() {
	flag.Parse()
	parseLangs()

	ret := make(chan int, 1)
	go func() {
		ret <- countDirLine(wd)
	}()
	if humanRead {
		fmt.Printf("totally: %s\n", toHuman(<-ret))
	} else {
		fmt.Printf("totally: %d\n", <-ret)
	}
}

var langs []string

func parseLangs() {
	langs = strings.Split(lang, ":")
}
