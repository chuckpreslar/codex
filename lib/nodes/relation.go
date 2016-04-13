package nodes

import (
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Relation ...
type Relation struct {
	Name string
}

// Accept ...
func (r *Relation) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(r)
}

// InitializeRelation ...
func InitializeRelation(name string) *Relation {
	var relation = new(Relation)

	relation.Name = name

	return relation
}
