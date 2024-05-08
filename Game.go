package main

import (
	"bufio"
	"fmt"
	"os"
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

		stdScanner.Scan()
		command := stdScanner.Text()

		if command == "x" {
			break
		}

		if command == "1" {
			game = &RockPaperScissor{}
		} else if command == "2" {
			game = &Palindrome{}
		}

		game.Play()
	}

}
