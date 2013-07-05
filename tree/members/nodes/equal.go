package nodes

// Equal node struct
type Equal Binary

// Returns an Or node with leafs containing references
// to the original and other
func (eq *Equal) Or(other interface{}) *Or {
  return &Or{eq, other}
}

// Returns an And node with leafs containing references
// to the original and other
func (eq *Equal) And(other interface{}) *And {
  return &And{eq, other}
}
