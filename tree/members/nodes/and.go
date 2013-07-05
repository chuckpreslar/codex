package nodes

// And node struct
type And Binary

// Returns an Or node with leafs containing references
// to the original and other
func (and *And) Or(other interface{}) *Or {
  return &Or{and, other}
}

// Returns an And node with leafs containing references
// to the original and other
func (and *And) And(other interface{}) *And {
  return &And{and, other}
}
