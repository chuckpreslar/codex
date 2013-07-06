package nodes

// LessThanOrEqual node struct
type LessThanOrEqual Binary

// Returns an Or node with leafs containing references
// to the original and other
func (lte *LessThanOrEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{lte, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (lte *LessThanOrEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{lte, other}}
}
