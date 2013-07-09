package codex

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
)

func Table(name string) managers.Accessor {
  relation := nodes.Relation(name)
  return func(name string) *nodes.AttributeNode {
    return nodes.Attribute(name, relation)
  }
}
