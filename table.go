package codex

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
)

func Table(name string) managers.Accessor {
  relation := new(nodes.Relation)
  relation.Name = name
  return func(name string) *nodes.Attribute {
    attr := new(nodes.Attribute)
    attr.Name = name
    attr.Relation = relation
    return attr
  }
}
