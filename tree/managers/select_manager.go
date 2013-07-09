package managers

import (
  "github.com/chuckpreslar/codex/factory"
  "github.com/chuckpreslar/codex/tree/nodes"
)

type SelectManager struct {
  Tree    *nodes.SelectStatement
  Context *nodes.SelectCore
  Engine  interface{}
}

func (mgmt *SelectManager) Project(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = factory.Literal(projection)
    default:
    }
    mgmt.Context.Projections = append(mgmt.Context.Projections, projection)
  }
  return mgmt
}

func (mgmt *SelectManager) Where(expr interface{}) *SelectManager {
  mgmt.Context.Wheres = append(mgmt.Context.Wheres, expr)
  return mgmt
}

func (mgmt *SelectManager) Offset(skip int) *SelectManager {
  mgmt.Tree.Offset.Expr = factory.Offset(skip)
  return mgmt
}

func (mgmt *SelectManager) Limit(take int) *SelectManager {
  mgmt.Tree.Limit = factory.Limit(take)
  return mgmt
}

func (mgmt *SelectManager) InnerJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, factory.InnerJoin(table.(Accessor).Relation()))
  case *nodes.Relation:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, factory.InnerJoin(table.(*nodes.Relation)))
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) OuterJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, factory.OuterJoin(table.(Accessor).Relation()))
  case *nodes.Relation:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, factory.OuterJoin(table.(*nodes.Relation)))
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) On(expr interface{}) *SelectManager {
  joins := mgmt.Context.Source.Right
  last := joins[len(joins)-1]

  switch last.(type) {
  case *nodes.InnerJoin:
    last.(*nodes.InnerJoin).Right = factory.On(expr)
  case *nodes.OuterJoin:
    last.(*nodes.OuterJoin).Right = factory.On(expr)
  default:
  }

  return mgmt
}

func (mgmt *SelectManager) SetEngine(engine interface{}) *SelectManager {
  mgmt.Engine = engine
  return mgmt
}

func (mgmt *SelectManager) ToSql() string {
  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }
  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}
