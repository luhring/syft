package cmd

import (
	"bytes"
	"fmt"

	"github.com/anchore/syft/internal/formats/syftjson"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/source"

	"github.com/spf13/cobra"
)

// attestCmd represents the attest command
var attestCmd = &cobra.Command{
	Use:   "attest",
	Short: "Generate and attest an SBOM for an image",
	RunE: func(cmd *cobra.Command, args []string) error {
		userInput := args[0]

		src, cleanup, err := source.New(userInput, appConfig.Registry.ToOptions())
		if err != nil {
			return fmt.Errorf("failed to determine image source: %w", err)
		}
		defer cleanup()

		scope := appConfig.Package.Cataloger.ScopeOpt
		catalog, d, err := syft.CatalogPackages(src, scope)
		if err != nil {
			return fmt.Errorf("failed to catalog input: %w", err)
		}

		format := syftjson.Format()

		var sbomBuffer *bytes.Buffer
		err = format.Encode(sbomBuffer, catalog, d, &src.Metadata, scope)
		if err != nil {
			return err
		}

		signer, err := cos.Signer(opts...)
		attester, err := cos.Attester(signer)
		attestation, err := attester.Attest(img, sbomBuffer, predicateType)

		// conditionally...
		attestation.Push()

		// conditionally...
		attestation.Sig.UploadToTLog()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(attestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
