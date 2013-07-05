package nodes

// LessThanOrEqual node struct
type LessThanOrEqual struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (lte *LessThanOrEqual) Or(other interface{}) *Or {
  return &Or{&Binary{lte, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (lte *LessThanOrEqual) And(other interface{}) *And {
  return &And{&Binary{lte, other}}
}
