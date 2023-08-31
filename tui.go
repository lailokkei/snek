package main

import (
	"snek/pkg/ring_array"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	m := model{
		snake:       ring_array.NewRingArray[vector](gridSize),
		direction:   vector{0, 1},
		inputBuffer: ring_array.NewRingArray[vector](5),
	}
	m.snake.PushFront(headStart)

	for i := 0; i < 20; i++ {
		m.grow()
	}
	return m
}

func tick() tea.Msg {
	time.Sleep(tickRate)
	return tickMsg{}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tick
}

var inputMap = map[string]vector{
	"w": {0, 1},
	"a": {-1, 0},
	"s": {0, -1},
	"d": {1, 0},
}

func redundantDirection(next vector, prev vector) bool {
	opposite := vectorEquals(vectorAdd(next, prev), vector{0, 0})
	return next == prev || opposite

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			val, ok := inputMap[msg.String()]
			if ok && !redundantDirection(val, m.inputBuffer.Head()) {
				m.inputBuffer.PushFront(val)
			}
		}

	case tickMsg:
		return tickUpdate(m), tick
	}

	return m, nil
}

func (m model) View() string {
	grid := make([]cellState, gridSize)
	var s string

	for _, snakeNode := range m.snake.Array() {
		grid[gridIndex(snakeNode)] = snakeCell
	}

	caps := "--"
	for i := 0; i < gridWidth; i++ {
		caps += "--"
	}
	caps += "\n"
	s += caps

	for y := gridHeight - 1; y >= 0; y-- {
		s += "|"
		for x := 0; x < gridWidth; x++ {
			cell := grid[gridIndex(vector{x, y})]
			s += cellChars[cell]
		}
		s += "|\n"
	}

	s += caps
	return s
}
