# go-sudoku-solver

A sudoku solver written in Go using recursive backtracking with the minimum remaining values (MRV) heuristic.

## Usage

```go
matrix := [9][9]int{
    {5, 3, 0, 6, 7, 8, 9, 1, 2},
    // ...
}

b, err := board.NewBoard(&matrix)
if err != nil {
    log.Fatal(err)
}

solution, ok := solver.Solve(b)
if !ok {
    log.Fatal("no solution found")
}
```

Use `0` to mark unsolved cells.

## How it works

The solver uses recursive backtracking. At each step, it picks the empty cell with the fewest legal values (MRV), tries each candidate, and recurses. If no candidate works, it backtracks to the previous cell and tries the next value.

The board validates rows, columns, and nonets on construction.
