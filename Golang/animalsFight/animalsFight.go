package animalsfight

import (
	"bufio"
	"fmt"
	"os"
)

type AnimalsFight struct{}

func (*AnimalsFight) Play() {

	stdScanner := bufio.NewScanner(os.Stdin)

	for {

		team1 := NewTeam()
		team2 := NewTeam()

		fmt.Println("Team1:")
		team1.Introduce()

		fmt.Println("Team2:")
		team2.Introduce()
		fmt.Println("Who Wins?")

		stdScanner.Scan()

		userChoice := stdScanner.Text()
		if userChoice == "x" {
			break
		}
		checkResult(*team1, *team2, userChoice)
	}
}

func checkResult(team1 Team, team2 Team, userChoice string) {

	if team1.GetPower() > team2.GetPower() {
		if userChoice == "1" {
			fmt.Println("You were correct! Team1 Won :)")
		} else {
			fmt.Println("You were wrong! Team2 Won :(")
		}
	} else if team1.GetPower() < team2.GetPower() {
		if userChoice == "2" {
			fmt.Println("You were correct! Team1 Won :)")
		} else {
			fmt.Println("You were wrong! Team2 Won :(")
		}
	} else {
		fmt.Println("It was a Draw! No one Won the Fight :|")
	}

	fmt.Println()
}
