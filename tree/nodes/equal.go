package nodes

// Equal node is a Binary node struct
type Equal Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Equal and other.
func (eq *Equal) Or(other interface{}) *Grouping {
  return &Grouping{&Or{eq, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Equal and other.
func (eq *Equal) And(other interface{}) *Grouping {
  return &Grouping{&And{eq, other}}
}

// Returns an Not node with and expression containing the
// Equal node.
func (eq *Equal) Not() *Not {
  return &Not{eq}
}
