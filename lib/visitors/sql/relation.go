package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Relation ...
type Relation struct{}

// VisitRelationNode ...
func (r Relation) VisitRelationNode(node *nodes.Relation, visitor interfaces.Visitor) string {
	return fmt.Sprintf("`%s`", node.Name)
}
