package librarian

func Search(table Table) *SelectStatement {
  stmt := new(SelectStatement)
  stmt.reference = table
  return stmt
}
