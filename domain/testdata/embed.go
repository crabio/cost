package testdata

import (
	_ "embed"
)

//go:embed scheme_config.yml
var SchemeCfg []byte
