package nodes

// GreaterThanOrEqual node struct
type GreaterThanOrEqual Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the GreaterThanOrEqual and other.
func (gte *GreaterThanOrEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{gte, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the GreaterThanOrEqual and other.
func (gte *GreaterThanOrEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{gte, other}}
}

// Returns an Not node with and expression containing the
// GreaterThanOrEqual node.
func (gte *GreaterThanOrEqual) Not() *Not {
  return &Not{gte}
}
