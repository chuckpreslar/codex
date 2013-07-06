package tree

import (
  "librarian/tree/nodes"
  "fmt"
)

type Accessor func(string) *nodes.Attribute

func (accessor Accessor) Relation() *nodes.Relation {
  return accessor("").Relation
}

func (accessor Accessor) Name() string {
  return accessor("").Relation.Name
}

func Table(name string) Accessor {
  relation := &nodes.Relation{name, ""}
  return func(name string) *nodes.Attribute {
    return &nodes.Attribute{name, relation}
  }
}
