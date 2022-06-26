package bin

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func splitLines(data []byte) [][]byte {
	res := make([][]byte, 0)
	row := make([]byte, 0)
	for _, v := range data {
		if v == '\n' {
			res = append(res, row)
			row = make([]byte, 0)
		} else {
			row = append(row, v)
		}
	}
	return res
}

func onPath(path []state, x, y int) bool {
	for _, v := range path {
		if v.x == x && v.y == y {
			return true
		}
	}
	return false
}

func heuristic(s state, goal state) int {
	return abs(s.x-goal.x) + abs(s.y-goal.y)
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
