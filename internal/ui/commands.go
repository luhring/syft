package ui

import (
	"fmt"

	"github.com/anchore/syft/syft/artifact"

	"github.com/anchore/syft/internal/config"

	"github.com/anchore/stereoscope"
	"github.com/anchore/stereoscope/pkg/event"
	"github.com/anchore/syft/internal"
	"github.com/anchore/syft/internal/version"
	"github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/wagoodman/go-partybus"
	goprogress "github.com/wagoodman/go-progress"
)

type AnalysisTask func(*sbom.Artifacts, *source.Source) ([]artifact.Relationship, error)

func doAnalysisCmd(userInput string, appConfig config.Application, analysisTasks []AnalysisTask) tea.Cmd {
	return func() tea.Msg {
		src, cleanup, err := source.New(userInput, appConfig.Registry.ToOptions(), appConfig.Exclusions)
		if err != nil {
			return errMsg{fmt.Errorf("failed to construct source from user input %q: %w", userInput, err)}
		}
		if cleanup != nil {
			defer cleanup()
		}
		defer stereoscope.Cleanup()

		sbom := sbom.SBOM{
			Source: src.Metadata,
			Descriptor: sbom.Descriptor{
				Name:          internal.ApplicationName,
				Version:       version.FromBuild().Version,
				Configuration: appConfig,
			},
		}

		for _, task := range analysisTasks {
			relationships, err := task(&sbom.Artifacts, src)
			if err != nil {
				return errMsg{err}
			}

			sbom.Relationships = append(sbom.Relationships, relationships...)
		}

		return analysisDoneMsg{
			Result{
				Source: *src,
				SBOM:   sbom,
			},
		}
	}
}

func waitForEventCmd(events <-chan partybus.Event) tea.Cmd {
	return func() tea.Msg {
		e := <-events

		switch e.Type {

		case event.FetchImage:
			if p, ok := e.Value.(goprogress.StagedProgressable); ok {
				return fetchingImageMsg{p}
			}

			return errMsg{fmt.Errorf("unexpected value type for FetchImage event: %#v", e)}

		case event.ReadImage:
			if p, ok := e.Value.(goprogress.Progressable); ok {
				return readingImageMsg{p}
			}

			return errMsg{fmt.Errorf("unexpected value type for ReadImage event: %#v", e)}

		}

		return nil
	}
}
