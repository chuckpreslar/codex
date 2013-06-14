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

func (table Table) Select(columns ...string) *SelectManager {
  return NewSelectManager().Select(columns...)
}

func (table Table) Where(filter nodes.ComparatorNode) *SelectManager {
  return NewSelectManager().Where(filter)
}

func (table Table) InnerJoin() *SelectManager {
  return NewSelectManager().InnerJoin()
}

func (table Table) OuterJoin() *SelectManager {
  return NewSelectManager().OuterJoin()
}

func (table Table) LeftJoin() *SelectManager {
  return NewSelectManager().LeftJoin()
}

func (table Table) RightJoin() *SelectManager {
  return NewSelectManager().RightJoin()
}

func (table Table) FullJoin() *SelectManager {
  return NewSelectManager().FullJoin()
}
