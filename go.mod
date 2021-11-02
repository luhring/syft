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
	github.com/docker/docker v20.10.7+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/facebookincubator/nvdtools v0.1.4
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gabriel-vasile/mimetype v1.3.1 // indirect
	github.com/go-test/deep v1.0.7
	github.com/google/go-cmp v0.5.6
	github.com/google/go-containerregistry v0.6.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/gookit/color v1.2.7
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-version v1.2.0
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.1.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.15.0 // indirect
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/profile v1.5.0
	github.com/scylladb/go-set v1.0.2
	github.com/sergi/go-diff v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/assertions v1.2.0 // indirect
	github.com/spdx/tools-golang v0.1.0
	github.com/spf13/afero v1.6.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/vifraa/gopom v0.1.0
	github.com/wagoodman/go-partybus v0.0.0-20210627031916-db1f5573bbc5
	github.com/wagoodman/go-progress v0.0.0-20200731105512-1020f39e6240
	github.com/wagoodman/jotframe v0.0.0-20200730190914-3517092dd163
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/mod v0.5.0
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/sigstore/cosign => ../cosign
