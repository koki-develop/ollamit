package ollamit

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleSpinner   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	styleCheckmark = lipgloss.NewStyle().SetString("✔ ").Foreground(lipgloss.Color("#00FF00"))
	styleMessage   = lipgloss.NewStyle().Padding(1, 2).Bold(true).Foreground(lipgloss.Color("#FFFFFF"))
)

func (m *Ollamit) View() string {
	s := new(strings.Builder)

	if m.status == statusGenerating {
		fmt.Fprintf(s, "%sGenerating...\n", m.spinner.View())
	} else {
		fmt.Fprintf(s, "%sGenerated!\n", styleCheckmark.Render())
	}

	fmt.Fprintln(s, styleMessage.Render(m.formatMsg(m.messageBuilder.String())))

	switch m.status {
	case statusGenerated:
		fmt.Fprintln(s, "Press [enter] to commit, [r] to regenerate, or [q] to quit.")
	case statusCommitting:
		fmt.Fprintf(s, "%sCommitting...\n", m.spinner.View())
	case statusSuccess:
		fmt.Fprintf(s, "%sCommit successful!\n", styleCheckmark.Render())
	}

	if m.quit {
		fmt.Fprintln(s, "Goodbye!")
	}

	return s.String()
}
