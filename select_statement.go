package librarian

type SelectStatement struct {
  reference   Table
  projections []Column
  filters     []Comparable
  offset      int
  limit       int
}

func (stmt *SelectStatement) Select(columns ...interface{}) *SelectStatement {
  for _, column := range columns {
    switch column.(type) {
    case string:
      column = stmt.reference(column.(string))
    case Column:
      // Already in needed type.
    default:
      // Unacceptable type.
    }
    stmt.projections = append(stmt.projections, column.(Column))
  }
  return stmt
}

func (stmt *SelectStatement) Where(comparator Comparable) *SelectStatement {
  stmt.filters = append(stmt.filters, comparator)
  return stmt
}
