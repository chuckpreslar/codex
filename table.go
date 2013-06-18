package librarian

import (
  "librarian/nodes"
)

type Table func(string) *nodes.AttributeNode

func NewTable(name string) Table {
  relation := nodes.Relation(name)
  return func(name string) *nodes.AttributeNode {
    return nodes.Attribute(name, relation)
  }
}
