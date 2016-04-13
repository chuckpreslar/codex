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

// Neq ...
type Neq struct{}

// VisitNeqNode ...
func (n Neq) VisitNeqNode(node *nodes.Neq, visitor interfaces.Visitor) string {
	var ok bool

	if _, ok = node.Right.(*base.Nil); ok {
		return fmt.Sprintf("%s IS NOT %s", node.Left.Accept(visitor), node.Right.Accept(visitor))
	}

	return fmt.Sprintf("%s %s %s", node.Left.Accept(visitor), consts.Neq, node.Right.Accept(visitor))
}
