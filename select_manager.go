package codex

import (
  "codex/tree/nodes"
  "codex/tree/visitors"
)

type SelectManager struct {
  Tree    *nodes.SelectStatement
  Context *nodes.SelectCore
}

func (mgmt *SelectManager) Project(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = &nodes.Literal{projection}
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
  if nil == mgmt.Tree.Offset {
    mgmt.Tree.Offset = &nodes.Offset{}
  }
  mgmt.Tree.Offset.Expr = skip
  return mgmt
}

func (mgmt *SelectManager) Limit(take int) *SelectManager {
  if nil == mgmt.Tree.Limit {
    mgmt.Tree.Limit = &nodes.Limit{}
  }
  mgmt.Tree.Limit.Expr = take
  return mgmt
}

func (mgmt *SelectManager) InnerJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.InnerJoin{table.(Accessor).Relation(), nil})
  case *nodes.Relation:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.InnerJoin{table, nil})
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) OuterJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.OuterJoin{table.(Accessor).Relation(), nil})
  case *nodes.Relation:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.OuterJoin{table, nil})
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) ToSql() string {
  visitor := &visitors.ToSqlVisitor{}
  return visitor.Accept(mgmt.Tree)
}

func (mgmt *SelectManager) On(expr interface{}) *SelectManager {
  joins := mgmt.Context.Source.Right
  last := joins[len(joins)-1]

  switch last.(type) {
  case *nodes.InnerJoin:
    last.(*nodes.InnerJoin).Right = &nodes.On{expr}
  case *nodes.OuterJoin:
    last.(*nodes.OuterJoin).Right = &nodes.On{expr}
  default:
  }

  return mgmt
}
