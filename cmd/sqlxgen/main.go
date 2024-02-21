package main

import (
	"github.com/mvoorberg/sqlxgen/internal/cli"
	"github.com/mvoorberg/sqlxgen/internal/utils"
)

func main() {
	cmd := cli.RootCmd(Version)

	err := cmd.Execute()

	if err == nil {
		return
	}

	utils.ExitWithError(err)
}

var Version = "v1.0.2"
