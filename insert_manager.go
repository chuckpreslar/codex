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
    if nil == mgmt.Tree.Values {
      mgmt.Tree.Values = &nodes.Values{Expressions: append([]interface{}{}, values...)}
    } else {
      mgmt.Tree.Values.Expressions = append([]interface{}{}, values...)
    }
  }
  return mgmt
}

func (mgmt *InsertManager) InsertValues(values Values) *InsertManager {
  mgmt.Tree.Columns = append(mgmt.Tree.Columns, values.Columns()...)
  if nil == mgmt.Tree.Values {
    mgmt.Tree.Values = &nodes.Values{values.Values(), values.Columns()}
  } else {
    mgmt.Tree.Values.Expressions = append([]interface{}{}, values.Values()...)
    mgmt.Tree.Values.Columns = append([]interface{}{}, values.Columns()...)
  }
  return mgmt
}

func (mgmt *InsertManager) Into(columns ...interface{}) *InsertManager {
  if nil == mgmt.Tree.Values {
    mgmt.Tree.Values = &nodes.Values{Columns: []interface{}{}}
  }
  mgmt.Tree.Values.Columns = append([]interface{}{}, columns...)
  mgmt.Tree.Columns = append([]interface{}{}, columns...)

  return mgmt
}

func (mgmt *InsertManager) ToSql() string {
  visitor := &visitors.ToSqlVisitor{}
  return visitor.Accept(mgmt.Tree)
}
