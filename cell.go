package main

type Cell struct {
    Alive bool
    Neighbors int
}

func NewCell() *Cell {
    return &Cell{false, 0}
}
