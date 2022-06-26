package bin

import (
	"fmt"
	"io/ioutil"
	"log"
)

type maze struct {
	height      int
	width       int
	walls       [][]bool
	start       state
	goal        state
	NumExplored int
	explored    map[state]bool
	solution
}

type solution struct {
	actions []action
	cells   []state
}

func Maze(filename string) *maze {
	self := new(maze)
	self.walls = make([][]bool, 0)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	contents := splitLines(file)
	self.height = len(contents)
	for _, v := range contents {
		if x := len(v); x > self.width {
			self.width = x
		}
	}

	for i := 0; i < self.height; i++ {
		row := make([]bool, self.width)
		for j, v := range contents[i] {
			if v == 'A' {
				self.start = state{x: i, y: j}
				row[j] = false
			} else if v == 'B' {
				self.goal = state{x: i, y: j}
				row[j] = false
			} else if v == ' ' {
				row[j] = false
			} else {
				row[j] = true
			}
		}
		self.walls = append(self.walls, row)
	}

	return self
}

func (self *maze) Print() {
	mySolution := self.solution.cells
	fmt.Println()
	for i, row := range self.walls {
		for j, col := range row {
			if col {
				fmt.Print("█")
			} else if self.start.x == i && self.start.y == j {
				fmt.Print("A")
			} else if self.goal.x == i && self.goal.y == j {
				fmt.Print("B")
			} else if onPath(mySolution, i, j) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type neighbor struct {
	action
	state
}

func (self *maze) neighbors(State state) []neighbor {
	row, col := State.x, State.y

	candidates := []neighbor{
		{"up", state{x: row - 1, y: col}},
		{"down", state{x: row + 1, y: col}},
		{"left", state{x: row, y: col - 1}},
		{"right", state{x: row, y: col + 1}},
	}

	result := make([]neighbor, 0)
	for _, v := range candidates {
		r, c := v.x, v.y
		if (0 <= r && r < self.height) && (0 <= c && c < self.width) && !self.walls[r][c] {
			result = append(result, v)
		}
	}

	return result
}

func (self *maze) Solve() {
	start := Node(self.start, nil, "", heuristic(self.start, self.goal), 0)
	frontier := ListFrontier()
	frontier.Add(start)

	self.explored = map[state]bool{}

	for {
		if frontier.Empty() {
			log.Fatalln("No solution")
		}

		curNode := frontier.Remove()
		self.NumExplored += 1

		if curNode.state == self.goal {
			actions, cells := make([]action, 0), make([]state, 0)
			for curNode.parent != nil {
				actions = append(actions, curNode.action)
				cells = append(cells, curNode.state)
				curNode = curNode.parent
			}
			reverse(actions)
			reverse(cells)
			self.solution = solution{actions: actions, cells: cells}
			return
		}

		self.explored[curNode.state] = true

		for _, v := range self.neighbors(curNode.state) {
			if _, ok := self.explored[v.state]; !(ok || frontier.ContainsState(v.state)) {
				frontier.Add(Node(v.state, curNode, v.action, heuristic(v.state, self.goal), curNode.reachCost+1))
			}
		}
	}
}
