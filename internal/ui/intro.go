package ui

import (
	"io/ioutil"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}
type goToMenuMsg struct{}

type introModel struct {
	fullText   string
	current    string
	index 	   int
	done 	   bool
	ascii      []string
	showASCII  bool
	asciiIndex int
}

func IntroModel() introModel {
	data, err := ioutil.ReadFile("assets/ascii.txt")
	var ascii []string
	if err == nil {
		ascii = append(ascii, string(data), "\n")
	}
	return introModel {
		fullText: "Hi! I'm Janpol Hidalgo,",
		ascii: ascii,
	}
}

func (m introModel) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(35*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m introModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		_ = msg
		if !m.done {
			if m.index < len(m.fullText) {
				m.current += string(m.fullText[m.index])
				m.index++
				return m, tick()
			}
			m.done = true
			return m, tick()
		} else if m.showASCII && m.asciiIndex < len(m.ascii) {
			m.asciiIndex++
			return m, tick()
		}
		return m, nil
		

	case tea.KeyMsg:
		if m.done {
			return m, func() tea.Msg { return goToMenuMsg{} }
		}
	}

	if m.done && !m.showASCII {
		m.showASCII = true
		return m, tick()
	}

	return m, nil
}

func (m introModel) View() string {
	s := m.current + "\n\n"

	if m.showASCII {
		for i := 0; i < m.asciiIndex && i < len(m.ascii); i++ {
			s += m.ascii[i] + "\n"
		}
	}

	if m.done && m.asciiIndex >= len(m.ascii) {
		s += "\n\nPress any key to continue..."
	}
	
	return s
}
