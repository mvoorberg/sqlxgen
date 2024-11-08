package pg

import (
	i "github.com/mvoorberg/sqlxgen/internal/introspect"
	"github.com/mvoorberg/sqlxgen/internal/utils/fs"
)

type source struct {
	args IntrospectArgs
	fd   fs.FileDiscovery
}

func NewIntrospect(fd fs.FileDiscovery, args IntrospectArgs) i.Introspect {
	return source{fd: fd, args: args}
}

type IntrospectArgs struct {
	Schemas         []string
	TableInclusions []string
	TableExclusions []string

	QueryDirs       []string
	QueryInclusions []string
	QueryExclusions []string
}
