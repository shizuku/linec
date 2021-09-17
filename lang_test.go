package main

import (
	"testing"
)

func TestIsSourceFile(t *testing.T) {
	if !isSourceFile("main.go") {
		t.Fail()
	}
	if !isSourceFile("main.rs") {
		t.Fail()
	}
	if !isSourceFile("main.c") {
		t.Fail()
	}
	if !isSourceFile("main.h") {
		t.Fail()
	}
	if !isSourceFile("main.cpp") {
		t.Fail()
	}
	if !isSourceFile("main.hpp") {
		t.Fail()
	}
	if !isSourceFile("main.cc") {
		t.Fail()
	}
	if !isSourceFile("main.hh") {
		t.Fail()
	}
	if !isSourceFile("main.cxx") {
		t.Fail()
	}
	if !isSourceFile("main.hxx") {
		t.Fail()
	}
	if !isSourceFile("main.js") {
		t.Fail()
	}
	if !isSourceFile("main.jsx") {
		t.Fail()
	}
	if !isSourceFile("main.ts") {
		t.Fail()
	}
	if !isSourceFile("main.tsx") {
		t.Fail()
	}
	if !isSourceFile("main.java") {
		t.Fail()
	}
}

func TestIsLangFile(t *testing.T) {
	if !isLangFile("main.go", "go") {
		t.Fail()
	}
	if !isLangFile("main.rs", "rust") {
		t.Fail()
	}
	if !isLangFile("main.c", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.h", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.cpp", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.hpp", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.cc", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.hh", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.cxx", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.hxx", "c/c++") {
		t.Fail()
	}
	if !isLangFile("main.js", "js") {
		t.Fail()
	}
	if !isLangFile("main.jsx", "js") {
		t.Fail()
	}
	if !isLangFile("main.ts", "ts") {
		t.Fail()
	}
	if !isLangFile("main.tsx", "ts") {
		t.Fail()
	}
	if !isLangFile("main.java", "java") {
		t.Fail()
	}
}
