package nodes

// GreaterThan node is a Binary node struct
type GreaterThan Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThan and other.
func (gt *GreaterThan) Or(other interface{}) *Grouping {
  return &Grouping{&Or{gt, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThan and other.
func (gt *GreaterThan) And(other interface{}) *Grouping {
  return &Grouping{&And{gt, other}}
}

// Returns an Not node with and expression containing the
// GreaterThan node.
func (gt *GreaterThan) Not() *Not {
  return &Not{gt}
}
