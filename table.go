package librarian

import (
  "librarian/nodes"
)

type Table func(string) *nodes.AttributeNode

func (table Table) Relation() *nodes.RelationNode {
  return table("").Relation
}

func (table Table) On(a interface{}) *nodes.OnNode {
  return nodes.On(table.Relation(), a)
}

func NewTable(name string) Table {
  relation := nodes.Relation(name)
  return func(name string) *nodes.AttributeNode {
    return nodes.Attribute(name, relation)
  }
}
