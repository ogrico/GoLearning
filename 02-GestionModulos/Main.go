package main

import "fmt"

func hello(name string) string {
	message := fmt.Sprintf("Hi %s", name)
	return message
}
