package main

import "fmt"

type Printer struct {
    board *Board
}

func NewPrinter(b *Board) *Printer {
    return &Printer{b}
}

func (self *Printer) Reprint() {
    self.moveToHome()
    for _, row := range self.board.matrix {
        var s string
        for _, cell := range row {
            if cell.Alive {
                s += "\033[32mX\033[0m"
            } else {
                s += "-"
            }
        }
        self.printRow(s)
    }
}

func (self *Printer) ClearScreen() {
    fmt.Print("\033[2J")
}

func (self *Printer) moveToHome() {
    fmt.Print("\033[H")
}

func (self *Printer) printRow(s string) {
    fmt.Println("\033[0K" + s)
}
