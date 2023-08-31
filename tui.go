package main

import (
	"snek/pkg/ring_array"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	m := model{
		snake:     ring_array.NewRingArray[vector](gridSize),
		direction: vector{0, 1},
	}
	m.snake.PushFront(headStart)
	m.grid[gridIndex(headStart)]++

	for i := 0; i < 5; i++ {
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

func handleInput() {

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "w":
			m.direction.x = 0
			m.direction.y = 1
		case "a":
			m.direction.x = -1
			m.direction.y = 0
		case "s":
			m.direction.x = 0
			m.direction.y = -1
		case "d":
			m.direction.x = 1
			m.direction.y = 0
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		return tickUpdate(m), tick
	}

	return m, nil
}

func (m model) View() string {
	return gridString(m)
}
