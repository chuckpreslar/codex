package nodes

// Equal node struct
type Unlike Binary

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Unlike and other.
func (unlike *Unlike) Or(other interface{}) *Grouping {
  return &Grouping{&Or{unlike, other}}
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Unlike and other.
func (unlike *Unlike) And(other interface{}) *Grouping {
  return &Grouping{&And{unlike, other}}
}

// Returns an Not node with and expression containing the
// Unlike node.
func (unlike *Unlike) Not() *Not {
  return &Not{unlike}
}
