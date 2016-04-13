package nodes

import (
	"github.com/chuckpreslar/codex/lib/aux"
	"github.com/chuckpreslar/codex/lib/interfaces"
)

// Attribute ...
type Attribute struct {
	Name     string
	Relation *Relation
}

// Eq ...
func (a *Attribute) Eq(any interface{}) *Eq {
	return InitializeEq(a, aux.CreateAcceptorFrom(any))
}

// Equals ...
func (a *Attribute) Equals(any interface{}) *Eq {
	return a.Eq(any)
}

// Neq ...
func (a *Attribute) Neq(any interface{}) *Neq {
	return InitializeNeq(a, aux.CreateAcceptorFrom(any))
}

// DoesNotEqual ...
func (a *Attribute) DoesNotEqual(any interface{}) *Neq {
	return a.Neq(any)
}

// Gt ...
func (a *Attribute) Gt(any interface{}) *Gt {
	return InitializeGt(a, aux.CreateAcceptorFrom(any))
}

// IsGreaterThan ...
func (a *Attribute) IsGreaterThan(any interface{}) *Gt {
	return a.Gt(any)
}

// Gte ...
func (a *Attribute) Gte(any interface{}) *Gte {
	return InitializeGte(a, aux.CreateAcceptorFrom(any))
}

// IsGreaterThanOrEqual ...
func (a *Attribute) IsGreaterThanOrEqual(any interface{}) *Gte {
	return a.Gte(any)
}

// Lt ...
func (a *Attribute) Lt(any interface{}) *Lt {
	return InitializeLt(a, aux.CreateAcceptorFrom(any))
}

// IsLessThan ...
func (a *Attribute) IsLessThan(any interface{}) *Lt {
	return a.Lt(any)
}

// Lte ...
func (a *Attribute) Lte(any interface{}) *Lte {
	return InitializeLte(a, aux.CreateAcceptorFrom(any))
}

// IsLessThanOrEqual ...
func (a *Attribute) IsLessThanOrEqual(any interface{}) *Lte {
	return a.Lte(any)
}

// Accept ...
func (a *Attribute) Accept(visitor interfaces.Visitor) string {
	return visitor.Visit(a)
}

// InitializeAttribute ...
func InitializeAttribute(name string, relation *Relation) *Attribute {
	var attribute = new(Attribute)

	attribute.Name = name
	attribute.Relation = relation

	return attribute
}
