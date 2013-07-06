package librarian

import (
  "librarian/tree/nodes"
)

type Accessor func(string) *nodes.Attribute

func (accessor Accessor) Relation() *nodes.Relation {
  return accessor("").Relation
}

func (accessor Accessor) Name() string {
  return accessor("").Relation.Name
}

func (accessor Accessor) From(table *nodes.Relation) *SelectManager {
  statement := &nodes.SelectStatement{
    Cores: make([]*nodes.SelectCore, 0),
  }
  core := &nodes.SelectCore{table,
    &nodes.JoinSource{table, []interface{}{}},
    []interface{}{}, []interface{}{},
  }
  statement.Cores = append(statement.Cores, core)
  return &SelectManager{statement, core}
}

func (accessor Accessor) Project(projections ...interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Project(projections...)
}

func (accessor Accessor) Where(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).Where(expr)
}

func (accessor Accessor) InnerJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).InnerJoin(expr)
}

func (accessor Accessor) OuterJoin(expr interface{}) *SelectManager {
  return accessor.From(accessor.Relation()).OuterJoin(expr)
}

func (accessor Accessor) ToSql() string {
  return accessor.From(accessor.Relation()).ToSql()
}

func Table(name string) Accessor {
  relation := &nodes.Relation{name, ""}
  return func(name string) *nodes.Attribute {
    return &nodes.Attribute{name, relation}
  }
}
