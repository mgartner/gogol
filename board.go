package main

type Board struct {
    matrix [][]*Cell
    Width int
    Height int
}

// NewBoard creates a new Board.
func NewBoard(x, y int) *Board {
    return &Board{newMatrix(x, y), x, y}
}

// newMatrix returns a blank matrix of cells.
func newMatrix(x, y int) [][]*Cell {
    matrix := make([][]*Cell, x)
    for i := 0; i < x; i++ {
        column := make([]*Cell, y)
        for j := 0; j < y; j++ {
            column[j] = NewCell()
        }
        matrix[i] = column
    }
    return matrix
}

// eachCell calls the function f on each cell of the board.
func (self *Board) eachCell(f func(int, int, *Cell)) {
    for i, row := range self.matrix {
        for j, cell := range row {
            f(i, j, cell)
        }
    }
}

// eachCell calls the function f on each cell of the board.
func (self *Board) eachNeighbor(x, y int, f func(*Cell)) {
    for i := x - 1; i <= x + 1; i++ {
        if i > -1 && i < self.Width {
            for j:= y - 1 ; j <= y + 1; j++ {
                if j > -1 && j < self.Height && !(i == x && j == y) {
                    f(self.matrix[i][j])
                }
            }
        }
    }
}

/*
 * Advances the board to the next state.
 * Based on the rules:
 *   1. Any live cell with fewer than two live neighbours dies, as if caused by under-population.
 *   2. Any live cell with two or three live neighbours lives on to the next generation.
 *   3. Any live cell with more than three live neighbours dies, as if by overcrowding.
 *   4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction
 */
func (self *Board) Advance() {

    // Calculate neighbors of each cell.
    for i, row := range self.matrix {
        for j, cell := range row {
            cell.Neighbors = self.neighborCount(i, j)
        }
    }

    // Kill, keep alive, or revive each cell based on neighbor count.
    for _, row := range self.matrix {
        for _, cell := range row {
            if cell.Alive {
                if cell.Neighbors != 2 && cell.Neighbors != 3 {
                    cell.Alive = false
                }
            } else {
                if cell.Neighbors == 3 {
                    cell.Alive = true
                }
            }
        }
    }
}

/*
 * Advances the board to the next state.
 * Based on the rules:
 *   1. Any live cell with fewer than two live neighbours dies, as if caused by under-population.
 *   2. Any live cell with two or three live neighbours lives on to the next generation.
 *   3. Any live cell with more than three live neighbours dies, as if by overcrowding.
 *   4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction
 */
func (self *Board) ConcurrentAdvance() {

    // Calculate neighbors of each cell.
    goCount := 0
    done := make(chan bool)
    for i := range self.matrix {
        goCount++
        go func(row int, done chan bool) {
            for j, cell := range self.matrix[row] {
                cell.Neighbors = self.neighborCount(row, j)
            }
            done <- true
        }(i, done)
    }

    for ; goCount > 0; goCount-- {
        <-done
    }

    // Kill, keep alive, or revive each cell based on neighbor count.
    for i := range self.matrix {
        goCount++
        go func(row int, done chan bool) {
            for _, cell := range self.matrix[row] {
                if cell.Alive {
                    if cell.Neighbors != 2 && cell.Neighbors != 3 {
                        cell.Alive = false
                    }
                } else {
                    if cell.Neighbors == 3 {
                        cell.Alive = true
                    }
                }
            }
            done <- true
        }(i, done)
    }

    for ; goCount > 0; goCount-- {
        <-done
    }
}

func (self *Board) neighborCount(x, y int) int {
    count := 0
    for i := x - 1; i <= x + 1; i++ {
        if i > -1 && i < self.Width {
            for j:= y - 1 ; j <= y + 1; j++ {
                if j > -1 && j < self.Height && !(i == x && j == y) && self.matrix[i][j].Alive {
                    count++
                }
            }
        }
    }
    return count
}
