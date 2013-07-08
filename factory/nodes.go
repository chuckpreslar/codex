package factory

import (
  "github.com/chuckpreslar/codex/tree/nodes"
)

// SQL Functions

// Creates a new Count SQL function node and returns it's pointer.
func Count(expr interface{}) *nodes.Count {
  count := new(nodes.Count)
  count.Expressions = append([]interface{}{}, expr)
  return count
}

// Creates a new Average SQL function node and returns it's pointer.
func Average(expr interface{}) *nodes.Average {
  average := new(nodes.Average)
  average.Expressions = append([]interface{}{}, expr)
  return average
}

// Creates a new Sum SQL function node and returns it's pointer.
func Sum(expr interface{}) *nodes.Sum {
  sum := new(nodes.Sum)
  sum.Expressions = append([]interface{}{}, expr)
  return sum
}

// Creates a new Minimum SQL function node and returns it's pointer.
func Minimum(expr interface{}) *nodes.Minimum {
  min := new(nodes.Minimum)
  min.Expressions = append([]interface{}{}, expr)
  return min
}

// Creates a new Maximum SQL function node and returns it's pointer.
func Maximum(expr interface{}) *nodes.Maximum {
  max := new(nodes.Maximum)
  max.Expressions = append([]interface{}{}, expr)
  return max
}

// Creates a new Not node and returns it's pointer.
func Not(expr interface{}) *nodes.Not {
  not := new(nodes.Not)
  not.Expr = expr
  return not
}

// Creates a new And node, wrapping it in a grouping and returning
// the groupings pointer.
func And(left, right interface{}) *nodes.Grouping {
  grouping := new(nodes.Grouping)
  and := new(nodes.And)
  and.Left = left
  and.Right = right
  grouping.Expr = and
  return grouping
}

// Creates a new Or node, wrapping it in a grouping and returning
// the groupings pointer.
func Or(left, right interface{}) *nodes.Grouping {
  grouping := new(nodes.Grouping)
  or := new(nodes.Or)
  or.Left = left
  or.Right = right
  grouping.Expr = or
  return grouping
}

// Creates a new As node returning it's pointer.
func As(left, right interface{}) *nodes.As {
  as := new(nodes.As)
  as.Left = left
  as.Right = right
  return as
}

// Creates a new Between node returning it's pointer.
func Between(left, right interface{}) *nodes.Between {
  between := new(nodes.Between)
  between.Left = left
  between.Right = right
  return between
}

// Creates a new InnerJoin node returning it's pointer.
func InnerJoin(left, right interface{}) *nodes.InnerJoin {
  join := new(nodes.InnerJoin)
  join.Left = left
  join.Right = right
  return join
}

// Creates a new OuterJoin node returning it's pointer.
func OuterJoin(left, right interface{}) *nodes.OuterJoin {
  join := new(nodes.OuterJoin)
  join.Left = left
  join.Right = right
  return join
}

// Creates a new Assignment node returning it's pointer.
func Assignment(left, right interface{}) *nodes.Assignment {
  assignment := new(nodes.Assignment)
  assignment.Left = left
  assignment.Right = right
  return assignment
}

// Creates a new Equal node returning it's pointer.
func Equal(left, right interface{}) *nodes.Equal {
  eq := new(nodes.Equal)
  eq.Left = left
  eq.Right = right
  return eq
}

// Creates a new NotEqual node returning it's pointer.
func NotEqual(left, right interface{}) *nodes.NotEqual {
  neq := new(nodes.NotEqual)
  neq.Left = left
  neq.Right = right
  return neq
}

// Creates a new GreaterThan node returning it's pointer.
func GreaterThan(left, right interface{}) *nodes.GreaterThan {
  gt := new(nodes.GreaterThan)
  gt.Left = left
  gt.Right = right
  return gt
}

// Creates a new GreaterThanOrEqual node returning it's pointer.
func GreaterThanOrEqual(left, right interface{}) *nodes.GreaterThanOrEqual {
  gte := new(nodes.GreaterThanOrEqual)
  gte.Left = left
  gte.Right = right
  return gte
}

// Creates a new LessThan node returning it's pointer.
func LessThan(left, right interface{}) *nodes.LessThan {
  lt := new(nodes.LessThan)
  lt.Left = left
  lt.Right = right
  return lt
}

// Creates a new LessThanOrEqual node returning it's pointer.
func LessThanOrEqual(left, right interface{}) *nodes.LessThanOrEqual {
  lte := new(nodes.LessThanOrEqual)
  lte.Left = left
  lte.Right = right
  return lte
}

// Creates a new Like node returning it's pointer.
func Like(left, right interface{}) *nodes.Like {
  like := new(nodes.Like)
  like.Left = left
  like.Right = right
  return like
}

// Creates a new Unlike node returning it's pointer.
func Unlike(left, right interface{}) *nodes.Unlike {
  unlike := new(nodes.Unlike)
  unlike.Left = left
  unlike.Right = right
  return unlike
}

// Creates a new Literal node returning it's pointer.
func Literal(expr interface{}) *nodes.Literal {
  literal := new(nodes.Literal)
  literal.Expr = expr
  return literal
}

// Creates a new On node returning it's pointer.
func On(expr interface{}) *nodes.On {
  on := new(nodes.On)
  on.Expr = expr
  return on
}

// Creates a new Limit node returning it's pointer.
func Limit(expr interface{}) *nodes.Limit {
  limit := new(nodes.Limit)
  limit.Expr = expr
  return limit
}

// Creates a new Offset node returning it's pointer.
func Offset(expr interface{}) *nodes.Offset {
  offset := new(nodes.Offset)
  offset.Expr = expr
  return offset
}

// Creates a new Group node returning it's pointer.
func Group(expr interface{}) *nodes.Group {
  group := new(nodes.Group)
  group.Expr = expr
  return group
}

// Creates a new Having node returning it's pointer.
func Having(expr interface{}) *nodes.Having {
  having := new(nodes.Having)
  having.Expr = expr
  return having
}

// Creates a new UnqualifiedColumn node returning it's pointer.
func UnqualifiedColumn(expr interface{}) *nodes.UnqualifiedColumn {
  column := new(nodes.UnqualifiedColumn)
  column.Expr = expr
  return column
}

// Creates a new JoinSource node returning it's pointer.
func JoinSource(relation *nodes.Relation) *nodes.JoinSource {
  source := new(nodes.JoinSource)
  source.Left = relation
  source.Right = make([]interface{}, 0)
  return source
}

// Creates a new SelectCore node returning it's pointer.
func SelectCore(relation *nodes.Relation) *nodes.SelectCore {
  core := new(nodes.SelectCore)
  core.Source = JoinSource(relation)
  core.Relation = relation
  core.Wheres = make([]interface{}, 0)
  core.Projections = make([]interface{}, 0)
  return core
}

// Creates a new SelectStatement node returning it's pointer.
func SelectStatement(relation *nodes.Relation) *nodes.SelectStatement {
  stmt := new(nodes.SelectStatement)
  stmt.Cores = []*nodes.SelectCore{SelectCore(relation)}
  return stmt
}

// Creates a new InsertStatement node returning it's pointer.
func InsertStatement(relation *nodes.Relation) *nodes.InsertStatement {
  stmt := new(nodes.InsertStatement)
  stmt.Relation = relation
  stmt.Columns = make([]interface{}, 0)
  stmt.Values = new(nodes.Values)
  return stmt
}

// Creates a new UpdateStatement node returning it's pointer.
func UpdateStatement(relation *nodes.Relation) *nodes.UpdateStatement {
  stmt := new(nodes.UpdateStatement)
  stmt.Relation = relation
  stmt.Values = make([]interface{}, 0)
  stmt.Wheres = make([]interface{}, 0)
  return stmt
}

// Creates a new DeleteStatement node returning it's pointer.
func DeleteStatement(relation *nodes.Relation) *nodes.DeleteStatement {
  stmt := new(nodes.DeleteStatement)
  stmt.Relation = relation
  stmt.Wheres = make([]interface{}, 0)
  return stmt
}
