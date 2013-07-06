package nodes

// Attribute struct.
type Attribute struct {
  Name     string    // Name of the Attribute.
  Relation *Relation // Relation the attribute belongs to.
}

// Returns and Equal node containing a reference to the
// attribute and other
func (attr *Attribute) Eq(other interface{}) *Equal {
  return &Equal{attr, other}
}

// Returns and NotEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Neq(other interface{}) *NotEqual {
  return &NotEqual{attr, other}
}

// Returns and GreaterThan node containing a reference to the
// attribute and other
func (attr *Attribute) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{attr, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{attr, other}
}

// Returns and LessThan node containing a reference to the
// attribute and other
func (attr *Attribute) Lt(other interface{}) *LessThan {
  return &LessThan{attr, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// attribute and other
func (attr *Attribute) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{attr, other}
}

// Returns and Like node containing a reference to the
// attribute and other
func (attr *Attribute) Like(other interface{}) *Like {
  return &Like{attr, other}
}

// Returns and Unlike node containing a reference to the
// attribute and other
func (attr *Attribute) Unlike(other interface{}) *Unlike {
  return &Unlike{attr, other}
}

// Returns and Or node containing a reference to the
// attribute and other
func (attr *Attribute) Or(other interface{}) *Grouping {
  return &Grouping{&Or{attr, other}}
}

// Returns and And node containing a reference to the
// attribute and other
func (attr *Attribute) And(other interface{}) *Grouping {
  return &Grouping{&And{attr, other}}
}
