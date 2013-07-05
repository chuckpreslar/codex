package nodes

// GreaterThan node struct
type GreaterThan Binary

// Returns an Or node with leafs containing references
// to the original and other
func (gt *GreaterThan) Or(other interface{}) *Or {
  return &Or{gt, other}
}

// Returns an And node with leafs containing references
// to the original and other
func (gt *GreaterThan) And(other interface{}) *And {
  return &And{gt, other}
}
