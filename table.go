package librarian

import (
  "librarian/nodes"
)

type Table func(string) nodes.AttributeInterface

func (table Table) As(aliases ...string) Table {
  table("").Reference().AddAliases(aliases...)
  return table
}

func (table Table) Reference() nodes.ReferenceInterface {
  return table("").Reference()
}

func (table Table) TableName() string {
  return table("").Reference().Name()
}

func NewTable(name string) Table {
  reference := nodes.Reference(name)
  return func(name string) nodes.AttributeInterface {
    return nodes.Attribute(reference, name)
  }
}
