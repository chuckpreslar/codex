package nodes

// Or node is a Binary node struct
type Or Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Or and other.
func (or *Or) Or(other interface{}) *Grouping {
  return &Grouping{&Or{or, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Or and other.
func (or *Or) And(other interface{}) *Grouping {
  return &Grouping{&And{or, other}}
}

// Returns an Not node with and expression containing the
// Or node.
func (or *Or) Not() *Not {
  return &Not{or}
}
