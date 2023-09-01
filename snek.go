package main

import (
	"snek/pkg/ring_array"
	"time"
)

type cellState uint8

const (
	emptyCell cellState = 0
	snakeCell cellState = 1
	foodCell  cellState = 2
)

const (
	gridWidth  int = 20
	gridHeight int = 15
	gridSize       = gridWidth * gridHeight
	initSize   int = 5
)

var headStart = vector{gridWidth / 2, gridHeight / 2}

var cellChars = map[cellState]string{
	emptyCell: "  ",
	snakeCell: "[]",
}

const tickRate time.Duration = time.Second / 4

type vector struct {
	x int
	y int
}

func vectorAdd(a vector, b vector) vector {
	return vector{a.x + b.x, a.y + b.y}
}

func vectorEquals(a vector, b vector) bool {
	return (a.x == b.x) && (a.y == b.y)
}

type model struct {
	currentView view

	snake       ring_array.RingArray[vector]
	direction   vector
	inputBuffer ring_array.RingArray[vector]
}

func (m *model) grow() {
	tail, _ := m.snake.Tail()
	m.snake.PushBack(tail)
}

func outOfBounds(position vector) bool {
	if position.x > gridWidth-1 || position.x < 0 {
		return true
	}
	if position.y > gridHeight-1 || position.y < 0 {
		return true
	}
	return false
}

func tickUpdate(m model) (model, bool) {
	input, err := m.inputBuffer.PopBack()
	if err == nil {
		m.direction = input
	}
	head, _ := m.snake.Head()
	newHead := vectorAdd(head, m.direction)

	if outOfBounds(newHead) {
		m.currentView = gameOver
		return m, true
	}

	m.snake.PushFront(newHead)
	m.snake.PopBack()

	return m, false
}

func gridIndex(vector vector) int {
	return vector.y*gridWidth + vector.x
}
