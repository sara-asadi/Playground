package rockPaperScissor

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type RockPaperScissor struct{}

func (*RockPaperScissor) Play() {

	stdScanner := bufio.NewScanner(os.Stdin)

	var compScore, userScore, winningScore int = 0, 0, 3

	var userChoice int

	fmt.Println("GAME STARTED")

	for compScore < winningScore && userScore < winningScore {
		fmt.Println("Enter Your Choice")
		fmt.Println("1.Rock")
		fmt.Println("2.Paper")
		fmt.Println("3.Scissor")

		stdScanner.Scan()

		userChoice, _ = strconv.Atoi(stdScanner.Text())

		if userChoice != 1 && userChoice != 2 && userChoice != 3 {
			fmt.Println("You made the wrong choice!")
			continue
		}

		randomNumber := rand.Intn(2) + 1

		result, err := checkWinner(userChoice, randomNumber)

		if err == nil {
			if result == 0 {
				fmt.Println("Draw!")
			} else if result == 1 {
				userScore++
				fmt.Println("You Won!")
			} else if result == 2 {
				compScore++
				fmt.Println("You Lost!")
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	if userScore > compScore {
		fmt.Println("You Won the Game!")
	} else {
		fmt.Println("You Lost the Game!")
	}
}

func checkWinner(userChoice int, compChoice int) (int, error) {

	choices := map[int]string{1: "Rock", 2: "Paper", 3: "Scissor"}

	fmt.Println("Computer chose " + choices[compChoice])

	if userChoice == compChoice {
		return 0, nil
	}

	switch userChoice {
	case 1:
		if compChoice == 2 {
			return 1, nil
		} else {
			return -1, nil
		}
	case 2:
		if compChoice == 3 {
			return -1, nil
		} else {
			return 1, nil
		}
	case 3:
		if compChoice == 1 {
			return -1, nil
		} else {
			return 1, nil
		}
	}

	return -2, errors.New("wrong choice")
}
