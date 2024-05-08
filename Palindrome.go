package main

import (
	"bufio"
	"fmt"
	"os"
)

type Palindrome struct {
}

func (*Palindrome) Play() {

	stdScanner := bufio.NewScanner(os.Stdin)
	gameOver := false
	score := 0

	for !gameOver {
		fmt.Println("Enter a word:")

		stdScanner.Scan()

		word := stdScanner.Text()

		result := CheckResult(word)

		if result {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("You've Failed")
			break
		}

		fmt.Println("Your Current Score is ", score)
	}

	fmt.Println("Your Total Score is ", score)

}

func CheckResult(word string) bool {
	wordRunes := []rune(word)
	var reversedWord []rune

	for i := len(word) - 1; i >= 0; i-- {
		reversedWord = append(reversedWord, wordRunes[i])
	}

	if string(reversedWord) == word {
		return true
	} else {
		return false
	}
}
