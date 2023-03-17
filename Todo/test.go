package main

import "fmt"

type some struct {
	q string
}

func main() {
	var text []some

	text = append(text, some{"aaa"})
	fmt.Println(text)
	text = append(text, some{"bbb"})
	fmt.Println(text)
}
