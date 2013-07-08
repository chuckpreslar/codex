package nodes

// SelectCore is a Nary node, normally contained in a SelectStatement node.
type SelectCore struct {
  Relation    *Relation     // Pointer to the relation the SelectCore is acting on.
  Source      *JoinSource   // JoinSouce for joining other SQL tables.
  Projections []interface{} // Projections is an array, normally columns found on the SQL table.
  Wheres      []interface{} // Wheres is an array of filters for the acting on the SelectCore.
}
