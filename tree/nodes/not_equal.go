package nodes

// NotEqual node struct
type NotEqual Binary

// Returns an Or node with leafs containing references
// to the original and other
func (neq *NotEqual) Or(other interface{}) *Grouping {
  return &Grouping{&Or{neq, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (neq *NotEqual) And(other interface{}) *Grouping {
  return &Grouping{&And{neq, other}}
}
