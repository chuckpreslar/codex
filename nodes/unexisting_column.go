// Package nodes provides nodes to use in codex AST's.
package nodes

import (
  "github.com/chuckpreslar/codex/sql"
)

// UnexistingColumnNode is a Nary node.
type UnexistingColumnNode struct {
  Name *UnqualifiedColumnNode
  Type sql.Type
  Constraints []sql.Constraint
}

func UnexistingColumn(name interface{}, typ sql.Type, constraints ...sql.Constraint) (column *UnexistingColumnNode) {
  column = new(UnexistingColumnNode)
  column.Name = UnqualifiedColumn(name)
  column.Type = typ
  column.Constraints = constraints
  return
}
