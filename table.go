package librarian

type Table func(string) Column

func NewTable(table string) {
  return func(column string) Column {
    return Column{table, column}
  }
}
