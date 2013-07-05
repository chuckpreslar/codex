package nodes

// Like node struct
type Like struct {
  *Binary
}

// Returns an Or node with leafs containing references
// to the original and other
func (like *Like) Or(other interface{}) *Or {
  return &Or{&Binary{like, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (like *Like) And(other interface{}) *And {
  return &And{&Binary{like, other}}
}
