package librarian

import (
  "librarian/nodes"
)

type Table func(string) nodes.Attribute

func (table Table) Select(projections ...interface{}) *SelectManager {
  return NewSelectManager(table("").Reference()).Select(projections...)
}

func (table Table) Where(comparison nodes.Comparison) *SelectManager {
  return NewSelectManager(table("").Reference()).Where(comparison)
}

func NewTable(name string) Table {
  r := nodes.NewReference(name)
  return func(name string) nodes.Attribute {
    return nodes.NewAttribute(name, r)
  }
}
