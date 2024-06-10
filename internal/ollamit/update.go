package ollamit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/ollamit/internal/ollama"
)

type errorMsg struct{ err error }
type generatingMsg struct{ chunk string }
type regenerateMsg struct{}
type generatedMsg struct{}
type commitMsg struct{}

func (m *Ollamit) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errorMsg:
		m.err = msg.err
		return m, tea.Quit
	case generatingMsg:
		m.messageBuilder.WriteString(msg.chunk)
		return m, nil
	case regenerateMsg:
		return m, m.generateCmd()
	case generatedMsg:
		m.status = statusGenerated
		return m, nil
	case commitMsg:
		m.status = statusSuccess
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.quit = true
			return m, tea.Quit
		}

		if m.status == statusGenerated {
			if msg.Type == tea.KeyEnter {
				return m, m.commitCmd()
			}
			switch string(msg.Runes) {
			case "r":
				return m, m.regenerateCmd()
			case "q":
				m.quit = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m *Ollamit) regenerateCmd() tea.Cmd {
	return func() tea.Msg {
		m.status = statusGenerating
		m.messageBuilder.Reset()
		return regenerateMsg{}
	}
}

func (m *Ollamit) generateCmd() tea.Cmd {
	return func() tea.Msg {
		diff, err := m.git.DiffStaged()
		if err != nil {
			return errorMsg{err}
		}

		resp, err := m.ollama.Chat(&ollama.ChatInput{
			Model: m.config.Model,
			Messages: []ollama.ChatMessage{
				{Role: "system", Content: prompt},
				{Role: "user", Content: diff},
			},
		})
		if err != nil {
			return errorMsg{err}
		}

		for resp.Scan() {
			ch, err := resp.Chunk()
			if err != nil {
				return err
			}
			m.program.Send(generatingMsg{ch.Message.Content})
		}

		return generatedMsg{}
	}
}

func (m *Ollamit) commitCmd() tea.Cmd {
	return func() tea.Msg {
		if !m.config.DryRun {
			if err := m.git.Commit(m.formatMsg(m.messageBuilder.String())); err != nil {
				return errorMsg{err}
			}
		}
		return commitMsg{}
	}
}