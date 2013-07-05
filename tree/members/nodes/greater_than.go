package nodes

// GreaterThan node struct
type GreaterThan struct {
  *Binary // Embedded Binary node.
}

// Returns an Or node with leafs containing references
// to the original and other
func (gt *GreaterThan) Or(other interface{}) *Or {
  return &Or{&Binary{gt, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (gt *GreaterThan) And(other interface{}) *And {
  return &And{&Binary{gt, other}}
}
