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
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zivlakmilos/rstax/models"
)

type setupModel struct {
	params models.SetupParams
}

func RunSetup(params models.SetupParams) {
	p := tea.NewProgram(newSetup(params))
	if _, err := p.Run(); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func newSetup(params models.SetupParams) setupModel {
	return setupModel{
		params: params,
	}
}

func (m setupModel) Init() tea.Cmd {
	return nil
}

func (m setupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m setupModel) View() string {
	return ""
}
