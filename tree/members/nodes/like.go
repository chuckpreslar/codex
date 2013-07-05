package nodes

// Like node struct
type Like Binary

// Returns an Or node with leafs containing references
// to the original and other
func (like *Like) Or(other interface{}) *Grouping {
  return &Grouping{&Or{like, other}}
}

// Returns an And node with leafs containing references
// to the original and other
func (like *Like) And(other interface{}) *Grouping {
  return &Grouping{&And{like, other}}
}
