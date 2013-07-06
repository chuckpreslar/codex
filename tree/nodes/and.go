package nodes

// And node struct
type And Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the And and other.
func (and *And) Or(other interface{}) *Grouping {
  return &Grouping{&Or{and, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the And and other.
func (and *And) And(other interface{}) *Grouping {
  return &Grouping{&And{and, other}}
}

// Returns an Not node with and expression containing the
// And node.
func (and *And) Not() *Not {
  return &Not{and}
}
