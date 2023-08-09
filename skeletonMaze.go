package main

import "fmt"

// Struct for Maze
type maze struct {
	maze [][] string 
	xSize int
	ySize int
}

type icon struct {
	xPos int
	yPos int
	icon string
}

type player struct {
	image icon
}

type exit struct {
	image icon
}

type key struct {
	image icon
}


func main() {

	// 2D String Slice for Map 🎯
	mazeMap := [][] string {
		{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"0", " ", " ", " ", " ", " ", "0", " ", " ", " ", "0"},
		{"0", "0", "0", "0", "0", " ", "0", "0", "0", " ", "0"},
		{"0", " ", " ", " ", "0", " ", " ", " ", " ", " ", "0"},
		{"0", " ", "0", " ", "0", "0", "0", "0", "0", " ", "0"},
		{"0", " ", "0", " ", " ", " ", " ", " ", " ", " ", "0"},
		{"0", "0", "0", " ", "0", "0", "0", " ", "0", " ", "0"},
		{"0", " ", " ", " ", "0", " ", " ", " ", "0", " ", "0"},
		{"0", " ", "0", "0", "0", "0", "0", "0", "0", " ", "0"},
		{"0", " ", " ", " ", " ", " ", "0", " ", " ", " ", "0"},
		{"0", "0", "0", "0", "0", " ", "0", "0", "0", "0", "0"},
	}

	gameMaze := maze {
		maze: mazeMap,
		xSize: len(mazeMap),
		ySize: len(mazeMap),
	}

	exit := exit{
		image: icon{
			xPos: 5,
			yPos: len(mazeMap) - 1,
			icon: "E",
		},
	}

	key := key {
		image: icon { 
			xPos: 1,
			yPos: 1,
			icon: "K", 
		},
	}

	// Create the player
	player := newPlayer(5, 1)

	// Game loop
	running := true
	printables := [] *icon {&player.image, &key.image, &exit.image}
	for running {
		draw(&gameMaze, printables)
		move(&gameMaze, player)
	}
}

// This function takes in a position and returns a player instance
func newPlayer(xPos int, yPos int) *player {
	p := player {
		image: icon { 
			xPos: xPos,
			yPos: yPos,
			icon: "P",
		},
	}

	return &p
}

func draw(gameMazeTrue *maze, icons []*icon) {
	gameMaze := maze{
		maze:  make([][]string, len(gameMazeTrue.maze)),
		xSize: gameMazeTrue.xSize,
		ySize: gameMazeTrue.ySize,
	}

	for i := range gameMazeTrue.maze {
		gameMaze.maze[i] = make([]string, len(gameMazeTrue.maze[i]))
		copy(gameMaze.maze[i], gameMazeTrue.maze[i])
	}

	for _, s := range icons {
		gameMaze.maze[s.yPos][s.xPos] = s.icon
	}

	for row := range gameMaze.maze {
		for col := range gameMaze.maze[row] {
			fmt.Print(gameMaze.maze[row][col] + " ")
		}
		// New line at the end of every row
		fmt.Println()
	}
}


func move(gameMaze *maze, player *player) {


	var i string

	// Initialise get input
	getInput := true

	// getInput loop
	for getInput {

		// Start scanner
		fmt.Scan(&i)

		// Start at false
		getInput = false

		// Movement input checking
		if i == "w" {
			player.image.yPos -= 1
		} else if i == "a" {
			player.image.xPos -= 1
		} else if i == "s" {
			player.image.yPos += 1
		} else if i == "d" { 
			player.image.xPos += 1
		} else {
			// If we get no inputs, set it back to true so it asks again
			getInput = true
		}
	}
}
