package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Attribute ...
type Attribute struct{}

// VisitAttributeNode ...
func (a Attribute) VisitAttributeNode(node *nodes.Attribute, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%s.`%s`", node.Relation.Accept(visitor), node.Name)
}
