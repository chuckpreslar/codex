package nodes

// GreaterThanOrEqual node struct
type GreaterThanOrEqual struct {
  *Binary // Embedded Binary node.
}

// Returns an Or node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) Or(other interface{}) *Or {
  return &Or{&Binary{gte, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (gte *GreaterThanOrEqual) And(other interface{}) *And {
  return &And{&Binary{gte, other}}
}
