package task

import (
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/spinner"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	goprogress "github.com/wagoodman/go-progress"
)

const progressCheckWaitTime = 100 * time.Millisecond

var (
	checkMarkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#22cc00"))
	titleStyle     = lipgloss.NewStyle().Bold(true)
)

type Model struct {
	id              string
	progress        progress.Model
	progressable    goprogress.Progressable
	titleInProgress string
	titleComplete   string
	isComplete      bool
	spinner         spinner.Model
}

func New(titleInProgress, titleComplete string, p goprogress.Progressable, id string) (*Model, tea.Cmd) {
	s := spinner.New()
	s.Spinner = spinner.Monkey
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := &Model{
		id:              id,
		progress:        progress.New(progress.WithDefaultScaledGradient()),
		progressable:    p,
		titleInProgress: titleInProgress,
		titleComplete:   titleComplete,
		spinner:         s,
	}

	cmd := tea.Batch(
		checkProgressCmd(p, id),
		s.Tick,
	)

	return m, cmd
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)

		return m, cmd

	case TaskUpdateMsg:
		if m.id != msg.id {
			return m, nil
		}

		if msg.percent >= 1 {
			m.isComplete = true
		}

		cmd := m.SetPercent(msg.percent)

		return m, cmd

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd

	default:
		return m, nil
	}
}

func (m Model) View() string {
	if !m.isComplete {
		return m.spinner.View() + " " + m.title() + " " + m.progress.View() + "\n"
	}

	return checkMarkStyle.Render("✔  ") + m.title() + "\n"
}

func (m *Model) SetPercent(p float64) tea.Cmd {
	return tea.Batch(
		m.progress.SetPercent(p),
		checkProgressCmd(m.progressable, m.id),
	)
}

func (m Model) title() string {
	if m.isComplete {
		return titleStyle.Render(m.titleComplete)
	}

	return titleStyle.Render(m.titleInProgress)
}

type TaskUpdateMsg struct {
	id      string
	percent float64
}

func checkProgressCmd(p goprogress.Progressable, id string) tea.Cmd {
	return tea.Tick(progressCheckWaitTime, func(t time.Time) tea.Msg {
		current := float64(p.Current())
		total := float64(p.Size())
		percent := current / total

		return TaskUpdateMsg{
			id:      id,
			percent: percent,
		}
	})
}
