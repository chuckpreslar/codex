package nodes

// Or node struct
type Or struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (or *Or) Or(other interface{}) *Or {
  return &Or{&Binary{or, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (or *Or) And(other interface{}) *And {
  return &And{&Binary{or, other}}
}
