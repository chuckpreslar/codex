package nodes

// UpdateStatement is the base node for SQL Update Statements.
type UpdateStatement struct {
  Relation *Relation     // Pointer to the Relation the Delete Statement is acting on.
  Values   []interface{} // Values is an array of expressions/nodes.
  Wheres   []interface{} // Wheres is an array of expressions/nodes.
  Limit    *Limit        // Potential Limit node for limiting the number of rows effected.
}
