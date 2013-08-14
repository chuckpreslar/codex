// Package nodes provides nodes to use in codex AST's.
package nodes

import (
  "github.com/chuckpreslar/codex/sql"
)

// ConstraintNode is a Nary node.
type ConstraintNode struct {
  Column interface{}
  Kind   sql.Constraint
  Expr   interface{}
}

// ConstraintNode factory method.
func Constraint(column interface{}, kind sql.Constraint, expr interface{}) (constraint *ConstraintNode) {
  constraint = new(ConstraintNode)
  constraint.Column = column
  constraint.Kind = kind
  constraint.Expr = expr
  return
}
