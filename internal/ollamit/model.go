package ollamit

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/ollamit/internal/git"
	"github.com/koki-develop/ollamit/internal/ollama"
)

const (
	prompt = `Based on the content of the git diff, generate a short and concise one-line commit message.
The commit message should clearly describe the specific changes without omitting any important details and exclude any unnecessary information.
Your entire response will be used directly as the commit message.`
)

type status int

const (
	_ status = iota
	statusGenerating
	statusGenerated
	statusCommitting
	statusSuccess
)

type Config struct {
	DryRun bool
	Model  string

	GitClient    *git.Client
	OllamaClient *ollama.Client
}

type Ollamit struct {
	config  *Config
	program *tea.Program

	// state
	diff           *git.Diff
	err            error
	status         status
	messageBuilder *strings.Builder
	width          int

	// component
	spinner spinner.Model
}

var _ tea.Model = (*Ollamit)(nil)

func New(cfg *Config) *Ollamit {
	m := &Ollamit{
		config: cfg,

		// state
		status:         statusGenerating,
		messageBuilder: new(strings.Builder),

		// component
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styleSpinner)),
	}

	p := tea.NewProgram(m)
	m.program = p

	return m
}

func (m *Ollamit) Start() error {
	if _, err := m.program.Run(); err != nil {
		return err
	}
	if m.err != nil {
		return m.err
	}

	return nil
}

func (m *Ollamit) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.diffCmd(),
	)

}

func (m *Ollamit) formatMsg(msg string) string {
	msg = strings.Trim(msg, `"`)
	msg = strings.TrimSuffix(msg, ".")
	return strings.TrimSpace(msg)
}
