package main

import "fmt"

var version = "dev"

func main() {

	fmt.Printf("Version: %s\n", version)

	msg := sayHello("Ravi")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
