package nodes

// Binary node struct
type Binary struct {
  Left  interface{} // Binary nodes left leaf.
  Right interface{} // Binary nodes right leaf.
}

type As Binary         // As node is a Binary node struct
type Between Binary    // Between node is a Binary node struct
type InnerJoin Binary  // InnerJoin node is a Binary node struct
type OuterJoin Binary  // OuterJoin node is a Binary node struct
type Assignment Binary // Assignment node is a Binary node struct
