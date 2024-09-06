package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	currentScene *Scene
	player       *Player
	scenes       map[string]*Scene
}

type Player struct {
	health    int
	inventory []string
}

type Scene struct {
	id          string
	description string
	choices     []Choice
}

type Choice struct {
	text    string
	next    string
	action  func(*Game)
}

func NewGame() *Game {
	game := &Game{
		player: &Player{
			health:    100,
			inventory: []string{},
		},
		scenes: make(map[string]*Scene),
	}
	
	game.initializeScenes()
	return game
}

func (g *Game) initializeScenes() {
	// Scene 1: Entrance
	g.scenes["entrance"] = &Scene{
		id: "entrance",
		description: "You stand at the temple entrance. Two paths diverge before you: one leading into a dark corridor to the LEFT, and another to a brightly lit hall to the RIGHT.",
		choices: []Choice{
			{text: "Go LEFT into the dark corridor", next: "dark_corridor"},
			{text: "Go RIGHT into the bright hall", next: "bright_hall"},
			{text: "Examine the entrance more closely", next: "entrance_examine"},
		},
	}
	
	// Scene 2: Dark Corridor
	g.scenes["dark_corridor"] = &Scene{
		id: "dark_corridor",
		description: "The air is cold and damp. You can barely see ahead. You hear faint scratching sounds from the darkness.",
		choices: []Choice{
			{text: "Continue forward cautiously", next: "spider_room"},
			{text: "Go back to the entrance", next: "entrance"},
			{text: "Feel the walls for clues", next: "corridor_secret"},
		},
	}
	
	// Scene 3: Bright Hall
	g.scenes["bright_hall"] = &Scene{
		id: "bright_hall",
		description: "The hall is illuminated by mysterious glowing crystals. Ancient murals cover the walls depicting forgotten rituals.",
		choices: []Choice{
			{text: "Examine the murals", next: "mural_study"},
			{text: "Search for hidden passages", next: "hidden_passage"},
			{text: "Return to entrance", next: "entrance"},
		},
	}
	
	// Scene 4: Spider Room
	g.scenes["spider_room"] = &Scene{
		id: "spider_room",
		description: "Giant spider webs cover the room. A massive spider drops from the ceiling!",
		choices: []Choice{
			{text: "Fight the spider", next: "spider_fight"},
			{text: "Try to sneak past", next: "sneak_past"},
			{text: "Retreat quickly", next: "dark_corridor"},
		},
	}
	
	// Scene 5: Spider Fight
	g.scenes["spider_fight"] = &Scene{
		id: "spider_fight",
		description: "You engage the giant spider in combat!",
		choices: []Choice{
			{text: "Swing your weapon", next: "spider_victory", action: func(g *Game) {
				fmt.Println("You defeat the spider and find a golden key!")
				g.player.inventory = append(g.player.inventory, "golden_key")
			}},
			{text: "Use a torch", next: "spider_retreat"},
		},
	}
	
	// Scene 6: Spider Victory
	g.scenes["spider_victory"] = &Scene{
		id: "spider_victory",
		description: "The spider lies defeated. Beyond it, you see a treasure chamber.",
		choices: []Choice{
			{text: "Enter the treasure chamber", next: "treasure_room"},
			{text: "Go back to the corridor", next: "dark_corridor"},
		},
	}
	
	// Scene 7: Treasure Room
	g.scenes["treasure_room"] = &Scene{
		id: "treasure_room",
		description: "The room glitters with gold and jewels! You have found the temple's treasure.",
		choices: []Choice{
			{text: "Take the treasure and escape", next: "victory"},
			{text: "Continue exploring", next: "dark_corridor"},
		},
	}
	
	// Scene 8: Mural Study
	g.scenes["mural_room"] = &Scene{
		id: "mural_room",
		description: "The murals reveal a secret: 'When the crystal glows blue, press the third stone from the left.'",
		choices: []Choice{
			{text: "Press the third stone", next: "secret_chamber"},
			{text: "Return to the hall", next: "bright_hall"},
		},
	}
	
	// Scene 9: Victory
	g.scenes["victory"] = &Scene{
		id: "victory",
		description: "CONGRATULATIONS! You have successfully navigated the temple and claimed the treasure!",
		choices: []Choice{
			{text: "Play again", next: "entrance"},
			{text: "Quit game", next: "quit"},
		},
	}
	
	g.currentScene = g.scenes["entrance"]
}

func (g *Game) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for g.currentScene != nil {
		g.displayScene()
		
		if len(g.currentScene.choices) == 0 {
			fmt.Println("The adventure ends here.")
			break
		}
		
		choice := g.getPlayerChoice(scanner)
		if choice == -1 {
			continue
		}
		
		g.executeChoice(choice)
	}
}

func (g *Game) displayScene() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println(g.currentScene.description)
	fmt.Println("\nWhat do you do?")
	
	for i, choice := range g.currentScene.choices {
		fmt.Printf("%d. %s\n", i+1, choice.text)
	}
}

func (g *Game) getPlayerChoice(scanner *bufio.Scanner) int {
	fmt.Print("\nEnter your choice (1-", len(g.currentScene.choices), "): ")
	
	if !scanner.Scan() {
		return -1
	}
	
	input := strings.TrimSpace(scanner.Text())
	var choice int
	_, err := fmt.Sscanf(input, "%d", &choice)
	
	if err != nil || choice < 1 || choice > len(g.currentScene.choices) {
		fmt.Println("Invalid choice. Please try again.")
		return -1
	}
	
	return choice - 1
}

func (g *Game) executeChoice(choiceIndex int) {
	choice := g.currentScene.choices[choiceIndex]
	
	// Execute any action associated with the choice
	if choice.action != nil {
		choice.action(g)
	}
	
	// Handle special cases
	if choice.next == "quit" {
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	}
	
	// Move to next scene
	if nextScene, exists := g.scenes[choice.next]; exists {
		g.currentScene = nextScene
	} else {
		fmt.Println("The path leads nowhere... returning to start.")
		g.currentScene = g.scenes["entrance"]
	}
}