package main

import (
	_ "embed"
	"encoding/json"
	"regexp"
)

var ext map[string]([]string)

//go:embed ext.json
var data []byte

func init() {
	json.Unmarshal(data, &ext)
}

func isSourceFile(fileName string) bool {
	if lang == "" {
		return isAllLangFile(fileName)
	} else {
		return isSomeLangFile(fileName, langs)
	}
}

func isAllLangFile(fileName string) bool {
	for _, v := range ext {
		for _, p := range v {
			match, err := regexp.MatchString(p, fileName)
			if err != nil {
				panic(err)
			}
			if match {
				return true
			}
		}
	}
	return false
}

func isSomeLangFile(fileName string, langs []string) bool {
	for _, lang := range langs {
		for _, p := range ext[lang] {
			match, err := regexp.MatchString(p, fileName)
			if err != nil {
				panic(err)
			}
			if match {
				return true
			}
		}
	}
	return false
}

func isLangFile(fileName, lang string) bool {
	for _, p := range ext[lang] {
		match, err := regexp.MatchString(p, fileName)
		if err != nil {
			panic(err)
		}
		if match {
			return true
		}
	}
	return false
}
