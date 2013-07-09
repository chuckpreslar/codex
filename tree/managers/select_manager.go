package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

type SelectManager struct {
  Tree    *nodes.SelectStatementNode
  Context *nodes.SelectCoreNode
  Engine  interface{}
}

func (mgmt *SelectManager) Project(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = nodes.Literal(projection)
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
  mgmt.Tree.Offset.Expr = nodes.Offset(skip)
  return mgmt
}

func (mgmt *SelectManager) Limit(take int) *SelectManager {
  mgmt.Tree.Limit = nodes.Limit(take)
  return mgmt
}

func (mgmt *SelectManager) InnerJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.InnerJoin(table.(Accessor).Relation(), nil))
  case *nodes.RelationNode:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.InnerJoin(table.(*nodes.RelationNode), nil))
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) OuterJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.OuterJoin(table.(Accessor).Relation(), nil))
  case *nodes.RelationNode:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.OuterJoin(table.(*nodes.RelationNode), nil))
  default:
    panic("Cannot join unknown type.")
  }
  return mgmt
}

func (mgmt *SelectManager) On(expr interface{}) *SelectManager {
  joins := mgmt.Context.Source.Right
  last := joins[len(joins)-1]

  switch last.(type) {
  case *nodes.InnerJoinNode:
    last.(*nodes.InnerJoinNode).Right = nodes.On(expr)
  case *nodes.OuterJoinNode:
    last.(*nodes.OuterJoinNode).Right = nodes.On(expr)
  default:
  }

  return mgmt
}

func (mgmt *SelectManager) SetEngine(engine interface{}) *SelectManager {
  mgmt.Engine = engine
  return mgmt
}

func (mgmt *SelectManager) ToSql() string {
  for _, core := range mgmt.Tree.Cores {
    if 0 == len(core.Projections) {
      core.Projections = append(core.Projections, nodes.Attribute(nodes.Star(), core.Relation))
    }
  }

  if nil == mgmt.Engine {
    mgmt.Engine = "to_sql"
  }

  return VISITORS[mgmt.Engine].Accept(mgmt.Tree)
}
