package nodes

// InsertStatement is the base node for SQL Insert Statements.
type InsertStatement struct {
  Relation *Relation     // Pointer to the Relation the Insert Statement is acting on.
  Columns  []interface{} // Columns the Insert Statement is effecting.
  Values   *Values       // Pointer to the Values for insertion.
}
