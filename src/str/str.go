package main

import "fmt"

func main() {
	text1 := "\"what's that?\", he said"
	text2 := `"what's that?", he said`
	text3 := "\u221A \U0000221a"
	fmt.Println(text1)
	fmt.Println(text2)
	fmt.Println(text3)
}