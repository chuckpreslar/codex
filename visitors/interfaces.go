package visitors

import (
  "librarian/nodes"
)

type VisitorInterface interface {
  Accept(nodes.NodeInterface) string
  visit(interface{}) string
}
