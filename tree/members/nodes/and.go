package nodes

// And node struct
type And struct {
  *Binary // Embedded Binary node.
}

// Returns an Or node with leafs containing references
// to the original and other
func (and *And) Or(other interface{}) *Or {
  return &Or{&Binary{and, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (and *And) And(other interface{}) *And {
  return &And{&Binary{and, other}}
}
