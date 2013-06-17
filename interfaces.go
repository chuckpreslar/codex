package librarian

import (
  "librarian/nodes"
  "librarian/visitors"
)

type ManagerInterface interface {
  Visitor() visitors.VisitorInterface
  Where(nodes.ComparisonInterface) ManagerInterface
  ToSql() string
}
