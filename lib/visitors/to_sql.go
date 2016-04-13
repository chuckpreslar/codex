package visitors

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
	"github.com/chuckpreslar/codex/lib/nodes"
	"github.com/chuckpreslar/codex/lib/nodes/base"
	"github.com/chuckpreslar/codex/lib/visitors/sql"
)

// ToSQL ...
type ToSQL struct {
	sql.Attribute
	sql.Relation
	sql.Eq
	sql.Neq
	sql.Gt
	sql.Gte
	sql.Lt
	sql.Lte
	sql.Or
	sql.And
	sql.Nil
	sql.Integer
	sql.String
	sql.Boolean
}

// Visit ...
func (t ToSQL) Visit(acceptor interfaces.Acceptor) string {
	switch acceptor.(type) {
	case *nodes.Attribute:
		return t.VisitAttributeNode(acceptor.(*nodes.Attribute), t)
	case *nodes.Relation:
		return t.VisitRelationNode(acceptor.(*nodes.Relation), t)
	case *nodes.Eq:
		return t.VisitEqNode(acceptor.(*nodes.Eq), t)
	case *nodes.Neq:
		return t.VisitNeqNode(acceptor.(*nodes.Neq), t)
	case *nodes.Gt:
		return t.VisitGtNode(acceptor.(*nodes.Gt), t)
	case *nodes.Gte:
		return t.VisitGteNode(acceptor.(*nodes.Gte), t)
	case *nodes.Lt:
		return t.VisitLtNode(acceptor.(*nodes.Lt), t)
	case *nodes.Lte:
		return t.VisitLteNode(acceptor.(*nodes.Lte), t)
	case *nodes.Or:
		return t.VisitOrNode(acceptor.(*nodes.Or), t)
	case *nodes.And:
		return t.VisitAndNode(acceptor.(*nodes.And), t)
	case *base.Nil:
		return t.VisitNilNode(acceptor.(*base.Nil), t)
	case *base.Integer:
		return t.VisitIntegerNode(acceptor.(*base.Integer), t)
	case *base.String:
		return t.VisitStringNode(acceptor.(*base.String), t)
	case *base.Boolean:
		return t.VisitBooleanNode(acceptor.(*base.Boolean), t)
	default:
		return ""
	}
}
