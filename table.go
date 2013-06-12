package librarian

import (
  "fmt"
)

type table func(string) column

func Table(name string) table {
  return func(column string) column {
    return column(column)
  }
}
