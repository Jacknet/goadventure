// GoAdventure - Windows \r\n version

package main

import (
	"bufio"
	"fmt"
	"os"
	//"math/rand"
)

func main() {

	// Game intro and reader is declared
	fmt.Println("-=GoAdventure!=-")
	fmt.Println("An adventure game made in Go by Joaquin Padilla!\n")
	reader := bufio.NewReader(os.Stdin)

	// Game asks for name.
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')

	// Easter egg condition.
	// Checks to see if you typed "Duke Nukem" and prints a special message.
	// Gives a generic introduction that prints out name or fallback otherwise.
	// Reference to the Duke Nukem video game series.
	if name == "Duke Nukem\r\n" {
		fmt.Println("It's time to kick ass and chew bubble gum. Unfortunately, you noticed that you're all out of gum...\n")
	} else if name == "\r\n" {
		fmt.Println("Alrighty, player.")
	} else {
		fmt.Println("Alrighty,", name)
	}

	//declare variable used to track objective
	obj := 0

	// Starting scenario is given along with the available commands
	fmt.Println("You are a gopher inside a room with the door locked. You see a key on the floor. What do you want to do?")
	fmt.Println("Options: [open door] [pick up key]")

	// First objective is to pick up the key
	for obj == 0 {
		option, _ := reader.ReadString('\n')
		if option == "pick up key\r\n" {
			obj++
			fmt.Println("You pick up the key from the floor with your paw.")
		} else if option == "open door\r\n" {
			fmt.Println("Your try to twist the knob with your paw, but it won't budge. The door is locked.")
		} else if obj == 0 {
			fmt.Println("You wonder what that means, but you can't seem to understand.")
		}
	}

	// Second objective is to open the door
	for obj == 1 {
		option, _ := reader.ReadString('\n')
		if option == "pick up key\r\n" {
			fmt.Println("You already have the key on your paw.")
		} else if option == "open door\r\n" {
			obj++
			fmt.Println("You unlock the door and twist the knob. The door opens and you exit the room.")
		} else if obj == 1 {
			fmt.Println("You wonder what that means, but you can't seem to understand.")
		}
	}

	// Third objective is to pick up the sword
	fmt.Println("You hear something approach you. You notice a sword as you look down.")
	fmt.Println("Options: [pick up sword]")
	for obj == 2 {
		option, _ := reader.ReadString('\n')
		if option == "pick up sword\r\n" {
			obj++
			fmt.Println("You pick up the sword with your paw. The sword feels solid.")
		} else if obj == 2 {
			fmt.Println("You wonder what that means, but you can't seem to understand.")
		}
	}


	fmt.Println("A bug approaches you! You must fight using the sword!")
	// TO DO: RNG mechanic fight with monster

	// Congratulatory message and end of game
	fmt.Println("You win! Press Enter/Return to exit.")
	reader.ReadString('\n')

}