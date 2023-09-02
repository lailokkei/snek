package main

import (
	"fmt"
	"snek/pkg/ring_array"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type currentView struct {
}

const (
	ingame   view = 0
	gameOver view = 1
	menu     view = 2
)

type view int

func initialModel() model {
	m := model{
		currentView: ingame,
		snake:       ring_array.NewRingArray[vector](gridSize),
		direction:   vector{0, 1},
		inputBuffer: ring_array.NewRingArray[vector](5),
	}
	m.snake.PushFront(headStart)

	for i := 0; i < initSize-1; i++ {
		m.grow()
	}

	m.food = randomEmpty(m.snake)

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
			next, ok := inputMap[msg.String()]
			prev, err := m.inputBuffer.Head()
			if err != nil {
				prev = m.direction
			}
			if ok && !redundantDirection(next, prev) {
				m.inputBuffer.PushFront(next)
			}
		}

	case tickMsg:
		var died bool
		m, died = tickUpdate(m)
		if !died {
			return m, tick
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	switch m.currentView {
	case ingame:
		s = inGameView(m)
	case gameOver:
		s = gameOverView(m)
	}
	// s += fmt.Sprintf("%v\n", m.snake.Length)
	return s
}

func gameOverView(m model) string {
	var s string
	s += "Game Over!\n"
	s += fmt.Sprintf("Your length : %d\n", m.snake.Length)
	return s
}

func inGameView(m model) string {
	grid := make([]cellState, gridSize)
	var s string
	for _, snakeNode := range m.snake.Array() {
		grid[gridIndex(snakeNode)] = snakeCell
	}

	grid[gridIndex(m.food)] = foodCell

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
