// Package nodes provides nodes to use in codex AST's.
package nodes

// FunctionNode is a Nary node.
type FunctionNode struct {
	Expressions []interface{} // Expressions the function is operating on.
	Alias       interface{}   // Alias the function result is reffered to as.
	Distinct    bool          // Function is distinct.
}

// FunctionNode factory method.
func Function(expressions ...interface{}) (function *FunctionNode) {
	function = new(FunctionNode)
	function.Expressions = expressions
	return
}
