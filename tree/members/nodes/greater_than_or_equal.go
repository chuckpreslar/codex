package nodes

// GreaterThanOrEqual node struct
type GreaterThanOrEqual Binary

// Returns an Or node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{gte, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{gte, other}}
}
