package nodes

// LessThan node is a Binary node struct
type LessThan Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThan and other.
func (lt *LessThan) Or(other interface{}) *Grouping {
  return &Grouping{&Or{lt, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThan and other.
func (lt *LessThan) And(other interface{}) *Grouping {
  return &Grouping{&And{lt, other}}
}

// Returns an Not node with and expression containing the
// LessThan node.
func (lt *LessThan) Not() *Not {
  return &Not{lt}
}
