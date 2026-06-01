package solver

import "github.com/BryanDGuy/go-sudoku-solver/board"

func Solve(b *board.Board) (*board.Board, bool) {
	work := b.GetMatrix()
	if !solve(&work) {
		return nil, false
	}
	result, err := board.NewBoard(&work)
	if err != nil {
		return nil, false
	}
	return result, true
}

func solve(m *[9][9]int) bool {
	j, k, found := nextCell(m)
	if !found {
		return true
	}
	for v, ok := range availableValues(m, j, k) {
		if !ok {
			continue
		}
		m[j][k] = v
		if solve(m) {
			return true
		}
		m[j][k] = 0
	}
	return false
}

// nextCell picks the empty cell with the fewest available values (MRV heuristic).
func nextCell(m *[9][9]int) (int, int, bool) {
	bestJ, bestK, bestCount := -1, -1, 10
	for j := range 9 {
		for k := range 9 {
			if m[j][k] != 0 {
				continue
			}
			count := countAvailable(m, j, k)
			if count < bestCount {
				bestJ, bestK, bestCount = j, k, count
			}
		}
	}
	return bestJ, bestK, bestJ != -1
}

// countAvailable returns how many values can legally be placed at (j, k).
func countAvailable(m *[9][9]int, j, k int) int {
	av := availableValues(m, j, k)
	n := 0
	for v := 1; v <= 9; v++ {
		if av[v] {
			n++
		}
	}
	return n
}

func availableValues(m *[9][9]int, j, k int) [10]bool {
	var used [10]bool
	for c := range 9 {
		used[m[j][c]] = true
	}
	for r := range 9 {
		used[m[r][k]] = true
	}
	nonetRow, nonetCol := (j/3)*3, (k/3)*3
	for r := range 3 {
		for c := range 3 {
			used[m[nonetRow+r][nonetCol+c]] = true
		}
	}
	var available [10]bool
	for v := 1; v <= 9; v++ {
		available[v] = !used[v]
	}
	return available
}
