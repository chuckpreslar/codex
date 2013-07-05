package nodes

// Or node struct
type Or Binary

// Returns an Or node with leafs containing references
// to the original and other
func (or *Or) Or(other interface{}) *Grouping {
  return &Grouping{&Or{or, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (or *Or) And(other interface{}) *Grouping {
  return &Grouping{&And{or, other}}
}
