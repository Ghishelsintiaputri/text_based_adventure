# Interactive Text Adventure Game

A simple text-based adventure game written in Go where players explore an ancient temple with multiple paths and choices.

## How to Use

### Prerequisites
- Go 1.21 or later installed on your system

### Installation & Running

1. **Create a new directory for the game:**
   ```bash
   mkdir text-adventure
   cd text-adventure
   ```

2. **Create the game files:**
   Copy the provided code into three separate files:
   - `main.go`
   - `game.go`
   - `go.mod`

3. **Initialize and run the game:**
   ```bash
   go mod tidy
   go run .
   ```

### Game Instructions

- **Starting the Game:** The game begins automatically when you run the program
- **Making Choices:** You'll be presented with numbered options. Type the number of your choice and press Enter
- **Game Flow:** 
  - Start at the temple entrance
  - Choose between different paths (left/right corridors)
  - Encounter various challenges and puzzles
  - Multiple endings based on your choices
  - Collect items that may help you progress

### Features

- **Multiple Paths:** Different choices lead to different outcomes
- **Inventory System:** Collect items that affect gameplay
- **Combat Scenarios:** Make strategic choices in dangerous situations
- **Puzzle Solving:** Discover clues and solve mysteries
- **Replayability:** Try different paths to discover all endings

### Controls
- Type numbers (1, 2, 3, etc.) to select options
- Press Enter to confirm your choice
- The game provides clear prompts for all interactions

### Game Structure
- **Scenes:** Different locations in the temple
- **Choices:** Multiple options at each scene
- **Actions:** Some choices trigger special events or item collection
- **Progression:** Your choices determine which scenes you visit next

Enjoy your adventure through the Forgotten Temple!