package librarian

import (
  "librarian/nodes"
)

type Table func(string) nodes.Attribute

func NewTable(name string) Table {
  r := nodes.NewReference(name)
  return func(name string) nodes.Attribute {
    return nodes.NewAttribute(name, r)
  }
}
