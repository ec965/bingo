package entities

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var dialect goqu.DialectWrapper

func init() {
	dialect = goqu.Dialect("postgres")
}
