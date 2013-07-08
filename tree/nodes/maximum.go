package nodes

type Maximum Function

// Returns and Equal node containing a reference to the
// function and other
func (max *Maximum) Eq(other interface{}) *Equal {
  return &Equal{max, other}
}

// Returns and NotEqual node containing a reference to the
// function and other
func (max *Maximum) Neq(other interface{}) *NotEqual {
  return &NotEqual{max, other}
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (max *Maximum) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{max, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (max *Maximum) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{max, other}
}

// Returns and LessThan node containing a reference to the
// function and other
func (max *Maximum) Lt(other interface{}) *LessThan {
  return &LessThan{max, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (max *Maximum) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{max, other}
}

// Returns and Like node containing a reference to the
// function and other
func (max *Maximum) Like(other interface{}) *Like {
  return &Like{max, other}
}

// Returns and Unlike node containing a reference to the
// function and other
func (max *Maximum) Unlike(other interface{}) *Unlike {
  return &Unlike{max, other}
}

// Returns and Or node containing a reference to the
// function and other
func (max *Maximum) Or(other interface{}) *Grouping {
  return &Grouping{&Or{max, other}}
}

// Returns and And node containing a reference to the
// function and other
func (max *Maximum) And(other interface{}) *Grouping {
  return &Grouping{&And{max, other}}
}
