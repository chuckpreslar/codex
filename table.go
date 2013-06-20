package librarian

import (
  "librarian/nodes"
)

type Table func(string) *nodes.AttributeNode

func (table Table) Relation() *nodes.RelationNode {
  return table("").Relation
}

func (table Table) As(alias string) Table {
  table.Relation().Alias = alias
  return table
}

func (table Table) On(a interface{}) *nodes.OnNode {
  return nodes.On(table.Relation(), a)
}

func (table Table) From(a interface{}) *SelectManager {
  return NewSelectManager(table.Relation()).From(a)
}

func (table Table) Project(a ...interface{}) *SelectManager {
  return NewSelectManager(table.Relation()).Project(a...)
}

func (table Table) Where(a interface{}) *SelectManager {
  return NewSelectManager(table.Relation()).Where(a)
}

func (table Table) InnerJoin(a interface{}) *SelectManager {
  return NewSelectManager(table.Relation()).InnerJoin(a)
}

func (table Table) ToSql() string {
  return NewSelectManager(table.Relation()).ToSql()
}

func NewTable(name string) Table {
  relation := nodes.Relation(name)
  return func(name string) *nodes.AttributeNode {
    return nodes.Attribute(name, relation)
  }
}
