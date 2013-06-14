package librarian

import (
  "librarian/visitors"
)

var VISITORS = map[string]visitors.Visitor{
  "base": visitors.BaseVisitor{},
}
