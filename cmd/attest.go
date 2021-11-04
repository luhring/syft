package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/sigstore/cosign/cmd/cosign/cli/sign"
	"github.com/sigstore/cosign/pkg/cosign/signature"

	"github.com/google/go-containerregistry/pkg/name"

	"github.com/anchore/syft/internal/formats/syftjson"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/source"

	"github.com/spf13/cobra"
)

const syftJSONPredicateTypeURI = "https://syft.dev/syft-json"

var privateKeyPath string

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

		sbomBuffer := new(bytes.Buffer)
		err = format.Encode(sbomBuffer, catalog, d, &src.Metadata, scope)
		if err != nil {
			return err
		}

		if src.Image == nil {
			return fmt.Errorf("user input did not specify an image: %q", userInput)
		}

		ref := src.Image.Ref()
		digest, err := name.NewDigest(ref.Name())
		if err != nil {
			return err
		}

		ctx := context.Background()

		keyOpts := sign.KeyOpts{
			KeyRef: privateKeyPath,
			PassFunc: func(_ bool) ([]byte, error) {
				// TODO: this should be config-driven
				pw, _ := os.LookupEnv("COSIGN_PASSWORD")
				return []byte(pw), nil
			},
			FulcioURL: "",
			RekorURL:  "",
		}

		att, err := signature.Attest(ctx, digest, syftJSONPredicateTypeURI, sbomBuffer, keyOpts)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", att.Envelope())

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

	attestCmd.Flags().StringVar(&privateKeyPath, "key", "", "path to cosign private key")
}
