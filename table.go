package tree

import (
  "librarian/tree/managers"
  "librarian/tree/nodes"
)

type Accessor func(string) *nodes.Attribute

func (accessor Accessor) Relation() *nodes.Relation {
  return accessor("").Relation
}

func (accessor Accessor) Name() string {
  return accessor("").Relation.Name
}

func (accessor Accessor) From(table *nodes.Relation) *managers.SelectManager {
  statement := &nodes.SelectStatement{
    make([]*nodes.SelectCore, 0),
    &nodes.Limit{}, &nodes.Offset{},
  }
  core := &nodes.SelectCore{table,
    &nodes.JoinSource{table, []interface{}{}},
    []interface{}{}, []interface{}{},
  }
  statement.Cores = append(statement.Cores, core)
  return &managers.SelectManager{statement, core}
}

func (accessor Accessor) Project(projections ...interface{}) *managers.SelectManager {
  return accessor.From(accessor.Relation()).Project(projections...)
}

func (accessor Accessor) Where(expr interface{}) *managers.SelectManager {
  return accessor.From(accessor.Relation()).Where(expr)
}

func Table(name string) Accessor {
  relation := &nodes.Relation{name, ""}
  return func(name string) *nodes.Attribute {
    return &nodes.Attribute{name, relation}
  }
}
