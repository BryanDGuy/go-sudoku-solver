package board

import "errors"

type Board struct {
	matrix *[9][9]int

	byColumns *[9][9]int
	byNonets  *[9][9]int
}

var (
	ErrInvalidRow    = errors.New("invalid row")
	ErrInvalidColumn = errors.New("invalid column")
	ErrInvalidNonet  = errors.New("invalid nonet")
)

func NewBoard(m *[9][9]int) (*Board, error) {
	byColumns := convertColumnsToRows(m)
	byNonets := convertNonetsToRows(m)

	for _, row := range m {
		if valid := validateSlice(&row); !valid {
			return nil, ErrInvalidRow
		}
	}

	for _, column := range byColumns {
		if valid := validateSlice(&column); !valid {
			return nil, ErrInvalidColumn
		}
	}

	for _, nonet := range byNonets {
		if valid := validateSlice(&nonet); !valid {
			return nil, ErrInvalidNonet
		}
	}

	return &Board{
		matrix:    m,
		byColumns: byColumns,
		byNonets:  byNonets,
	}, nil
}

func (b *Board) GetMatrix() [9][9]int {
	return *b.matrix
}

func (b *Board) GetRow(i int) *[9]int {
	return &b.matrix[i]
}

func (b *Board) GetColumn(i int) *[9]int {
	return &b.byColumns[i]
}

func (b *Board) GetNonet(j, k int) *[9]int {
	return &b.byNonets[(j/3)*3+(k/3)]
}

func (b *Board) GetValue(j, k int) int {
	return b.matrix[j][k]
}

func convertColumnsToRows(m *[9][9]int) *[9][9]int {
	converted := [9][9]int{}
	for k := range 9 {
		column := [9]int{}
		for j := range 9 {
			column[j] = m[j][k]
		}
		converted[k] = column
	}
	return &converted
}

func convertNonetsToRows(m *[9][9]int) *[9][9]int {
	converted := [9][9]int{}
	n := 0
	for j := range 3 {
		for k := range 3 {
			nonet := [9]int{}
			i := 0
			for x := range 3 {
				for y := range 3 {
					nonet[i] = m[j*3+x][k*3+y]
					i++
				}
			}
			converted[n] = nonet
			n++
		}
	}
	return &converted
}

// All numbered 0-9 with no duplicates
func validateSlice(b *[9]int) bool {
	seen := make(map[int]struct{}, 9)
	for _, value := range b {
		if value < 0 || value > 9 {
			return false
		}

		// 0 indicates an unsolved position, allowed to have multiple
		if value != 0 {
			if _, ok := seen[value]; ok {
				return false
			}

			seen[value] = struct{}{}
		}
	}
	return true
}
