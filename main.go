package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	game := NewGame()
	game.Start()
}

type Game struct {
	scanner *bufio.Scanner
	currentScene *Scene
	playerName string
}

func NewGame() *Game {
	return &Game{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (g *Game) Start() {
	fmt.Println("=== WELCOME TO THE FOREST ADVENTURE ===")
	fmt.Print("Enter your name: ")
	g.scanner.Scan()
	g.playerName = strings.TrimSpace(g.scanner.Text())
	
	fmt.Printf("\nHello, %s! Your adventure begins...\n\n", g.playerName)
	
	// Start with the first scene
	g.currentScene = GetStartingScene()
	g.playScene()
}

func (g *Game) playScene() {
	for g.currentScene != nil {
		g.currentScene.Display()
		
		if len(g.currentScene.Choices) == 0 {
			fmt.Println("\n=== GAME OVER ===")
			break
		}
		
		fmt.Print("\nChoose an option: ")
		g.scanner.Scan()
		choice := strings.TrimSpace(g.scanner.Text())
		
		nextScene := g.currentScene.HandleChoice(choice)
		g.currentScene = nextScene
		
		if g.currentScene == nil {
			fmt.Println("\n=== Thanks for playing! ===")
		}
	}
}