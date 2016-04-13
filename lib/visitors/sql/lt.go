package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Lt ...
type Lt struct{}

// VisitLtNode ...
func (l Lt) VisitLtNode(node *nodes.Lt, visitor interfaces.Visitor) string {
	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Lt, node.Right.Accept(visitor))
}
