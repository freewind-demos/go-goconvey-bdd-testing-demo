package main

import "fmt"

func Greeting(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Greeting("convey"))
}
