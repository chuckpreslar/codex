package librarian

import (
  "fmt"
  "librarian/nodes"
)

type SelectManager struct {
  ast *nodes.SelectStatement
  ctx *nodes.SelectCore
}

func (mgmt *SelectManager) Select(columns ...string) *SelectManager {
  return mgmt
}

func (mgmt *SelectManager) Where(filter nodes.ComparatorNode) *SelectManager {
  fmt.Println(filter)
  return mgmt
}

func (mgmt *SelectManager) InnerJoin() *SelectManager {
  return mgmt
}

func (mgmt *SelectManager) OuterJoin() *SelectManager {
  return mgmt
}

func (mgmt *SelectManager) LeftJoin() *SelectManager {
  return mgmt
}

func (mgmt *SelectManager) RightJoin() *SelectManager {
  return mgmt
}

func (mgmt *SelectManager) FullJoin() *SelectManager {
  return mgmt
}

func NewSelectManager() *SelectManager {
  mgmt := new(SelectManager)
  return mgmt
}
