package main

import (
	"Maze_Solver_Astar/bin"
	"fmt"
)

func main() {
	m := bin.Maze("mazefiles/maze2.txt")
	fmt.Println("Maze: ")
	m.Print()
	fmt.Println("Solving.... ")
	m.Solve()
	fmt.Println("Sates Explored: ", m.NumExplored)
	fmt.Println("Solution: ")
	m.Print()
}
