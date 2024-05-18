package main

import (
	"bufio"
	"fmt"
	"os"

	animalsfight "github.com/sara-asadi/Playground/animalsFight"
	"github.com/sara-asadi/Playground/palindrome"
	"github.com/sara-asadi/Playground/rockPaperScissor"
)

type Game interface {
	Play()
}

func main() {

	stdScanner := bufio.NewScanner(os.Stdin)
	var game Game

	for {

		fmt.Println("Enter a game:")
		fmt.Println("1.RockPaperScissor")
		fmt.Println("2.Palindrome")
		fmt.Println("3.AnimalsFight")

		stdScanner.Scan()
		command := stdScanner.Text()

		if command == "x" {
			break
		}

		if command == "1" {
			game = &rockPaperScissor.RockPaperScissor{}
		} else if command == "2" {
			game = &palindrome.Palindrome{}
		} else if command == "3" {
			game = &animalsfight.AnimalsFight{}
		}

		game.Play()
	}

}
