package ollamit

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	styleSpinner   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	styleInfomark  = lipgloss.NewStyle().SetString("ℹ ").Foreground(lipgloss.Color("#FF00FF"))
	styleCheckmark = lipgloss.NewStyle().SetString("✔ ").Foreground(lipgloss.Color("#00FF00"))
	styleMessage   = lipgloss.NewStyle().Padding(1, 2).Bold(true).Foreground(lipgloss.Color("#FFFFFF"))
	styleQuestion  = lipgloss.NewStyle().SetString("? ").Foreground(lipgloss.Color("#FF00FF"))
)

func (m *Ollamit) View() string {
	s := new(strings.Builder)

	if m.diff != nil {
		fmt.Fprintf(s, "%sStaged files\n\n", styleInfomark.Render())
		for _, f := range m.diff.Files {
			fmt.Fprintf(s, "  %s\n", f)
		}
		fmt.Fprintln(s)
	}

	if m.status == statusGenerating {
		fmt.Fprintf(s, "%sGenerating...\n", m.spinner.View())
	} else {
		fmt.Fprintf(s, "%sGenerated!\n", styleCheckmark.Render())
	}

	msg := wordwrap.String(m.messageBuilder.String(), m.width-4)
	fmt.Fprintln(s, styleMessage.Render(m.formatMsg(msg)))

	switch m.status {
	case statusGenerated:
		fmt.Fprintf(s, "%s%s\n", styleQuestion.Render(), wordwrap.String("Press [enter] to commit, [r] to regenerate, or [q] to quit.", m.width))
	case statusCommitting:
		fmt.Fprintf(s, "%sCommitting...\n", m.spinner.View())
	case statusSuccess:
		fmt.Fprintf(s, "%sCommit successful!\n", styleCheckmark.Render())
	}

	return s.String()
}
