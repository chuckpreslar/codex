package librarian

type Table func(string) Column

func NewTable(table string) Table {
  return func(column string) Column {
    return Column{table, column, ""}
  }
}
