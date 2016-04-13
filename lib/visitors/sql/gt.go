package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Gt ...
type Gt struct{}

// VisitGtNode ...
func (g Gt) VisitGtNode(node *nodes.Gt, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Gt, node.Right.Accept(visitor))
}
