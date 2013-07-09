package codex

import (
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
)

func Table(name string) managers.Accessor {
  relation := nodes.Relation(name)
  return func(name interface{}) *nodes.AttributeNode {
    switch name.(type) {
    case string:
      return nodes.Attribute(nodes.Column(name), relation)
    default:
      return nodes.Attribute(name, relation)
    }

  }
}
