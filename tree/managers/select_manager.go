package managers

import (
  "librarian/tree/nodes"
  "librarian/tree/visitors"
)

type SelectManager struct {
  Tree *nodes.SelectStatement
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
  mgmt.Tree.Offset.Expr = skip
  return mgmt
}

func (mgmt *SelectManager) Limit(take int) *SelectManager {
  mgmt.Tree.Limit.Expr = take
  return mgmt
}

func (mgmt *SelectManager) InnerJoin(relation *nodes.Relation) *SelectManager {
  mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.InnerJoin{relation, nil})
  return mgmt
}

func (mgmt *SelectManager) OuterJoin(relation *nodes.Relation) *SelectManager {
  mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, &nodes.OuterJoin{relation, nil})
  return mgmt
}

func (mgmt *SelectManager) ToSql() string {
  visitor := &visitors.ToSqlVisitor{}
  return visitor.Accept(mgmt.Tree)
}

func (mgmt *SelectManager) On(expr interface{}) *SelectManager {
  joins := mgmt.Context.Source.Right
  last := joins[len(joins) - 1]
  
  switch last.(type) {
  case *nodes.InnerJoin:
    join := last.(*nodes.InnerJoin)
    join.Right = &nodes.On{expr}
  case *nodes.OuterJoin:
    join := last.(*nodes.OuterJoin)
    join.Right = &nodes.On{expr}
  default:
  }

  return mgmt
}