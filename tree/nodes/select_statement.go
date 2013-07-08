package nodes

// SelectStatement is the base node for SQL Select Statements.
type SelectStatement struct {
  Cores  []*SelectCore // An array of SelectCores.
  Limit  *Limit        // Potential Limit node for limiting the number of results returned.
  Offset *Offset       // Potential Offset node for skipping records.
}
