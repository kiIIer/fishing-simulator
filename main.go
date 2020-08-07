package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	stateOfGame := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to fishingSimulator.")
	fmt.Println()
	fmt.Println("Type number and then 'Enter' to navugate")
	fmt.Println()
	for {
		switch {
		case stateOfGame == 0:
			mainMenuDisplay(scanner, &stateOfGame)
		case stateOfGame == 1:
			playGameDisplay(scanner, &stateOfGame)
		case stateOfGame == 2:
			tutorialDisplay(scanner, &stateOfGame)
		case stateOfGame == 3:
			supportDisplay(scanner, &stateOfGame)
		}
	}
}

func mainMenuDisplay(scanner *bufio.Scanner, stateOfGame *int) {
	fmt.Println("Main menu")
	fmt.Println()
	fmt.Println("1. Play game")
	fmt.Println()
	fmt.Println("2. Tutorial")
	fmt.Println()
	fmt.Println("3. Support creator")
	fmt.Println()
	fmt.Println("4. Exit")
	for scanner.Scan() {
		userInput := scanner.Text()
		switch {
		case userInput == "1":
			*stateOfGame = 1
			break
		case userInput == "2":
			*stateOfGame = 2
			break
		case userInput == "3":
			*stateOfGame = 3
			break
		case userInput == "4":
			return
		default:
			fmt.Println("Invalid input. Try again)")
			fmt.Println()
			continue
		}
		return
	}
}

func playGameDisplay(scanner *bufio.Scanner, stateOfGame *int) {
	var d game
	fmt.Println("Play game")
	fmt.Println()
	fmt.Println("Select difficulty(default is easy")
	fmt.Println()
	fmt.Println("1. Easy")
	fmt.Println()
	fmt.Println("2. Normal")
	fmt.Println()
	fmt.Println("3. Hard")
	fmt.Println()
	fmt.Println("4. Return to main menu")
	for scanner.Scan() {
		userInput := scanner.Text()
		switch {
		case userInput == "1":

			d.timeout = 2500 * time.Millisecond
			d.sleep = 10000
			d.startGame()
			break
		case userInput == "2":
			d.timeout = 1000 * time.Millisecond
			d.sleep = 7500
			d.startGame()
			break
		case userInput == "3":
			d.timeout = 500 * time.Millisecond
			d.sleep = 5000
			d.startGame()
			break
		case userInput == "4":
			*stateOfGame = 0
		default:
			fmt.Println("Invalid input. Try again)")
			fmt.Println()
			continue
		}
		return
	}
}

func tutorialDisplay(scanner *bufio.Scanner, stateOfGame *int) {
	fmt.Println("Tutorial")
	fmt.Println()
	fmt.Println("You are a fisher.")
	fmt.Println()
	fmt.Println("You want to eat, so you went to lake with fish.")
	fmt.Println()
	fmt.Println("To catch fish you need to press right button in a limited amout of time. Be ready for a challange!")
	fmt.Println()
	fmt.Println("If you press it - you get your fish. If not...")
	fmt.Println()
	fmt.Println("Try again or change difficulty")
	fmt.Println()
	fmt.Println("1. Return to main menu")
	for scanner.Scan() {
		userInput := scanner.Text()
		switch {
		case userInput == "1":
			*stateOfGame = 0
			break
		default:
			fmt.Println("Invalid input. Try again)")
			fmt.Println()
			continue
		}
		break

	}
	return
}

func supportDisplay(scanner *bufio.Scanner, stateOfGame *int) {
	fmt.Println("Support")
	fmt.Println()
	fmt.Println("If you appreciate my work you can say thank you Mike)")
	fmt.Println()
	fmt.Println("I will be happy")
	fmt.Println()
	fmt.Println("Check out my github - 'https://github.com/kiIIer'")
	fmt.Println()
	fmt.Println("1. Return to main menu")
	for scanner.Scan() {
		userInput := scanner.Text()
		switch {
		case userInput == "1":
			*stateOfGame = 0
			break
		default:
			fmt.Println("Invalid input. Try again)")
			fmt.Println()
			continue
		}
		break

	}
	return
}

type game struct {
	sleep   int
	timeout time.Duration
}

func (g game) startGame() {
	rand.Seed(time.Now().UnixNano())
	fishCount := 0
	sl := time.Duration(rand.Intn(g.sleep))

	fmt.Println("Fishing...")
	fmt.Println()
	for {

		chFishForFish := make(chan bool)
		char := randChar()

		time.Sleep(sl * time.Millisecond)
		fmt.Println("FISH!!!")
		fmt.Println()
		fmt.Println(char)
		fmt.Println()

		go fishForFish(char, chFishForFish)

		select {
		case <-chFishForFish:
			fishCount++
			fmt.Println("YOU GOT IT!!! YAY")
			fmt.Println()
			fmt.Print("Score:")
			fmt.Println(fishCount)
		case <-time.After(g.timeout):
			fmt.Println("You failed to catch the fish(((")

		}
	}
}

func randChar() string {
	s := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	r := rand.Intn(len(s))
	return s[r]
}

func fishForFish(char string, ch chan bool) {
	gscanner := bufio.NewScanner(os.Stdin)
	for gscanner.Scan() {
		gUserInput := gscanner.Text()
		if gUserInput == char {
			ch <- true
			break
		}
	}
}
