package nodes

// Equal node struct
type Equal Binary

// Returns an Or node with leafs containing references
// to the original and other
func (eq *Equal) Or(other interface{}) *Grouping {
  return &Grouping{&Or{eq, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (eq *Equal) And(other interface{}) *Grouping {
  return &Grouping{&And{eq, other}}
}
