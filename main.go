package main

import (
	"flag"
	"fmt"
	"time"
)

var speed = flag.Int("speed", 300, "speed to run in milliseconds")

// glider
// var aliveCells = []Cell{{x: 18, y: 18}, {x: 18, y: 19}, {x: 18, y: 20}, {x: 17, y: 20}, {x: 16, y: 19}}

// random
var aliveCells = []Cell{{x: 18, y: 18}, {x: 18, y: 19}, {x: 17, y: 18}, {x: 16, y: 18}, {x: 16, y: 19}, {x: 18, y: 20}, {x: 17, y: 18}}

type Cell struct {
	x, y int
}

func (c Cell) Alive() bool {
	// Any live cell with two or three neighbors survives.
	// Any dead cell with three live neighbors becomes a live cell.
	if hasCorrectNeighbours(aliveCells, c) {
		return true
	}
	// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	return false
}

func hasCorrectNeighbours(aliveCells []Cell, cell Cell) bool {
	neighbours := 0
	cellsToCheck := []Cell{
		Cell{x: cell.x - 1, y: cell.y - 1},
		Cell{x: cell.x - 1, y: cell.y},
		Cell{x: cell.x - 1, y: cell.y + 1},
		Cell{x: cell.x + 1, y: cell.y - 1},
		Cell{x: cell.x + 1, y: cell.y},
		Cell{x: cell.x + 1, y: cell.y + 1},
		Cell{x: cell.x, y: cell.y - 1},
		Cell{x: cell.x, y: cell.y + 1},
	}
	for i := 0; i < len(cellsToCheck); i++ {
		if Contains(aliveCells, cellsToCheck[i]) {
			neighbours++
		}
	}
	if neighbours == 3 {
		return true
	}
	if Contains(aliveCells, cell) && neighbours == 2 {
		return true
	}
	return false
}

func main() {
	flag.Parse()
	start := true
	ticker := time.Tick(time.Duration(*speed) * time.Millisecond)
	for _ = range ticker {
		newAliveCells := []Cell{}
		fmt.Printf("\033[0;0H")
		for l := 0; l <= 50; l++ {
			line := ""
			for c := 0; c <= 50; c++ {
				cell := Cell{
					x: l,
					y: c,
				}
				cellPosition := Cell{x: cell.x, y: cell.y}
				if (!start && cell.Alive()) || (start && Contains(aliveCells, cellPosition)) {
					newAliveCells = append(newAliveCells, Cell{x: cell.x, y: cell.y})
					line += "O "
				} else {
					line += ". "
				}
			}
			fmt.Println(line)
		}
		aliveCells = newAliveCells
		start = false
	}
}

func Contains(a []Cell, x Cell) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
