package managers

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// SelectManager manages a tree that compiles to a SQL select statement.
type SelectManager struct {
  Tree    *nodes.SelectStatementNode // The AST for the SQL SELECT statement.
  Context *nodes.SelectCoreNode      // Reference to the Core the manager is curretly operating on.
  Engine  interface{}                // The SQL Engine.
}

// Appends a projection to the current Context's Projections slice,
// typically an AttributeNode or string.  If a string is provided, it is
// inserted as a LiteralNode.
func (mgmt *SelectManager) Project(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = nodes.Literal(projection)
    }
    mgmt.Context.Projections = append(mgmt.Context.Projections, projection)
  }
  return mgmt
}

// Appends an expression to the current Context's Wheres slice,
// typically a comparison, i.e. 1 = 1
func (mgmt *SelectManager) Where(expr interface{}) *SelectManager {
  mgmt.Context.Wheres = append(mgmt.Context.Wheres, expr)
  return mgmt
}

// Sets the Tree's Offset to the given integer.
func (mgmt *SelectManager) Offset(skip int) *SelectManager {
  mgmt.Tree.Offset = nodes.Offset(skip)
  return mgmt
}

// Sets the Tree's Limit to the given integer.
func (mgmt *SelectManager) Limit(take int) *SelectManager {
  mgmt.Tree.Limit = nodes.Limit(take)
  return mgmt
}

// Appends a new InnerJoin to the current Context's SourceNode.
func (mgmt *SelectManager) InnerJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.InnerJoin(table.(Accessor).Relation(), nil))
  case *nodes.RelationNode:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.InnerJoin(table.(*nodes.RelationNode), nil))
  }
  return mgmt
}

// Appends a new InnerJoin to the current Context's SourceNode.
func (mgmt *SelectManager) OuterJoin(table interface{}) *SelectManager {
  switch table.(type) {
  case Accessor:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.OuterJoin(table.(Accessor).Relation(), nil))
  case *nodes.RelationNode:
    mgmt.Context.Source.Right = append(mgmt.Context.Source.Right, nodes.OuterJoin(table.(*nodes.RelationNode), nil))
  }
  return mgmt
}

// Sets the last stored Join's Right leaf to a OnNode containing the
// given expression.
func (mgmt *SelectManager) On(expr interface{}) *SelectManager {
  joins := mgmt.Context.Source.Right

  if 0 == len(joins) {
    return mgmt
  }

  last := joins[len(joins)-1]

  switch last.(type) {
  case *nodes.InnerJoinNode:
    last.(*nodes.InnerJoinNode).Right = nodes.On(expr)
  case *nodes.OuterJoinNode:
    last.(*nodes.OuterJoinNode).Right = nodes.On(expr)
  }

  return mgmt
}

// Appends an expression to the current Context's Orders slice,
// typically an attribute.
func (mgmt *SelectManager) Order(expr interface{}) *SelectManager {
  mgmt.Tree.Orders = append(mgmt.Tree.Orders, expr)
  return mgmt
}

// Appends a node to the current Context's Groups slice,
// typically an attribute or column.
func (mgmt *SelectManager) Group(groupings ...interface{}) *SelectManager {
  for _, group := range groupings {
    switch group.(type) {
    case string:
      group = nodes.Literal(group)
    }
    mgmt.Context.Groups = append(mgmt.Context.Groups, group)
  }
  return mgmt
}

// Sets the Context's Having member to the given expression.
func (mgmt *SelectManager) Having(expr interface{}) *SelectManager {
  mgmt.Context.Having = nodes.Having(expr)
  return mgmt
}

// Sets the SQL Enginge.
func (mgmt *SelectManager) SetEngine(engine interface{}) *SelectManager {
  if _, ok := VISITORS[engine]; ok {
    mgmt.Engine = engine
  }
  return mgmt
}

// Calls a visitor's Accept method based on the manager's SQL Engine.
func (mgmt *SelectManager) ToSql() (string, error) {
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

// SelectManager factory method.
func Selection(relation *nodes.RelationNode) *SelectManager {
  selection := new(SelectManager)
  selection.Tree = nodes.SelectStatement(relation)
  selection.Context = selection.Tree.Cores[0]
  return selection
}
