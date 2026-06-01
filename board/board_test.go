package board_test

import (
	"errors"
	"testing"

	"github.com/BryanDGuy/go-sudoku-solver/board"
)

var validMatrix = [9][9]int{
	{5, 3, 0, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 0, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 0, 7, 9},
}

func TestNewBoard_Valid(t *testing.T) {
	m := validMatrix
	_, err := board.NewBoard(&m)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestNewBoard_InvalidRow(t *testing.T) {
	m := validMatrix
	m[0][0] = 3
	_, err := board.NewBoard(&m)
	if !errors.Is(err, board.ErrInvalidRow) {
		t.Fatalf("expected ErrInvalidRow, got %v", err)
	}
}

func TestNewBoard_InvalidColumn(t *testing.T) {
	m := [9][9]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	_, err := board.NewBoard(&m)
	if !errors.Is(err, board.ErrInvalidColumn) {
		t.Fatalf("expected ErrInvalidColumn, got %v", err)
	}
}

func TestNewBoard_InvalidNonet(t *testing.T) {
	m := [9][9]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
	}
	_, err := board.NewBoard(&m)
	if !errors.Is(err, board.ErrInvalidNonet) {
		t.Fatalf("expected ErrInvalidNonet, got %v", err)
	}
}

func TestGetMatrix_ReturnsCopy(t *testing.T) {
	m := validMatrix
	b, _ := board.NewBoard(&m)
	got := b.GetMatrix()
	got[0][0] = 99
	if b.GetValue(0, 0) == 99 {
		t.Fatal("GetMatrix should return a copy, not a reference")
	}
}
