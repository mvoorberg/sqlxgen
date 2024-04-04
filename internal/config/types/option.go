package types

import (
	"fmt"
	"strings"
)

type Option struct {
	MysqlModelBanner        *string `json:"mysqlModelBanner" yaml:"mysqlModelBanner"`
	PostgresModelBanner     *string `json:"postgresModelBanner" yaml:"postgresModelBanner"`
	PostgresInt64JsonString *bool   `json:"postgresInt64JsonString,string" yaml:"postgresInt64JsonString"`
}

func (q *Option) String() string {
	if q == nil {
		return "Option{nil}"
	}

	content := strings.Join(
		[]string{
			fmt.Sprintf("mysqlModelBanner: %v", q.MysqlModelBanner),
			fmt.Sprintf("postgresModelBanner: %v", q.PostgresModelBanner),
			fmt.Sprintf("postgresInt64JsonString: %v", q.PostgresInt64JsonString),
		},
		", ",
	)

	return fmt.Sprintf("Option{%s}", content)
}

func (q *Option) Merge(other *Option) *Option {
	if other == nil {
		return q
	}

	if q == nil {
		return other
	}

	if other.PostgresInt64JsonString != nil {
		q.PostgresInt64JsonString = other.PostgresInt64JsonString
	}

	if other.MysqlModelBanner != nil {
		q.MysqlModelBanner = other.MysqlModelBanner
	}

	if other.PostgresModelBanner != nil {
		q.PostgresModelBanner = other.PostgresModelBanner
	}

	return q
}
