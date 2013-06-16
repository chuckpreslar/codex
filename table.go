package librarian

import (
  "librarian/nodes"
)

type Table func(string) nodes.AttributeInterface

func (table Table) As(aliases ...string) Table {
  table("").Reference().AddAliases(aliases...)
  return table
}

func NewTable(name string) Table {
  reference := nodes.Reference(name)
  return func(name string) nodes.AttributeInterface {
    return nodes.Attribute(reference, name)
  }
}
