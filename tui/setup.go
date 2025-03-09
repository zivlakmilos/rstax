/*
Copyright © 2025 Milos Zivlak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package tui

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zivlakmilos/rstax/models"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Save ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Save"))

	fields = []struct {
		name  string
		width int
	}{
		{
			name:  "Name",
			width: 64,
		},
		{
			name:  "Jmbg",
			width: 13,
		},
		{
			name:  "Phone",
			width: 32,
		},
		{
			name:  "Email",
			width: 64,
		},
		{
			name:  "Address",
			width: 64,
		},
		{
			name:  "Place",
			width: 64,
		},
		{
			name:  "Bank Account",
			width: 32,
		},
	}
)

type setupModel struct {
	params models.SetupParams

	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

func RunSetup(params models.SetupParams) {
	p := tea.NewProgram(newSetup(params))
	if _, err := p.Run(); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func newSetup(params models.SetupParams) setupModel {
	m := setupModel{
		params: params,
	}

	for i, field := range fields {
		t := textinput.New()

		t.Cursor.Style = cursorStyle
		t.CharLimit = field.width
		t.Placeholder = field.name

		switch i {
		case 0:
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

			t.SetValue(m.params.Name)
		case 1:
			t.SetValue(m.params.Jmbg)
		case 2:
			t.SetValue(m.params.Phone)
		case 3:
			t.SetValue(m.params.Email)
		case 4:
			t.SetValue(m.params.Address)
		case 5:
			t.SetValue(m.params.Address)
		case 6:
			t.SetValue(m.params.BankAccount)
		}

		m.inputs = append(m.inputs, t)
	}

	return m
}

func (m setupModel) Init() tea.Cmd {
	return nil
}

func (m setupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "enter", "down", "ctrl+n",
			"shift+tab", "up", "ctrl+p":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				// TODO: Handle save
				return m, tea.Quit
			}

			if s == "tab" || s == "enter" || s == "down" || s == "ctrl+n" {
				m.focusIndex++
				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				}
			} else {
				m.focusIndex--
				if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs)
				}
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m setupModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(fmt.Sprintf("%s:\n", fields[i].name))
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (m setupModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
