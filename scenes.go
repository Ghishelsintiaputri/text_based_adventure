package main

import (
	"fmt"
	"strconv"
)

type Scene struct {
	Title   string
	Text    string
	Choices []Choice
}

type Choice struct {
	Text    string
	NextSceneID string
}

var scenes = map[string]*Scene{
	"start": {
		Title: "The Mysterious Forest",
		Text: `You wake up in a dense forest with no memory of how you got here.
The air is cool and fresh, and sunlight filters through the canopy above.
You see two paths ahead of you: one leading deeper into the dark woods,
and another following a small stream.`,
		Choices: []Choice{
			{"Follow the stream", "stream"},
			{"Go deeper into the woods", "woods"},
			{"Stay where you are", "stay"},
		},
	},
	
	"stream": {
		Title: "The Gentle Stream",
		Text: `You follow the stream as it winds through the forest. The water is crystal clear
and you can see small fish swimming. After walking for a while, you come across
an old wooden bridge and a cave entrance partially hidden by vines.`,
		Choices: []Choice{
			{"Cross the bridge", "bridge"},
			{"Enter the cave", "cave"},
			{"Go back to where you started", "start"},
		},
	},
	
	"woods": {
		Title: "Deep Woods",
		Text: `The forest grows darker as you venture deeper. Strange sounds echo around you,
and the trees seem to watch your every move. You find an ancient stone altar
with glowing runes, and notice a hidden path behind a large oak tree.`,
		Choices: []Choice{
			{"Examine the altar", "altar"},
			{"Take the hidden path", "hidden_path"},
			{"Retreat to safety", "start"},
		},
	},
	
	"stay": {
		Title: "A Wise Decision",
		Text: `You decide to stay put and gather your thoughts. As you sit quietly,
a friendly forest spirit appears before you. "You are wise to be cautious, traveler,"
it says. "I shall guide you safely home."`,
		Choices: []Choice{},
	},
	
	"bridge": {
		Title: "The Old Bridge",
		Text: `The bridge creaks ominously as you cross, but it holds your weight.
On the other side, you discover a small village where the villagers welcome you
and help you find your way back to civilization.`,
		Choices: []Choice{},
	},
	
	"cave": {
		Title: "The Dark Cave",
		Text: `Inside the cave, you find ancient writings on the walls and a treasure chest.
The writings tell of a lost civilization, and the chest contains a map that leads
you safely out of the forest.`,
		Choices: []Choice{},
	},
	
	"altar": {
		Title: "The Ancient Altar",
		Text: `As you touch the glowing runes, a portal opens before you. You step through
and find yourself in your own bed, wondering if it was all just a dream...`,
		Choices: []Choice{},
	},
	
	"hidden_path": {
		Title: "The Hidden Path",
		Text: `The path leads you to a beautiful clearing with a mystical fountain.
Drinking from the fountain grants you eternal wisdom, and you live out your days
as the forest's guardian.`,
		Choices: []Choice{},
	},
}

func GetStartingScene() *Scene {
	return scenes["start"]
}

func (s *Scene) Display() {
	fmt.Printf("\n=== %s ===\n", s.Title)
	fmt.Println(s.Text)
	
	if len(s.Choices) > 0 {
		fmt.Println("\nWhat will you do?")
		for i, choice := range s.Choices {
			fmt.Printf("%d. %s\n", i+1, choice.Text)
		}
	}
}

func (s *Scene) HandleChoice(choice string) *Scene {
	if len(s.Choices) == 0 {
		return nil
	}
	
	choiceNum, err := strconv.Atoi(choice)
	if err != nil || choiceNum < 1 || choiceNum > len(s.Choices) {
		fmt.Println("Invalid choice! Please enter a number from the options.")
		return s
	}
	
	selectedChoice := s.Choices[choiceNum-1]
	return scenes[selectedChoice.NextSceneID]
}