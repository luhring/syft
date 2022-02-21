package ui

import (
	"github.com/anchore/syft/internal/config"
	"github.com/anchore/syft/internal/ui/components/task"
	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wagoodman/go-partybus"
)

type Model struct {
	// input
	userInput     string
	subscription  *partybus.Subscription
	appConfig     config.Application
	analysisTasks []AnalysisTask

	// view model
	fetchImage *task.Model
	readImage  *task.Model

	// output
	Result Result
}

type Result struct {
	Source source.Source
	SBOM   sbom.SBOM
	Error  error
}

func New(
	userInput string,
	subscription *partybus.Subscription,
	appConfig config.Application,
	analysisTasks []AnalysisTask,
) Model {
	return Model{
		userInput:     userInput,
		subscription:  subscription,
		appConfig:     appConfig,
		analysisTasks: analysisTasks,
	}
}

func (m Model) Init() tea.Cmd {
	return ready
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case readyMsg:
		return m, tea.Batch(
			waitForEventCmd(m.subscription.Events()),
			doAnalysisCmd(m.userInput, m.appConfig, m.analysisTasks),
		)

	case task.TaskUpdateMsg, progress.FrameMsg, spinner.TickMsg:
		// `task.TaskUpdateMsg` is sent when a task has just checked for an update to its progress
		// `progress.FrameMsg` is sent when a task's progress bar wants to animate itself
		// `spinner.TickMsg` is sent when it's time for a task's spinner to change its state
		return updateAllTasksWithMsg(msg, m)

	case fetchingImageMsg:
		var initialTaskCmd tea.Cmd
		m.fetchImage, initialTaskCmd = task.New("Loading image", "Loaded image", msg.progressable, "fetchImage")

		return m, tea.Batch(
			waitForEventCmd(m.subscription.Events()),
			initialTaskCmd,
		)

	case fetchImageProgressMsg:
		cmd := m.fetchImage.SetPercent(msg.percent)

		return m, cmd

	case readingImageMsg:
		var checkProgressCmd tea.Cmd
		m.readImage, checkProgressCmd = task.New("Parsing image", "Parsed image", msg.progressable, "readImage")

		return m, tea.Batch(
			waitForEventCmd(m.subscription.Events()),
			checkProgressCmd,
		)

	case readImageProgressMsg:
		cmd := m.readImage.SetPercent(msg.percent)

		return m, cmd

	case analysisDoneMsg:
		m.Result = msg.result

		return m, tea.Quit

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.

		m.Result.Error = msg.error

		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	// If we happen to get any other messages, don't do anything.
	return m, nil
}

func (m Model) View() string {
	output := ""

	if t := m.fetchImage; t != nil {
		output += t.View()
	}

	if t := m.readImage; t != nil {
		output += t.View()
	}

	output += "\n"

	return output
}

func ready() tea.Msg {
	return readyMsg{}
}

func updateAllTasksWithMsg(msg tea.Msg, currentModel Model) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	if t := currentModel.fetchImage; t != nil {
		model, cmd := t.Update(msg)
		cmds = append(cmds, cmd)
		updatedModel := model.(task.Model)
		currentModel.fetchImage = &updatedModel
	}

	if t := currentModel.readImage; t != nil {
		model, cmd := t.Update(msg)
		cmds = append(cmds, cmd)
		updatedModel := model.(task.Model)
		currentModel.readImage = &updatedModel
	}

	return currentModel, tea.Batch(cmds...)
}
