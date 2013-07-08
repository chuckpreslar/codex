package nodes

// Values is a specific Binary node.
type Values struct {
  Expressions []interface{} // Array of expressions/nodes, normally assignments.
  Columns     []interface{} // Array of columns the expressions effect.
}
