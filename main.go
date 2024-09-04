package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== THE FORGOTTEN TEMPLE ADVENTURE ===")
	fmt.Println("You stand before an ancient temple, its stone walls covered in moss.")
	
	// Start the game
	game := NewGame()
	game.Start()
}