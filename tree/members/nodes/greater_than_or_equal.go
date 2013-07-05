package nodes

// GreaterThanOrEqual node struct
type GreaterThanOrEqual Binary

// Returns an Or node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) Or(other interface{}) *Or {
  return &Or{gte, other}
}

// Returns an And node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) And(other interface{}) *And {
  return &And{gte, other}
}
