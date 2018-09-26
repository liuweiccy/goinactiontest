package main

import (
	"fmt"
	"strconv"
)

func main() {
	str5()
}

func str1() {
	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)

	y, err := strconv.ParseInt("71309", 10, 0)
	fmt.Printf("%8T %6v %v\n", y, y, err)

	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)
}

func str2() {
	text1 := "\"what's that?\", he said"
	text2 := `"what's that?", he said`
	text3 := "\u221A \U0000221a"
	fmt.Println(text1)
	fmt.Println(text2)
	fmt.Println(text3)
}

func str3() {
	z := 10000
	s := strconv.FormatBool(z > 100)
	fmt.Println(s)

	i, err := strconv.ParseInt("0xDEED", 0, 32)
	fmt.Println(i, err)

	j, err := strconv.ParseInt("0707", 0, 32)
	fmt.Println(j, err)

	k, err := strconv.ParseInt("10111010001", 2, 32)
	fmt.Println(k, err)
}

func str4() {
	i := 16769023
	fmt.Println(strconv.Itoa(i))
	fmt.Println(strconv.FormatInt(int64(i), 10))
	fmt.Println(strconv.FormatInt(int64(i), 2))
	fmt.Println(strconv.FormatInt(int64(i), 16))
}

func str5() {
	s := "Alice is g刘irl ,, hafas"
	quoted := strconv.Quote(s)
	fmt.Println(quoted)
	fmt.Println(strconv.Unquote(quoted))
}

func str6() {
}




