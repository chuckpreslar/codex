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

func (table Table) Project(names ...interface{}) *SelectManager {
  mgr := Select(table.Relation())
  for _, name := range names {
    switch name.(type) {
    case string:
      name = nodes.Attribute(table.Relation(), name.(string))
    case *nodes.AttributeNode:
    default:
      panic("Unkown argument type.")
    }
    mgr.Project(name.(nodes.AttributeInterface))
  }
  return mgr
}

func NewTable(name string) Table {
  relation := nodes.Relation(name)
  return func(name string) nodes.AttributeInterface {
    return nodes.Attribute(relation, name)
  }
}
