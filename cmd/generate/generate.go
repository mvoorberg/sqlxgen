package main

import (
	_ "embed"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mvoorberg/sqlxgen/internal/config"
	"github.com/mvoorberg/sqlxgen/internal/utils"
	"github.com/mvoorberg/sqlxgen/internal/utils/fs"
	"github.com/mvoorberg/sqlxgen/internal/utils/writer"
)

func main() {
	workingDir, err := os.Getwd()

	if err != nil {
		utils.ExitWithError(err)

		return
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
			SqlxAltPath:   "",
		},
	)

	if err != nil {
		utils.ExitWithError(err)

		return
	}

	sqlxGenCfg.InitLogger()

	err = sqlxGenCfg.Generate()

	if err != nil {
		utils.ExitWithError(err)

		return
	}

	println("all done!")
}
