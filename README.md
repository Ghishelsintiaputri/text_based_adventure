# Interactive Text Adventure Game

A simple text-based adventure game written in Go where players make choices that lead to different story paths and endings.

## How to Run

1. **Prerequisites**: Make sure you have Go installed on your system (version 1.16 or higher).

2. **Setup**:
   ```bash
   # Clone or create the project directory
   mkdir text-adventure
   cd text-adventure
   
   # Place all the provided files in this directory
   ```

3. **Run the game**:
   ```bash
   go run main.go scenes.go
   ```

## How to Play

1. **Starting the Game**:
   - Run the program and enter your name when prompted
   - The adventure begins in a mysterious forest

2. **Making Choices**:
   - Read each scene description carefully
   - Choose an option by entering the corresponding number (1, 2, 3, etc.)
   - Press Enter after typing your choice

3. **Game Features**:
   - Multiple story paths and endings
   - Different locations to explore
   - Various outcomes based on your choices
   - Simple text-based interface

4. **Game Endings**:
   - The game ends when you reach a scene with no more choices
   - Each path leads to a different conclusion
   - You can replay to explore all possible endings

## File Structure

- `main.go`: Contains the main game loop and player interaction logic
- `scenes.go`: Defines all game scenes, choices, and story content
- `go.mod`: Go module configuration file

## Customization

You can easily modify the game by:
- Adding new scenes to the `scenes` map in `scenes.go`
- Creating new choices and connecting them to different scenes
- Changing existing story text and descriptions

Enjoy your adventure in the mysterious forest!