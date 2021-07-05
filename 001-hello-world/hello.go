package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(names ...string) string {
	var name string

	if len(names) == 0 {
		name = "World"
	} else {
		name = names[0]
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello("Chris"))
}