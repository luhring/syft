package ruby

import (
	"bytes"
	"os"
	"testing"

	"github.com/anchore/syft/syft/pkg"
	"github.com/go-test/deep"
)

func TestParseGemspec(t *testing.T) {
	var expectedPkg = pkg.Package{
		Name:         "bundler",
		Version:      "2.1.4",
		Type:         pkg.GemPkg,
		Licenses:     []string{"MIT"},
		Language:     pkg.Ruby,
		MetadataType: pkg.GemMetadataType,
		Metadata: pkg.GemMetadata{
			Name:     "bundler",
			Version:  "2.1.4",
			Files:    []string{"exe/bundle", "exe/bundler"},
			Authors:  []string{"André Arko", "Samuel Giddins", "Colby Swandale", "Hiroshi Shibata", "David Rodríguez", "Grey Baker", "Stephanie Morillo", "Chris Morris", "James Wen", "Tim Moore", "André Medeiros", "Jessica Lynn Suttles", "Terence Lee", "Carl Lerche", "Yehuda Katz"},
			Licenses: []string{"MIT"},
			Homepage: "https://bundler.io",
		},
	}

	fixture, err := os.Open("test-fixtures/bundler.gemspec")
	if err != nil {
		t.Fatalf("failed to open fixture: %+v", err)
	}

	// TODO: no relationships are under test yet
	actual, _, err := parseGemSpecEntries(fixture.Name(), fixture)
	if err != nil {
		t.Fatalf("failed to parse gemspec: %+v", err)
	}

	if len(actual) != 1 {
		for _, a := range actual {
			t.Log("   ", a)
		}
		t.Fatalf("unexpected package count: %d!=1", len(actual))
	}

	for _, d := range deep.Equal(actual[0], &expectedPkg) {
		t.Errorf("diff: %+v", d)
	}
}

func FuzzParseGemSpecEntries(f *testing.F) {
	file, err := os.ReadFile("test-fixtures/bundler.gemspec")
	if err != nil {
		f.Fatal(err)
	}
	f.Add(file)

	f.Fuzz(func(t *testing.T, b []byte) {
		reader := bytes.NewBuffer(b)

		pkgs, _, err := parseGemSpecEntries("", reader)
		if err != nil && pkgs != nil {
			t.Errorf("%q, %v", pkgs, err)
		}
	})
}
