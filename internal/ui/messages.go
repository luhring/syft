package ui

import goprogress "github.com/wagoodman/go-progress"

type readyMsg struct{}

type fetchingImageMsg struct {
	progressable goprogress.StagedProgressable
}

type fetchImageProgressMsg struct {
	percent float64
}

type readingImageMsg struct {
	progressable goprogress.Progressable
}

type readImageProgressMsg struct {
	percent float64
}

type analysisDoneMsg struct {
	result Result
}

type errMsg struct {
	error error
}
