package main

import (
	"os"
	"testing"
)

func TestCountFileLine(t *testing.T) {
	if countFileLine("file_test.go") != 21 {
		t.Fail()
	}
}

func TestCountDirLine(t *testing.T) {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	n := countDirLine(d)
	t.Log(n)
}
