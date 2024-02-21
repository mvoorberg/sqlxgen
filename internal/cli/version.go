package cli

import (
	"github.com/spf13/cobra"
)

func versionCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print the sqlxgen version",
		Run: func(_ *cobra.Command, _ []string) {
			println("https://github.com/mvoorberg/sqlxgen")
			println(version)
		},
	}

	return cmd
}
