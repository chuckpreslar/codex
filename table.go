package librarian

import (
  "librarian/nodes"
)

type Table func(string) nodes.AttributeInterface

func (table Table) As(aliases ...string) Table {
  table("").Relation().AddAliases(aliases...)
  return table
}

func (table Table) Relation() nodes.RelationInterface {
  return table("").Relation()
}

func (table Table) TableName() string {
  return table("").Relation().Name()
}

func NewTable(name string) Table {
  relation := nodes.Relation(name)
  return func(name string) nodes.AttributeInterface {
    return nodes.Attribute(relation, name)
  }
}
