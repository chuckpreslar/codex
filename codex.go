// Package codex provides a Relational Algebra.
package codex

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
)

// Table returns an Accessor
func Table(name string) managers.Accessor {
  relation := nodes.Relation(name)
  return func(name interface{}) *nodes.AttributeNode {
    if _, ok := name.(string); ok {
      return nodes.Attribute(nodes.Column(name), relation)
    }

    return nodes.Attribute(name, relation)
  }
}
