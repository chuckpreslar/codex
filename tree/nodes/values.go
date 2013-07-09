package nodes

// ValuesNode is a specific BinaryNode.
type ValuesNode struct {
  Expressions []interface{} // Array of expressions/nodes, normally assignments.
  Columns     []interface{} // Array of columns the expressions effect.
}

// ValuesNode factory method.
func Values() *ValuesNode {
  values := new(ValuesNode)
  values.Expressions = make([]interface{}, 0)
  values.Columns = make([]interface{}, 0)
  return values
}
