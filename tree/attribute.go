package tree

import (
  "librarian/tree/nodes"
)

type Attribute struct {
  Name     string
  Relation *Relation
}

func (a *Attribute) Eq(other interface{}) *nodes.Equal {
  return &nodes.Equal{&nodes.Binary{a, other}}
}
