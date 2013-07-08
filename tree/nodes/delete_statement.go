package nodes

// DeleteStatement is the base node for SQL Delete Statements.
type DeleteStatement struct {
  Relation *Relation     // Pointer to the Relation the Delete Statement is acting on.
  Wheres   []interface{} // Wheres is an array of expressions/nodes.
}
