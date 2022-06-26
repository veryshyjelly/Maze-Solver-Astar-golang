package bin

type node struct {
	state
	action
	reachCost int
	heuristic int
	parent    *node
}

func Node(State state, Parent *node, Action action, Heuristic, ReachCost int) *node {
	self := new(node)
	self.state = State
	self.parent = Parent
	self.action = Action
	self.heuristic = Heuristic
	self.reachCost = ReachCost
	return self
}

type state struct {
	x int
	y int
}

type action string
