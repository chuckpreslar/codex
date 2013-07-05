package nodes

// Unlike node struct
type Unlike Binary

// Returns an Or node with leafs containing references
// to the original and other
func (unlike *Unlike) Or(other interface{}) *Or {
  return &Or{unlike, other}
}

// Returns an And node with leafs containing references
// to the original and other
func (unlike *Unlike) And(other interface{}) *And {
  return &And{unlike, other}
}
