package librarian

func Search(table Table) (stmt *SelectStatement) {
  stmt = new(SelectStatement)
  stmt.table = table
  return
}
