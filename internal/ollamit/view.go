package ollamit

import "strings"

func (m *Ollamit) View() string {
	s := new(strings.Builder)

	s.WriteString(m.messagePreview())

	if m.status == statusGenerated {
		s.WriteString(m.commands())
	}

	if m.status == statusCommitting {
		s.WriteString("Committing...\n")
	}

	if m.status == statusSuccess {
		s.WriteString("Commit successful!\n")
	}

	if m.quit {
		s.WriteString("Goodbye!\n")
	}

	return s.String()
}

func (m *Ollamit) messagePreview() string {
	s := new(strings.Builder)
	msg := m.formatMsg(m.messageBuilder.String())

	if m.status == statusGenerating {
		s.WriteString("Generating commit message...\n")
		s.WriteString(msg)
	}

	if m.status == statusGenerated || m.status == statusSuccess {
		s.WriteString("Generated message:\n")
		s.WriteString(msg)
	}

	s.WriteByte('\n')
	return s.String()
}

func (m *Ollamit) commands() string {
	return "Press [enter] to commit, [r] to regenerate, or [q] to quit.\n"
}
