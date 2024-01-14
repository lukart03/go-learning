package main

import "fmt"

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

func main() {
	// text := "Here's my spammy page: http://hehefouls.netHAHAHA see you http://hehefouls.netHAHAHA dassdf http://hehefouls.netHAHAHA ."
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// fmt.Println(masker(text))
	var s Some = Some{"Test String"}
	s.Write() // Что выдаст программа и почему?
	fmt.Println(s)
}

type Some struct {
	test string
}

func (s Some) Write() {
	s.test = "Hello from Write method"
}
