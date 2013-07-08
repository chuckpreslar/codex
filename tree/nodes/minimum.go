package nodes

type Minimum Function

// Returns and Equal node containing a reference to the
// function and other
func (min *Minimum) Eq(other interface{}) *Equal {
  return &Equal{min, other}
}

// Returns and NotEqual node containing a reference to the
// function and other
func (min *Minimum) Neq(other interface{}) *NotEqual {
  return &NotEqual{min, other}
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (min *Minimum) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{min, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (min *Minimum) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{min, other}
}

// Returns and LessThan node containing a reference to the
// function and other
func (min *Minimum) Lt(other interface{}) *LessThan {
  return &LessThan{min, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (min *Minimum) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{min, other}
}

// Returns and Like node containing a reference to the
// function and other
func (min *Minimum) Like(other interface{}) *Like {
  return &Like{min, other}
}

// Returns and Unlike node containing a reference to the
// function and other
func (min *Minimum) Unlike(other interface{}) *Unlike {
  return &Unlike{min, other}
}

// Returns and Or node containing a reference to the
// function and other
func (min *Minimum) Or(other interface{}) *Grouping {
  return &Grouping{&Or{min, other}}
}

// Returns and And node containing a reference to the
// function and other
func (min *Minimum) And(other interface{}) *Grouping {
  return &Grouping{&And{min, other}}
}
