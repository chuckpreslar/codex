package nodes

// Equal node struct
type Equal struct {
  *Binary // Embedded Binary node.
}

// Returns an Or node with leafs containing references
// to the original and other
func (eq *Equal) Or(other interface{}) *Or {
  return &Or{&Binary{eq, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (eq *Equal) And(other interface{}) *And {
  return &And{&Binary{eq, other}}
}
