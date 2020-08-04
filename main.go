package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type difficulty struct {
	difficulty string
	sleep      int
	answer     int
}

func randButton() string {
	r := rand.Intn(25)
	switch {
	case r == 0:
		return "q"
	case r == 1:
		return "w"
	case r == 2:
		return "e"
	case r == 3:
		return "r"
	case r == 4:
		return "t"
	case r == 5:
		return "y"
	case r == 6:
		return "u"
	case r == 7:
		return "i"
	case r == 8:
		return "o"
	case r == 9:
		return "p"
	case r == 10:
		return "a"
	case r == 11:
		return "s"
	case r == 12:
		return "d"
	case r == 13:
		return "f"
	case r == 14:
		return "g"
	case r == 15:
		return "h"
	case r == 16:
		return "j"
	case r == 17:
		return "k"
	case r == 18:
		return "l"
	case r == 19:
		return "z"
	case r == 20:
		return "x"
	case r == 21:
		return "c"
	case r == 22:
		return "v"
	case r == 23:
		return "b"
	case r == 24:
		return "n"
	case r == 25:
		return "m"
	}
	return "omg"
}

func fishFish(t int, ch chan bool) {
	ttt := time.Duration(t)
	time.Sleep(ttt)
	ch <- true
}

func game(selectedDif difficulty) {

	gscanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	fishCount := 0
	sl := time.Duration(selectedDif.sleep)

	fmt.Println("Fishing...")
	fmt.Println()
	for {
		ch := make(chan bool)
		button := randButton()
		time.Sleep(sl * time.Millisecond)
		fmt.Println("FISH!!!")
		go fishFish(selectedDif.answer, ch)
		fmt.Println()
		fmt.Println(button)
		fmt.Println()
		for gscanner.Scan() {
			gUserInput := gscanner.Text()
			if gUserInput == button && !<-ch {
				fishCount++
				fmt.Println("YOU GOT IT!!! YAY")
				fmt.Println()
				fmt.Print("Score:")
				fmt.Println(fishCount)
				break
			}
			fmt.Println("You failed to catch the fish(((")

		}

	}
}

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
					stateOfGame = 1
					break
				case userInput == "2":
					stateOfGame = 2
					break
				case userInput == "3":
					stateOfGame = 3
					break
				case userInput == "4":
					return
				default:
					fmt.Println("Invalid input. Try again)")
					fmt.Println()
					continue
				}
				break

			}
		case stateOfGame == 1:
			var d difficulty
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

					d.difficulty = "Easy"
					d.answer = 2500
					d.sleep = 5000
					game(d)
					break
				case userInput == "2":
					d.difficulty = "Normal"
					d.answer = 1000
					d.sleep = 10000
					game(d)
					break
				case userInput == "3":
					d.difficulty = "Hard"
					d.answer = 500
					d.sleep = 15000
					game(d)
					break
				case userInput == "4":
					return
				default:
					fmt.Println("Invalid input. Try again)")
					fmt.Println()
					continue
				}
				break
			}
		case stateOfGame == 2:
			fmt.Println("Tutorial")
			fmt.Println()
			fmt.Println("You are a fisher.")
			fmt.Println()
			fmt.Println("You want to eat, so you went to lake with fish.")
			fmt.Println()
			fmt.Println("To catch fish you need to press right buttons in a limited amout of time.")
			fmt.Println()
			fmt.Println("If you press all of them - you get your fish. If not...")
			fmt.Println()
			fmt.Println("Try again or change difficulty")
			fmt.Println()
			fmt.Println("1. Return to main menu")
			for scanner.Scan() {
				userInput := scanner.Text()
				switch {
				case userInput == "1":
					fmt.Println("1")
					stateOfGame = 0
					break
				default:
					fmt.Println("Invalid input. Try again)")
					fmt.Println()
					continue
				}
				break

			}
			break
		case stateOfGame == 3:
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
					fmt.Println("1")
					stateOfGame = 0
					break
				default:
					fmt.Println("Invalid input. Try again)")
					fmt.Println()
					continue
				}
				break

			}
			break
		}
	}
}
