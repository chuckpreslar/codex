package nodes

// LessThanOrEqual node is a Binary node struct
type LessThanOrEqual Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the LessThanOrEqual and other.
func (lte *LessThanOrEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{lte, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the LessThanOrEqual and other.
func (lte *LessThanOrEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{lte, other}}
}

// Returns an Not node with and expression containing the
// LessThanOrEqual node.
func (lte *LessThanOrEqual) Not() *Not {
  return &Not{lte}
}
