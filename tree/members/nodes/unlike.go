package nodes

// Unlike node struct
type Unlike struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (unlike *Unlike) Or(other interface{}) *Or {
  return &Or{&Binary{unlike, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (unlike *Unlike) And(other interface{}) *And {
  return &And{&Binary{unlike, other}}
}
