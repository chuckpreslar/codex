package nodes

// And node struct
type And Binary

// Returns an Or node with leafs containing references
// to the original and other
func (and *And) Or(other interface{}) *Grouping {
  return &Grouping{&Or{and, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (and *And) And(other interface{}) *Grouping {
  return &Grouping{&And{and, other}}
}
