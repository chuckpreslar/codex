package nodes

// LessThan node struct
type LessThan struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (lt *LessThan) Or(other interface{}) *Or {
  return &Or{&Binary{lt, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (lt *LessThan) And(other interface{}) *And {
  return &And{&Binary{lt, other}}
}
