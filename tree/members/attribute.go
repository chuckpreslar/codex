package members

import (
  "librarian/tree/members/nodes"
)

// Attribute struct.
type Attribute struct {
  Name     string    // Name of the Attribute.
  Relation *Relation // Relation the attribute belongs to.
}

// Returns and Equal node containing a reference to the
// attribute and other
func (attr *Attribute) Eq(other interface{}) *nodes.Equal {
  return &nodes.Equal{&nodes.Binary{attr, other}}
}

// Returns and NotEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Neq(other interface{}) *nodes.NotEqual {
  return &nodes.NotEqual{&nodes.Binary{attr, other}}
}

// Returns and GreaterThan node containing a reference to the
// attribute and other
func (attr *Attribute) Gt(other interface{}) *nodes.GreaterThan {
  return &nodes.GreaterThan{&nodes.Binary{attr, other}}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Gte(other interface{}) *nodes.GreaterThanOrEqual {
  return &nodes.GreaterThanOrEqual{&nodes.Binary{attr, other}}
}

// Returns and LessThan node containing a reference to the
// attribute and other
func (attr *Attribute) Lt(other interface{}) *nodes.LessThan {
  return &nodes.LessThan{&nodes.Binary{attr, other}}
}

// Returns and LessThanOrEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Lte(other interface{}) *nodes.LessThanOrEqual {
  return &nodes.LessThanOrEqual{&nodes.Binary{attr, other}}
}

// Returns and Like node containing a reference to the
// attribute and other
func (attr *Attribute) Like(other interface{}) *nodes.Like {
  return &nodes.Like{&nodes.Binary{attr, other}}
}

// Returns and Unlike node containing a reference to the
// attribute and other
func (attr *Attribute) Unlike(other interface{}) *nodes.Unlike {
  return &nodes.Unlike{&nodes.Binary{attr, other}}
}

// Returns and Or node containing a reference to the
// attribute and other
func (attr *Attribute) Or(other interface{}) *nodes.Or {
  return &nodes.Or{&nodes.Binary{attr, other}}
}

// Returns and And node containing a reference to the
// attribute and other
func (attr *Attribute) And(other interface{}) *nodes.And {
  return &nodes.And{&nodes.Binary{attr, other}}
}
