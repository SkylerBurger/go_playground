package main

import (
	"fmt"
)

// Example of a function closure
func returnGreetingFunc (salutation string) (greetingFunc func (string)) {
	greetingFunc = func (name string) {
		fmt.Printf("%v, %v", salutation, name)
	}
	return
}

func main() {
	newFunc := returnGreetingFunc("Good morning")
	// newFunc has "Good Morning" enclosed within
	newFunc("Skyler")
	newFunc("Pimm")

	// If a function is returned, you can immediately invoke it and pass args
	returnGreetingFunc("Good evening")("Skyler")
}