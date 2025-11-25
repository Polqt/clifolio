package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SpinnerComponent struct {
	Model spinner.Model
}

func NewSpinner() SpinnerComponent {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return SpinnerComponent{Model: s}
}

func (s SpinnerComponent) Init() tea.Cmd {
	return s.Model.Tick
}

func (s SpinnerComponent) Update(msg tea.Msg) (SpinnerComponent, tea.Cmd) {
	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return s, cmd
}

func (s SpinnerComponent) View() string {
	return s.Model.View()
}