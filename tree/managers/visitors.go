package managers

import (
  "github.com/chuckpreslar/codex/tree/visitors"
)

func VisitorFor(engine interface{}) visitors.VisitorInterface {
  switch engine {
  case "mysql":
    return &visitors.MySqlVisitor{&visitors.ToSqlVisitor{}}
  case "postgres":
    return &visitors.PostgresVisitor{&visitors.ToSqlVisitor{}, 0}
  default:
    return &visitors.ToSqlVisitor{}
  }
}
