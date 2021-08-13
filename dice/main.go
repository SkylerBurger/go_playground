package main

import (
	"fmt"
	"math/rand"
	"time" // used to seed the deterministic rand.Intn()
)

// The die interface
type die interface {
	rollDie() int
	getValue() int
}

// helper function
func roller(max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	return randomizer.Intn(max) + 1
}

//==========
// The Dice!
//==========

type d4 struct {
	current_value int
}

func (die *d4) rollDie() int {
	value := roller(4)
	die.current_value = value
	return die.current_value
}

func (die *d4) getValue() int {
	return die.current_value
}

type d6 struct {
	current_value int
}

func (die *d6) rollDie() int {
	value := roller(6)
	die.current_value = value
	return die.current_value
}

func (die *d6) getValue() int {
	return die.current_value
}

type d8 struct {
	current_value int
}

func (die *d8) rollDie() int {
	value := roller(8)
	die.current_value = value
	return die.current_value
}

func (die *d8) getValue() int {
	return die.current_value
}

type d10 struct {
	current_value int
}

func (die *d10) rollDie() int {
	die.current_value = roller(10)
	return die.current_value
}

func (die *d10) getValue() int {
	return die.current_value
}

type d12 struct {
	current_value int
}

func (die *d12) rollDie() int {
	die.current_value = roller(12)
	return die.current_value
}

func (die *d12) getValue() int {
	return die.current_value
}

type d20 struct {
	current_value int
}

func (die *d20) rollDie() int {
	die.current_value = roller(20)
	return die.current_value
}

func (die *d20) getValue() int {
	return die.current_value
}

func main() {
	myDie1 := d4{}
	myDie2 := d6{}
	myDie3 := d8{}
	myDie4 := d10{}
	myDie5 := d12{}
	myDie6 := d20{}
	
	myDice := []die{&myDie1, &myDie2, &myDie3, &myDie4, &myDie5, &myDie6}
	
	for _, die := range myDice {
		die.rollDie()
		fmt.Println(die.getValue())
	}
}