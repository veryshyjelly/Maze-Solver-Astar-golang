package bin

import "log"

type listFrontier struct {
	data *node
	next *listFrontier
}

func ListFrontier() *listFrontier {
	return new(listFrontier)
}

func (self *listFrontier) Add(Node *node) {
	if self.Empty() || (self.next.data.heuristic+self.next.data.reachCost > Node.heuristic+Node.reachCost) {
		newFront := new(listFrontier)
		newFront.data, newFront.next = Node, self.next
		self.next = newFront
	} else {
		self.next.Add(Node)
	}
}

func (self *listFrontier) ContainsState(State state) bool {
	if self.Empty() {
		return false
	} else if self.next.data.state == State {
		return true
	}
	return self.next.ContainsState(State)
}

func (self *listFrontier) Empty() bool {
	return self.next == nil
}

func (self *listFrontier) Remove() *node {
	if self.Empty() {
		log.Fatalln("empty frontier")
	}
	nd := self.next
	self.next = self.next.next
	return nd.data
}

