package nodes

// Unlike node struct
type Unlike Binary

// Returns an Or node with leafs containing references
// to the original and other
func (unlike *Unlike) Or(other interface{}) *Grouping {
  return &Grouping{&Or{unlike, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (unlike *Unlike) And(other interface{}) *Grouping {
  return &Grouping{&And{unlike, other}}
}
