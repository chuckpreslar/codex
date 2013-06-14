package librarian

import (
  "librarian/nodes"
)

type SelectManager struct {
  ast       *nodes.SelectStatement
  reference *nodes.Reference
}

func (mgmt *SelectManager) Select(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = nodes.NewAttribute(projection.(string), mgmt.reference)
    case nodes.Attribute:
    default:
      // Error.
    }
    mgmt.ast.Select(projection.(nodes.Attribute))
  }
  return mgmt
}

func (mgmt *SelectManager) Where(comparison nodes.Comparison) *SelectManager {
  mgmt.ast.Where(comparison)
  return mgmt
}

func (mgmt *SelectManager) AST() *nodes.SelectStatement {
  return mgmt.ast
}

func (mgmt *SelectManager) ToSql() string {
  visitor, ok := VISITORS["base"]
  if ok {
    visitor.Accept(mgmt.ast)
  }
  return ""
}

func NewSelectManager(reference *nodes.Reference) *SelectManager {
  mgmt := new(SelectManager)
  mgmt.reference = reference
  mgmt.ast = nodes.NewSelectStatement(reference)
  return mgmt
}
