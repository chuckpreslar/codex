package codex

import (
  "codex/tree/nodes"
  "codex/tree/visitors"
)

type InsertManager struct {
  Tree *nodes.InsertStatement
}

type Values map[interface{}]interface{}

func (values Values) Columns() []interface{} {
  cols := make([]interface{}, len(values))
  for column, _ := range values {
    cols = append(cols, column)
  }

  return cols
}

func (values Values) Values() []interface{} {
  vals := make([]interface{}, len(values))
  for _, value := range values {
    vals = append(vals, value)
  }

  return vals
}

func (mgmt *InsertManager) Insert(values ...interface{}) *InsertManager {
  value := values[0]
  switch value.(type) {
  case Values:
    return mgmt.InsertValues(value.(Values))
  default:
    mgmt.Tree.Values = append(mgmt.Tree.Values, values...)
  }
  return mgmt
}

func (mgmt *InsertManager) InsertValues(values Values) *InsertManager {
  for column, value := range values {
    mgmt.Tree.Columns = append(mgmt.Tree.Columns, column)
    mgmt.Tree.Values = append(mgmt.Tree.Values, value)
  }
  return mgmt
}

func (mgmt *InsertManager) Into(columns ...interface{}) *InsertManager {
  mgmt.Tree.Columns = append(mgmt.Tree.Columns, columns...)

  return mgmt
}

func (mgmt *InsertManager) ToSql() string {
  visitor := &visitors.ToSqlVisitor{}
  return visitor.Accept(mgmt.Tree)
}
