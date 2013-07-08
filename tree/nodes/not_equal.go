package nodes

// NotEqual node is a Binary node struct
type NotEqual Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the NotEqual and other.
func (neq *NotEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{neq, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the NotEqual and other.
func (neq *NotEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{neq, other}}
}

// Returns an Not node with and expression containing the
// NotEqual node.
func (neq *NotEqual) Not() *Not {
  return &Not{neq}
}
