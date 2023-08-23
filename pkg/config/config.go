package config

import (
	"os"
	"path/filepath"
)

var PORJECT_ROOT string

func init() {
	cliPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	PORJECT_ROOT = cliPath + "/../../"
}
