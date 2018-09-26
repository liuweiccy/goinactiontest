package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	TestSimplifyWitespace()
}

func log(str string) {
	fmt.Println(str)
}

func str1() {
	line := "ra ssd \u2028 va"
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line[j:])
	lastWord := line[j+size:]
	fmt.Println(firstWord, lastWord)

	var str string
	str = "VV"
	fmt.Println("Name: ", str)
}

func SimplifyWitespace(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}

	s = buffer.String()

	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}

	return s
}

func TestSimplifyWitespace() {
	s := "ssss   ssssss ss        sss ssssssss     "
	s = SimplifyWitespace(s)
	fmt.Println(s)
}
