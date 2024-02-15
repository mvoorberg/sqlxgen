package cli

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mvoorberg/sqlxgen/internal/config"
	"github.com/mvoorberg/sqlxgen/internal/utils"
	"github.com/mvoorberg/sqlxgen/internal/utils/fs"
	"github.com/mvoorberg/sqlxgen/internal/utils/writer"
	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate code from sqlxgen config",
		Run: func(cmd *cobra.Command, _ []string) {
			sqlxGenAltPath, err := cmd.Flags().GetString("config")

			if err != nil {
				utils.ExitWithError(err)

				return
			}

			err = runGenerate(sqlxGenAltPath)

			if err == nil {
				return
			}

			utils.ExitWithError(err)
		},
	}

	cmd.
		PersistentFlags().
		String("config", "", "path to sqlxgen config")

	return cmd
}

func runGenerate(sqlxGenAltPath string) error {
	workingDir, err := os.Getwd()

	if err != nil {
		return err
	}

	fd := fs.NewFileDiscovery()

	wc := writer.NewFileWriter

	connect := sqlx.Connect

	sqlxGenCfg, err := config.NewSqlxGen(
		config.SqlxGenArgs{
			Connect:       connect,
			Fd:            fd,
			WriterCreator: wc,
			WorkingDir:    workingDir,
			SqlxAltPath:   sqlxGenAltPath,
		},
	)

	if err != nil {
		return err
	}

	sqlxGenCfg.InitLogger()

	return sqlxGenCfg.Generate()
}
