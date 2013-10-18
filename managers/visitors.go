// Package managers provides AST managers for the codex package.
package managers

import (
	"github.com/chuckpreslar/codex/visitors"
)

// VisitorFor returns a AST visitor for the adapter argument.
func VisitorFor(adapter interface{}) visitors.VisitorInterface {
	switch adapter {
	case "mysql":
		return &visitors.MySqlVisitor{&visitors.ToSqlVisitor{}}
	case "postgres":
		return &visitors.PostgresVisitor{&visitors.ToSqlVisitor{}, 0}
	default:
		return &visitors.ToSqlVisitor{}
	}
}
