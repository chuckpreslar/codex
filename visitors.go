package librarian

import (
  "librarian/visitors"
)

var VISITORS = map[string]visitors.VisitorInterface{
  "to_sql": visitors.ToSql(),
}
