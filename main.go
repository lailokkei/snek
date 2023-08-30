package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
}

func initialModel() model {
	return model{
		1,
	}
}

func tickUpdate(m model) model {
	m.count++

	return m
}

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		return tickUpdate(m), tick

	}

	return m, nil
}

func (m model) View() string {
	s := strconv.Itoa(m.count)
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
