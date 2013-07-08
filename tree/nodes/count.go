package nodes

type Count Function

// Returns and Equal node containing a reference to the
// function and other
func (count *Count) Eq(other interface{}) *Equal {
  return &Equal{count, other}
}

// Returns and NotEqual node containing a reference to the
// function and other
func (count *Count) Neq(other interface{}) *NotEqual {
  return &NotEqual{count, other}
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (count *Count) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{count, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (count *Count) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{count, other}
}

// Returns and LessThan node containing a reference to the
// function and other
func (count *Count) Lt(other interface{}) *LessThan {
  return &LessThan{count, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (count *Count) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{count, other}
}

// Returns and Like node containing a reference to the
// function and other
func (count *Count) Like(other interface{}) *Like {
  return &Like{count, other}
}

// Returns and Unlike node containing a reference to the
// function and other
func (count *Count) Unlike(other interface{}) *Unlike {
  return &Unlike{count, other}
}

// Returns and Or node containing a reference to the
// function and other
func (count *Count) Or(other interface{}) *Grouping {
  return &Grouping{&Or{count, other}}
}

// Returns and And node containing a reference to the
// function and other
func (count *Count) And(other interface{}) *Grouping {
  return &Grouping{&And{count, other}}
}
