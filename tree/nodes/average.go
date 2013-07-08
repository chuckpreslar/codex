package nodes

type Average Function

// Returns and Equal node containing a reference to the
// function and other
func (avg *Average) Eq(other interface{}) *Equal {
  return &Equal{avg, other}
}

// Returns and NotEqual node containing a reference to the
// function and other
func (avg *Average) Neq(other interface{}) *NotEqual {
  return &NotEqual{avg, other}
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (avg *Average) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{avg, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (avg *Average) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{avg, other}
}

// Returns and LessThan node containing a reference to the
// function and other
func (avg *Average) Lt(other interface{}) *LessThan {
  return &LessThan{avg, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (avg *Average) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{avg, other}
}

// Returns and Like node containing a reference to the
// function and other
func (avg *Average) Like(other interface{}) *Like {
  return &Like{avg, other}
}

// Returns and Unlike node containing a reference to the
// function and other
func (avg *Average) Unlike(other interface{}) *Unlike {
  return &Unlike{avg, other}
}

// Returns and Or node containing a reference to the
// function and other
func (avg *Average) Or(other interface{}) *Grouping {
  return &Grouping{&Or{avg, other}}
}

// Returns and And node containing a reference to the
// function and other
func (avg *Average) And(other interface{}) *Grouping {
  return &Grouping{&And{avg, other}}
}
