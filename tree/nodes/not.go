package nodes

type Not Unary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Not and other.
func (not *Not) Or(other interface{}) *Grouping {
  return &Grouping{&Or{not, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Not and other.
func (not *Not) And(other interface{}) *Grouping {
  return &Grouping{&And{not, other}}
}

// Returns an Not node with and expression containing the
// Not node.
func (not *Not) Not() *Not {
  return &Not{not}
}
