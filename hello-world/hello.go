package main

import "fmt"

const englishEllowPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World" 
	}

	return englishEllowPrefix + name
}

func main() {
	fmt.Println(Hello("word"))
}
