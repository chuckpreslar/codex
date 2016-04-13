package sql

import (
	"github.com/chuckpreslar/codex/lib/aux/consts"
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

// Nil ...
type Nil struct{}

// VisitNilNode ...
func (l Nil) VisitNilNode(node *base.Nil, visitor interfaces.Visitor) string {
	return consts.Nil
}
