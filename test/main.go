package main

import (
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func main() {
	filename := "hyzx2.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if (len(content) >= 3 && content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF) || utf8.Valid(content) {
		fmt.Printf("%s is UTF-8 encoded with BOM\n", filename)
	} else {
		fmt.Printf("%s is not UTF-8 encoded\n", filename)
	}
}
