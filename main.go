package main

import "fmt"

func main() {
	msg := sayHello("Ravi")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
