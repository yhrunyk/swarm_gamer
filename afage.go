package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Go Adventure Game!")
	fmt.Println("You wake up in a mysterious place. Do you explore or stay?")
	fmt.Print("Type 'explore' or 'stay': ")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(strings.ToLower(action))

	if action == "explore" {
		fmt.Println("You start exploring and find a cave. Do you enter or walk past?")
		fmt.Print("Type 'enter' or 'walk': ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice == "enter" {
			fmt.Println("Inside the cave, you find a treasure chest! But there's a sleeping dragon. Do you take the treasure or sneak away?")
			fmt.Print("Type 'take' or 'sneak': ")
			treasureChoice, _ := reader.ReadString('\n')
			treasureChoice = strings.TrimSpace(strings.ToLower(treasureChoice))

			if treasureChoice == "take" {
				if rand.Intn(2) == 0 {
					fmt.Println("You grab the treasure and escape! You win!")
				} else {
					fmt.Println("The dragon wakes up and chases you. Game over!")
				}
			} else if treasureChoice == "sneak" {
				fmt.Println("You quietly sneak away and live to tell the tale. You win!")
			} else {
				fmt.Println("Invalid choice. Game over.")
			}
		} else if choice == "walk" {
			fmt.Println("You continue walking and find a village where you are safe. You win!")
		} else {
			fmt.Println("Invalid choice. Game over.")
		}
	} else if action == "stay" {
		fmt.Println("You decide to stay and nothing interesting happens. Game over.")
	} else {
		fmt.Println("Invalid choice. Game over.")
	}
}
