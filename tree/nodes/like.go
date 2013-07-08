package nodes

// Like node is a Binary node struct
type Like Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Like and other.
func (like *Like) Or(other interface{}) *Grouping {
  return &Grouping{&Or{like, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Like and other.
func (like *Like) And(other interface{}) *Grouping {
  return &Grouping{&And{like, other}}
}

// Returns an Not node with and expression containing the
// Like node.
func (like *Like) Not() *Not {
  return &Not{like}
}
