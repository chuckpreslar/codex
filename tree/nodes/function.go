package nodes

// Function node is a Nary node.
type Function struct {
  Expressions []interface{} // Expressions the function is operating on.
  Alias       interface{}   // Alias the function result is reffered to as.
  Distinct    bool          // Function is distinct.
}
