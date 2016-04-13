package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

// Eq ...
type Eq struct{}

// VisitEqNode ...
func (e Eq) VisitEqNode(node *nodes.Eq, visitor interfaces.Visitor) string {
	var ok bool

	if _, ok = node.Right.(*base.Nil); ok {
		return fmt.Sprintf("%s IS %s", node.Left.Accept(visitor), node.Right.Accept(visitor))
	}

	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Eq, node.Right.Accept(visitor))
}
