package main

import (
    "time"
    "fmt"
    "runtime"
)

type Game struct {
    board *Board
    printer *Printer
}

func newGame(x, y int) *Game {
    board := NewBoard(x, y)
    return &Game{board, NewPrinter(board)}
}

func (self *Game) seed(x, y int) {
    self.board.matrix[x][y].Alive = true
}

func (self *Game) setupToad() {
    self.seed(10, 10)
    self.seed(10, 11)
    self.seed(10, 12)
    self.seed(11, 9)
    self.seed(11, 10)
    self.seed(11, 11)
}

func main() {

    runtime.GOMAXPROCS(1)

    game := newGame(30, 100)
    game.setupToad()

    game.printer.ClearScreen()

    startTime := time.Now()

    iterations := 10000

    for i := 0; i < iterations; i++ {
        game.board.Advance()
        game.printer.Reprint()
    }

    endTime := time.Now()

    fmt.Println(endTime.Sub(startTime))
}
