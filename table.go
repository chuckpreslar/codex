package codex

import (
  "codex/tree/nodes"
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
  return &SelectManager{Tree: statement, Context: core}
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

func (accessor Accessor) Insert(expr ...interface{}) *InsertManager {
  statement := &nodes.InsertStatement{Relation: accessor.Relation()}
  magager := &InsertManager{Tree: statement}
  return magager.Insert(expr...)
}

func (accessor Accessor) Set(expr ...interface{}) *UpdateManager {
  statement := &nodes.UpdateStatement{Relation: accessor.Relation()}
  manager := &UpdateManager{Tree: statement}
  return manager.Set(expr...)
}

func (accessor Accessor) Delete(expr interface{}) *DeleteManager {
  statement := &nodes.DeleteStatement{accessor.Relation(), []interface{}{expr}}
  return &DeleteManager{Tree: statement}
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
