package main

import (
	"snek/pkg/ring_array"
	"time"
)

type cellState uint8

const (
	free   cellState = 0
	filled cellState = 1
)

const (
	gridWidth  int = 20
	gridHeight int = 15
	gridSize       = gridWidth * gridHeight
)

var headStart = vector{gridWidth / 2, gridHeight / 2}

var cellChars = map[cellState]string{
	free:   "  ",
	filled: "[]",
}

// var freeCells = make(map[vector]struct{})
// var filledCells = make(map[vector]struct{})

const tickRate time.Duration = time.Second / 4

type vector struct {
	x int
	y int
}

type model struct {
	grid      [gridWidth * gridHeight]int
	snake     ring_array.RingArray[vector]
	direction vector
}

func (m *model) grow() {
	tail := m.snake.Tail()
	m.snake.PushBack(tail)
	m.grid[gridIndex(tail)]++
}

func tickUpdate(m model) model {
	newHead := vector{
		m.snake.Head().x + m.direction.x,
		m.snake.Head().y + m.direction.y,
	}

	m.snake.PushFront(newHead)
	m.grid[gridIndex(newHead)]++

	tail, _ := m.snake.PopBack()
	m.grid[gridIndex(tail)]--

	return m
}

func gridIndex(vector vector) int {
	return vector.y*gridWidth + vector.x
}

func gridString(m model) string {
	var s string
	for y := gridHeight - 1; y >= 0; y-- {
		for x := 0; x < gridWidth; x++ {
			snakeNodes := m.grid[gridIndex(vector{x, y})]
			if snakeNodes == 0 {
				s += cellChars[free]
			}
			if snakeNodes > 0 {
				s += cellChars[filled]
			}
		}
		s += "\n"
	}
	return s
}
