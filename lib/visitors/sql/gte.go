package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Gte ...
type Gte struct{}

// VisitGteNode ...
func (g Gte) VisitGteNode(node *nodes.Gte, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Gte, node.Right.Accept(visitor))
}
