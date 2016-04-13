package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
)

// Or ...
type Or struct{}

// VisitOrNode ...
func (o Or) VisitOrNode(node *nodes.Or, visitor interfaces.Visitor) string {
	return fmt.Sprintf("(%s %s %s)", node.Left.Accept(visitor), consts.Or, node.Right.Accept(visitor))
}
