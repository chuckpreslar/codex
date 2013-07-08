package nodes

// Unary node struct
type Unary struct {
  Expr interface{} // Single leaf for the Unary node.
}

type Literal Unary           // Literal node is a Unary node struct
type On Unary                // On node is a Unary node struct
type Limit Unary             // Limit node is a Unary node struct
type Offset Unary            // Offset node is a Unary node struct
type Group Unary             // Group node is a Unary node struct
type Having Unary            // Having node is a Unary node struct
type UnqualifiedColumn Unary // UnqualifiedColumn node is a Unary node struct
