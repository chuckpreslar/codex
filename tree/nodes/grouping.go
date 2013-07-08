package nodes

// Grouping node is a Unary node struct
type Grouping Unary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Grouping and other.
func (grouping *Grouping) Or(other interface{}) *Grouping {
  return &Grouping{&Or{grouping, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Grouping and other.
func (grouping *Grouping) And(other interface{}) *Grouping {
  return &Grouping{&And{grouping, other}}
}

// Returns an Not node with and expression containing the
// Grouping node.
func (grouping *Grouping) Not() *Not {
  return &Not{grouping}
}
