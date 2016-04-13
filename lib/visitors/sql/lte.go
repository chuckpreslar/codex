package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Lte ...
type Lte struct{}

// VisitLteNode ...
func (l Lte) VisitLteNode(node *nodes.Lte, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Lte, node.Right.Accept(visitor))
}
