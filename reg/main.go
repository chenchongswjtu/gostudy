package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	// 匹配 "metadata": null 的正则表达式
	re, err := regexp.Compile(`"metadata": \w*`)
	if err != nil {
		panic(err)
	}

	res := re.Find(readFile("test.json"))
	log.Println(string(res))
}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return data
}
