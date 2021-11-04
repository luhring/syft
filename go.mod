module github.com/anchore/syft

go 1.16

require (
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d
	github.com/adrg/xdg v0.2.1
	github.com/alecthomas/jsonschema v0.0.0-20210301060011-54c507b6f074
	github.com/anchore/client-go v0.0.0-20210222170800-9c70f9b80bcf
	github.com/anchore/go-rpmdb v0.0.0-20210914181456-a9c52348da63
	github.com/anchore/go-testutils v0.0.0-20200925183923-d5f45b0d3c04
	github.com/anchore/go-version v1.2.2-0.20200701162849-18adb9c92b9b
	github.com/anchore/packageurl-go v0.0.0-20210922164639-b3fa992ebd29
	github.com/anchore/stereoscope v0.0.0-20211024152658-003132a67c10
	github.com/antihax/optional v1.0.0
	github.com/bmatcuk/doublestar/v2 v2.0.4
	github.com/docker/docker v20.10.8+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/facebookincubator/nvdtools v0.1.4
	github.com/go-test/deep v1.0.7
	github.com/google/go-cmp v0.5.6
	github.com/google/go-containerregistry v0.6.1-0.20210922191434-34b7f00d7a60
	github.com/google/uuid v1.3.0
	github.com/gookit/color v1.2.7
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-version v1.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.1.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/profile v1.5.0
	github.com/scylladb/go-set v1.0.2
	github.com/sergi/go-diff v1.2.0
	github.com/sigstore/cosign v1.2.2-0.20211101203451-ae36ba5bd7fd
	github.com/sirupsen/logrus v1.8.1
	github.com/spdx/tools-golang v0.1.0
	github.com/spf13/afero v1.6.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/vifraa/gopom v0.1.0
	github.com/wagoodman/go-partybus v0.0.0-20210627031916-db1f5573bbc5
	github.com/wagoodman/go-progress v0.0.0-20200731105512-1020f39e6240
	github.com/wagoodman/jotframe v0.0.0-20200730190914-3517092dd163
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/mod v0.5.1
	golang.org/x/net v0.0.0-20210928044308-7d9f5e0b762b
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/sigstore/cosign => ../cosign

replace github.com/anchore/stereoscope => ../stereoscope
