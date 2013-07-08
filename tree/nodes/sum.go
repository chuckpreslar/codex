package nodes

type Sum Function

// Returns and Equal node containing a reference to the
// function and other
func (sum *Sum) Eq(other interface{}) *Equal {
  return &Equal{sum, other}
}

// Returns and NotEqual node containing a reference to the
// function and other
func (sum *Sum) Neq(other interface{}) *NotEqual {
  return &NotEqual{sum, other}
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (sum *Sum) Gt(other interface{}) *GreaterThan {
  return &GreaterThan{sum, other}
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (sum *Sum) Gte(other interface{}) *GreaterThanOrEqual {
  return &GreaterThanOrEqual{sum, other}
}

// Returns and LessThan node containing a reference to the
// function and other
func (sum *Sum) Lt(other interface{}) *LessThan {
  return &LessThan{sum, other}
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (sum *Sum) Lte(other interface{}) *LessThanOrEqual {
  return &LessThanOrEqual{sum, other}
}

// Returns and Like node containing a reference to the
// function and other
func (sum *Sum) Like(other interface{}) *Like {
  return &Like{sum, other}
}

// Returns and Unlike node containing a reference to the
// function and other
func (sum *Sum) Unlike(other interface{}) *Unlike {
  return &Unlike{sum, other}
}

// Returns and Or node containing a reference to the
// function and other
func (sum *Sum) Or(other interface{}) *Grouping {
  return &Grouping{&Or{sum, other}}
}

// Returns and And node containing a reference to the
// function and other
func (sum *Sum) And(other interface{}) *Grouping {
  return &Grouping{&And{sum, other}}
}
