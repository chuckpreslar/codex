package sql

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes/base"
)

// String ...
type String struct{}

// VisitStringNode ...
func (s String) VisitStringNode(node *base.String, visitor interfaces.Visitor) string {
	return fmt.Sprintf("'%s'", node.Value)
}
