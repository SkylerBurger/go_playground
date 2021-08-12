package main

import (
	"fmt"
)

type die interface {
	rollDie() int
}

type d4 struct {
	current_value int
}

func (die d4) rollDie() int {
	return 4
}

type d6 struct {
	current_value int
}

func (die d6) rollDie() int {
	return 6
}

func main() {
	myDie1 := d4{}
	myDie2 := d6{}
	myDice := []die{myDie1, myDie2}
	
	for _, die := range myDice {
		fmt.Println(die.rollDie())
	}
}