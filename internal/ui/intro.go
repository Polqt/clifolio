package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"clifolio/internal/services"
)

type tickMsg struct{}
type goToMenuMsg struct{}

type introModel struct {
	fullRunes  		[]rune
	pos	       		int
	lines 	   		[]string
	done 	   		bool
	ascii      		[]string
	showASCII  		bool
	asciiIndex 		int
	asciiOpacity 	float64
}

func IntroModel() introModel {
	introData, err := services.LoadASCII("assets/intro.txt")
	fullText := ""
	if err == nil {
		fullText = string(introData)
	} else {
		fullText = `Press any key to continue...`
	}

	fullText = strings.ReplaceAll(fullText, "\r\n", "\n")

	data, err := services.LoadASCII("assets/ascii.txt")
	var ascii []string
	if err == nil {
		ascii = append(ascii, string(data), "\n")
	}
	
	return introModel {
		fullRunes: []rune(fullText),
		lines: []string{""},
		ascii: ascii,
		asciiOpacity: 0.0,
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
	switch msg.(type) {

	case tickMsg:
		if m.pos < len(m.fullRunes) {
			r := m.fullRunes[m.pos]
			m.pos++

			if r == '\n' {
				m.lines = append(m.lines, "")
			} else {
				if len(m.lines) == 0 {
					m.lines = append(m.lines, "")
				}
				last := m.lines[len(m.lines)-1]
				last += string(r)
				m.lines[len(m.lines)-1] = last
			}
			return m, tick()
		}

		if !m.done {
			m.done = true
			if len(m.ascii) > 0 {
				m.showASCII = true
				return m, tick()
			}
			return m, nil
		} 

		if m.showASCII && m.asciiOpacity < 1.0 {
			m.asciiOpacity += 0.1
			if m.asciiOpacity > 1.0 {
				m.asciiOpacity = 1.0
			}
			return m, tick()
		}

		return m, nil

	case tea.KeyMsg:
		if m.done {
			return m, func() tea.Msg { return goToMenuMsg{} }
		}
	}


	return m, nil
}

func (m introModel) View() string {
	s := ""

	for i := 0; i < len(m.lines); i++ {
		if i > 0 {
			s += "\n"
		} 
		s += m.lines[i]
	}

	if m.showASCII {
		for i := 0; i < m.asciiIndex && i < len(m.ascii); i++ {
			s += m.ascii[i] + "\n"
		}
	}

	if m.done {
		s += "\nPress any key to continue..."
	}
	
	return s
}
