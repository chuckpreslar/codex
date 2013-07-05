package nodes

// NotEqual node struct
type NotEqual struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (neq *NotEqual) Or(other interface{}) *Or {
  return &Or{&Binary{neq, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (neq *NotEqual) And(other interface{}) *And {
  return &And{&Binary{neq, other}}
}
