package main

import "time"

type cellState uint8

const (
	free   cellState = 0
	filled cellState = 1
)

const (
	gridWidth  int = 40
	gridHeight int = 30
)

var cellChars = map[cellState]string{
	free:   "  ",
	filled: "[]",
}

var freeCells = make(map[vector]struct{})
var filledCells = make(map[vector]struct{})

const tickRate time.Duration = time.Second / 4

type vector struct {
	x int
	y int
}

type model struct {
	grid [gridWidth * gridHeight]cellState

	head      vector
	direction vector
}

func tickUpdate(m model) model {
	m.head.x += m.direction.x
	m.head.y += m.direction.y

	m.grid[gridIndex(m.head)] = filled

	return m
}

func gridIndex(vector vector) int {
	return vector.y*gridWidth + vector.x
}

func gridString(m model) string {
	var s string
	for y := gridHeight - 1; y >= 0; y-- {
		for x := 0; x < gridWidth; x++ {
			s += cellChars[m.grid[gridIndex(vector{x, y})]]
		}
		s += "\n"
	}
	return s
}
