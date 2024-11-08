package config

import (
	"github.com/mvoorberg/sqlxgen/internal/config/types"
	"github.com/mvoorberg/sqlxgen/internal/utils"
)

func defaultMysqlConfig() *Config {
	cfg := &Config{
		Name:   utils.PointerTo("default"),
		Engine: utils.PointerTo("mysql"),
		Database: &types.Database{
			Host:     utils.PointerTo("localhost"),
			Port:     utils.PointerTo("3306"),
			Db:       utils.PointerTo("mysql"),
			User:     utils.PointerTo("root"),
			Password: utils.PointerTo(""),
			Url:      utils.PointerTo("mysql://root:@localhost:3306/mysql"),
			SslMode:  nil,
		},
		Source: &types.Source{
			Models: &types.Model{
				Schemas: []string{"public"},
				Include: []string{"^.+$"},
				Exclude: []string{},
			},
			Queries: &types.Query{
				Paths:   []string{},
				Include: []string{"^.+$"},
				Exclude: []string{},
			},
		},
		Options: &types.Option{
			MysqlModelBanner: utils.PointerTo("This is a MySql generated model. DO NOT EDIT."),
		},
		Gen: &types.Gen{
			Store: &types.GenPartial{
				Path: "gen/store",
			},
			Model: &types.GenPartial{
				Path: "gen/mysql",
			},
		},
	}

	return cfg
}
