// Package visitors provides AST visitors for the codex package.
package visitors

import (
  "errors"
  "fmt"
  "strings"
)

import (
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

var DEBUG = false

const (
  SPACE = ` `
  COMMA = `, `
  STAR  = `*`

  // Keywords
  FROM     = ` FROM `
  WHERE    = ` WHERE `
  ORDER_BY = ` ORDER BY `
  GROUP_BY = ` GROUP BY `
  AND      = ` AND `
)

type ToSqlVisitor struct{}

func (self *ToSqlVisitor) Accept(o interface{}) (result string, err error) {
  defer func() {
    if r := recover(); r != nil {
      err = errors.New(fmt.Sprintf("%v", r))
    }
  }()

  result = self.Visit(o, self)

  return
}

func (self *ToSqlVisitor) Visit(o interface{}, visitor VisitorInterface) string {

  if DEBUG {
    fmt.Printf("DEBUG: Visiting %T\n", o)
  }

  switch o.(type) {
  // Unary node visitors.
  case *nodes.GroupingNode:
    return visitor.VisitGrouping(o.(*nodes.GroupingNode), visitor)
  case *nodes.NotNode:
    return visitor.VisitNot(o.(*nodes.NotNode), visitor)
  case *nodes.LiteralNode:
    return visitor.VisitLiteral(o.(*nodes.LiteralNode), visitor)
  case *nodes.OnNode:
    return visitor.VisitOn(o.(*nodes.OnNode), visitor)
  case *nodes.ColumnNode:
    return visitor.VisitColumn(o.(*nodes.ColumnNode), visitor)
  case *nodes.StarNode:
    return visitor.VisitStar(o.(*nodes.StarNode), visitor)
  case *nodes.BindingNode:
    return visitor.VisitBinding(o.(*nodes.BindingNode), visitor)
  case *nodes.UnqualifiedColumnNode:
    return visitor.VisitUnqualifiedColumn(o.(*nodes.UnqualifiedColumnNode), visitor)
  case *nodes.LimitNode:
    return visitor.VisitLimit(o.(*nodes.LimitNode), visitor)
  case *nodes.OffsetNode:
    return visitor.VisitOffset(o.(*nodes.OffsetNode), visitor)
  case *nodes.HavingNode:
    return visitor.VisitHaving(o.(*nodes.HavingNode), visitor)
  case *nodes.AscendingNode:
    return visitor.VisitAscending(o.(*nodes.AscendingNode), visitor)
  case *nodes.DescendingNode:
    return visitor.VisitDescending(o.(*nodes.DescendingNode), visitor)
  case *nodes.EngineNode:
    return visitor.VisitEngine(o.(*nodes.EngineNode), visitor)
  case *nodes.IndexNameNode:
    return visitor.VisitIndexName(o.(*nodes.IndexNameNode), visitor)

  // Binary node visitors.
  case *nodes.AssignmentNode:
    return visitor.VisitAssignment(o.(*nodes.AssignmentNode), visitor)
  case *nodes.EqualNode:
    return visitor.VisitEqual(o.(*nodes.EqualNode), visitor)
  case *nodes.NotEqualNode:
    return visitor.VisitNotEqual(o.(*nodes.NotEqualNode), visitor)
  case *nodes.GreaterThanNode:
    return visitor.VisitGreaterThan(o.(*nodes.GreaterThanNode), visitor)
  case *nodes.GreaterThanOrEqualNode:
    return visitor.VisitGreaterThanOrEqual(o.(*nodes.GreaterThanOrEqualNode), visitor)
  case *nodes.LessThanNode:
    return visitor.VisitLessThan(o.(*nodes.LessThanNode), visitor)
  case *nodes.LessThanOrEqualNode:
    return visitor.VisitLessThanOrEqual(o.(*nodes.LessThanOrEqualNode), visitor)
  case *nodes.LikeNode:
    return visitor.VisitLike(o.(*nodes.LikeNode), visitor)
  case *nodes.UnlikeNode:
    return visitor.VisitUnlike(o.(*nodes.UnlikeNode), visitor)
  case *nodes.OrNode:
    return visitor.VisitOr(o.(*nodes.OrNode), visitor)
  case *nodes.AndNode:
    return visitor.VisitAnd(o.(*nodes.AndNode), visitor)
  case *nodes.RelationNode:
    return visitor.VisitRelation(o.(*nodes.RelationNode), visitor)
  case *nodes.AttributeNode:
    return visitor.VisitAttribute(o.(*nodes.AttributeNode), visitor)
  case *nodes.InnerJoinNode:
    return visitor.VisitInnerJoin(o.(*nodes.InnerJoinNode), visitor)
  case *nodes.OuterJoinNode:
    return visitor.VisitOuterJoin(o.(*nodes.OuterJoinNode), visitor)
  case *nodes.JoinSourceNode:
    return visitor.VisitJoinSource(o.(*nodes.JoinSourceNode), visitor)
  case *nodes.ValuesNode:
    return visitor.VisitValues(o.(*nodes.ValuesNode), visitor)
  case *nodes.UnionNode:
    return visitor.VisitUnion(o.(*nodes.UnionNode), visitor)
  case *nodes.IntersectNode:
    return visitor.VisitIntersect(o.(*nodes.IntersectNode), visitor)
  case *nodes.ExceptNode:
    return visitor.VisitExcept(o.(*nodes.ExceptNode), visitor)
  case *nodes.UnexistingColumnNode:
    return visitor.VisitUnexistingColumn(o.(*nodes.UnexistingColumnNode), visitor)
  case *nodes.ExistingColumnNode:
    return visitor.VisitExistingColumn(o.(*nodes.ExistingColumnNode), visitor)

  // Nary node visitors.
  case *nodes.ConstraintNode:
    return visitor.VisitConstraint(o.(*nodes.ConstraintNode), visitor)
  case *nodes.NotNullNode:
    return visitor.VisitNotNull(o.(*nodes.NotNullNode), visitor)
  case *nodes.UniqueNode:
    return visitor.VisitUnique(o.(*nodes.UniqueNode), visitor)
  case *nodes.PrimaryKeyNode:
    return visitor.VisitPrimaryKey(o.(*nodes.PrimaryKeyNode), visitor)
  case *nodes.ForeignKeyNode:
    return visitor.VisitForeignKey(o.(*nodes.ForeignKeyNode), visitor)
  case *nodes.CheckNode:
    return visitor.VisitCheck(o.(*nodes.CheckNode), visitor)
  case *nodes.DefaultNode:
    return visitor.VisitDefault(o.(*nodes.DefaultNode), visitor)
  case *nodes.SelectCoreNode:
    return visitor.VisitSelectCore(o.(*nodes.SelectCoreNode), visitor)
  case *nodes.SelectStatementNode:
    return visitor.VisitSelectStatement(o.(*nodes.SelectStatementNode), visitor)
  case *nodes.InsertStatementNode:
    return visitor.VisitInsertStatement(o.(*nodes.InsertStatementNode), visitor)
  case *nodes.UpdateStatementNode:
    return visitor.VisitUpdateStatement(o.(*nodes.UpdateStatementNode), visitor)
  case *nodes.DeleteStatementNode:
    return visitor.VisitDeleteStatement(o.(*nodes.DeleteStatementNode), visitor)
  case *nodes.AlterStatementNode:
    return visitor.VisitAlterStatement(o.(*nodes.AlterStatementNode), visitor)

  // Function node visitors.
  case *nodes.CountNode:
    return visitor.VisitCount(o.(*nodes.CountNode), visitor)
  case *nodes.AverageNode:
    return visitor.VisitAverage(o.(*nodes.AverageNode), visitor)
  case *nodes.SumNode:
    return visitor.VisitSum(o.(*nodes.SumNode), visitor)
  case *nodes.MaximumNode:
    return visitor.VisitMaximum(o.(*nodes.MaximumNode), visitor)
  case *nodes.MinimumNode:
    return visitor.VisitMinimum(o.(*nodes.MinimumNode), visitor)

  // SQL constant visitors.
  case sql.Type:
    return visitor.VisitSqlType(o.(sql.Type), visitor)

  // Base visitors.
  case string:
    return visitor.VisitString(o)
  case int, int16, int32, int64:
    return visitor.VisitInteger(o)
  case float32, float64:
    return visitor.VisitFloat(o)
  case bool:
    return visitor.VisitBool(o)
  default:
    panic(fmt.Sprintf("No visitor method for <%T>.", o))
  }
}

// Begin Unary node visitors.

func (self *ToSqlVisitor) VisitGrouping(o *nodes.GroupingNode, visitor VisitorInterface) string {
  return fmt.Sprintf("(%v)", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitNot(o *nodes.NotNode, visitor VisitorInterface) string {
  return fmt.Sprintf("NOT (%v)", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitLiteral(o *nodes.LiteralNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v", o.Expr)
}

func (self *ToSqlVisitor) VisitOn(o *nodes.OnNode, visitor VisitorInterface) string {
  return fmt.Sprintf("ON %v", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitColumn(o *nodes.ColumnNode, visitor VisitorInterface) string {
  return visitor.QuoteColumnName(o.Expr)
}

func (self *ToSqlVisitor) VisitStar(o *nodes.StarNode, visitor VisitorInterface) string {
  return "*"
}

func (self *ToSqlVisitor) VisitBinding(o *nodes.BindingNode, visitor VisitorInterface) string {
  return "?"
}

func (self *ToSqlVisitor) VisitUnqualifiedColumn(o *nodes.UnqualifiedColumnNode, visitor VisitorInterface) string {
  return visitor.QuoteColumnName(o.Expr)
}

func (self *ToSqlVisitor) VisitLimit(o *nodes.LimitNode, visitor VisitorInterface) string {
  return fmt.Sprintf("LIMIT %v", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitOffset(o *nodes.OffsetNode, visitor VisitorInterface) string {
  return fmt.Sprintf("OFFSET %v", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitHaving(o *nodes.HavingNode, visitor VisitorInterface) string {
  return fmt.Sprintf("HAVING %v", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitAscending(o *nodes.AscendingNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v ASC", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitDescending(o *nodes.DescendingNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v DESC", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitEngine(o *nodes.EngineNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v", visitor.Visit(o.Expr, visitor))
}

func (self *ToSqlVisitor) VisitIndexName(o *nodes.IndexNameNode, visitor VisitorInterface) string {
  return fmt.Sprintf(`"%v"`, o.Expr)
}

// End Unary node visitors.

// Begin Binary node visitors.

func (self *ToSqlVisitor) VisitAssignment(o *nodes.AssignmentNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitEqual(o *nodes.EqualNode, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitNotEqual(o *nodes.NotEqualNode, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v != %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitGreaterThan(o *nodes.GreaterThanNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitGreaterThanOrEqual(o *nodes.GreaterThanOrEqualNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitLessThan(o *nodes.LessThanNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitLessThanOrEqual(o *nodes.LessThanOrEqualNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitLike(o *nodes.LikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitUnlike(o *nodes.UnlikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitOr(o *nodes.OrNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitAnd(o *nodes.AndNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitRelation(o *nodes.RelationNode, visitor VisitorInterface) string {
  if 0 < len(o.Alias) {
    return visitor.QuoteTableName(o.Alias)
  }

  return visitor.QuoteTableName(o.Name)
}

func (self *ToSqlVisitor) VisitAttribute(o *nodes.AttributeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v.%v", visitor.Visit(o.Relation, visitor), visitor.Visit(o.Name, visitor))
}

func (self *ToSqlVisitor) VisitInnerJoin(o *nodes.InnerJoinNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("INNER JOIN %v", visitor.Visit(o.Left, visitor))
  if nil != o.Right {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Right, visitor))
  }
  return str
}

func (self *ToSqlVisitor) VisitOuterJoin(o *nodes.OuterJoinNode, visitor VisitorInterface) string {
  return fmt.Sprintf("LEFT OUTER JOIN %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitJoinSource(o *nodes.JoinSourceNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("%v", visitor.Visit(o.Left, visitor))
  if length := len(o.Right) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, SPACE)
    for index, join := range o.Right {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(join, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, SPACE)
      }
    }
  }
  return strings.Trim(str, " ")
}

func (self *ToSqlVisitor) VisitValues(o *nodes.ValuesNode, visitor VisitorInterface) string {
  str := ""

  if length := len(o.Expressions) - 1; 0 <= length {
    str = fmt.Sprintf("%vVALUES (", str)
    for index, value := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(value, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  return str
}

func (self *ToSqlVisitor) VisitUnion(o *nodes.UnionNode, visitor VisitorInterface) string {
  return fmt.Sprintf("(%s UNION %s)", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitIntersect(o *nodes.IntersectNode, visitor VisitorInterface) string {
  return fmt.Sprintf("(%s INTERSECT %s)", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitExcept(o *nodes.ExceptNode, visitor VisitorInterface) string {
  return fmt.Sprintf("(%s EXCEPT %s)", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitUnexistingColumn(o *nodes.UnexistingColumnNode, visitor VisitorInterface) string {
  return fmt.Sprintf("ADD %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (self *ToSqlVisitor) VisitExistingColumn(o *nodes.ExistingColumnNode, visitor VisitorInterface) string {
  return fmt.Sprintf("ALTER COLUMN %v TYPE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// End Binary node visitors.

// Begin Nary node visitors.

func (self *ToSqlVisitor) VisitConstraint(o *nodes.ConstraintNode, visitor VisitorInterface) string {
  panic("VisitConstraint is unimplemented.")
}

func (self *ToSqlVisitor) VisitNotNull(o *nodes.NotNullNode, visitor VisitorInterface) string {
  return fmt.Sprintf("ALTER %v SET NOT NULL", visitor.Visit(o.Column, visitor))
}

func (self *ToSqlVisitor) VisitUnique(o *nodes.UniqueNode, visitor VisitorInterface) string {
  str := "ADD "
  // Optional index name.
  if 0 < len(o.Options) {
    expr := o.Options[0]
    if _, ok := expr.(string); ok {
      expr = nodes.IndexName(expr)
    }

    str = fmt.Sprintf("%vCONSTRAINT %v ", str, visitor.Visit(expr, visitor))
  }

  return fmt.Sprintf("%vUNIQUE(%v)", str, visitor.Visit(o.Column, visitor))
}

func (self *ToSqlVisitor) VisitPrimaryKey(o *nodes.PrimaryKeyNode, visitor VisitorInterface) string {
  str := "ADD "
  // Optional index name.
  if 0 < len(o.Options) {
    expr := o.Options[0]
    if _, ok := expr.(string); ok {
      expr = nodes.IndexName(expr)
    }

    str = fmt.Sprintf("%vCONSTRAINT %v ", str, visitor.Visit(expr, visitor))
  }

  return fmt.Sprintf("%vPRIMARY KEY(%v)", str, visitor.Visit(o.Column, visitor))
}

func (self *ToSqlVisitor) VisitForeignKey(o *nodes.ForeignKeyNode, visitor VisitorInterface) string {
  str := "ADD "
  // For FOREIGN KEY, index name is optional, REFERENCES is not.
  //
  // No index name ex.
  //
  //  CreateTable("orders").
  //    AddColumn("user_id").
  //    AddConstraint("user_id", FOREIGN_KEY, "users")
  //
  // With option index name ex.
  //
  //  CreateTable("orders").
  //    AddColumn("user_id").
  //    AddConstraint("user_id", FOREIGN_KEY, "users_fkey", "users")
  if 1 < len(o.Options) {
    expr := o.Options[0]
    if _, ok := expr.(string); ok {
      expr = nodes.IndexName(expr)
    }

    // Remove this item from the array, avoiding any potential memory leak.
    // https://code.google.com/p/go-wiki/wiki/SliceTricks
    length := len(o.Options) - 1
    copy(o.Options[0:], o.Options[1:])
    o.Options[length] = nil
    o.Options = o.Options[:length]

    str = fmt.Sprintf("%vCONSTRAINT %v%v", str, visitor.Visit(expr, visitor), SPACE)
  }

  str = fmt.Sprintf("%vFOREIGN KEY(%v)", str, visitor.Visit(o.Column, visitor))

  // If option is not here, user didn't do it right, but don't dereference and panic.
  if 0 < len(o.Options) {
    relation := o.Options[0]
    if _, ok := relation.(string); ok {
      relation = nodes.Relation(relation.(string))
    }

    str = fmt.Sprintf("%v%vREFERENCES %v", str, SPACE, visitor.Visit(relation, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitCheck(o *nodes.CheckNode, visitor VisitorInterface) string {
  panic("VisitCheck is unimplemented.")
}

func (self *ToSqlVisitor) VisitDefault(o *nodes.DefaultNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("ALTER %v SET DEFAULT", visitor.Visit(o.Column, visitor))

  if 0 < len(o.Options) {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Options[0], visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitSelectCore(o *nodes.SelectCoreNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("SELECT")

  if length := len(o.Projections) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, SPACE)
    for index, projection := range o.Projections {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(projection, visitor))

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  }

  str = fmt.Sprintf("%v%v%v", str, FROM, visitor.Visit(o.Source, visitor))

  if length := len(o.Wheres) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, WHERE)
    for index, where := range o.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(where, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  if length := len(o.Groups) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, GROUP_BY)
    for index, group := range o.Groups {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(group, visitor))

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  }

  if nil != o.Having {
    str = fmt.Sprintf("%v %v", str, visitor.Visit(o.Having, visitor))
  }

  return strings.Trim(str, " ")
}

func (self *ToSqlVisitor) VisitSelectStatement(o *nodes.SelectStatementNode, visitor VisitorInterface) string {
  str := ""

  if nil != o.Combinator {
    combinator := o.Combinator
    o.Combinator = nil
    return visitor.Visit(combinator, visitor)
  }

  for _, core := range o.Cores {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(core, visitor))
  }

  if length := len(o.Orders) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, ORDER_BY)
    for index, order := range o.Orders {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(order, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  }

  if nil != o.Limit {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Limit, visitor))
  }

  if nil != o.Offset {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Offset, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitInsertStatement(o *nodes.InsertStatementNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("INSERT INTO %v%v", visitor.Visit(o.Relation, visitor), SPACE)

  if length := len(o.Columns) - 1; 0 <= length {
    str = fmt.Sprintf("%v(", str)
    for index, column := range o.Columns {
      switch column.(type) {
      case *nodes.LiteralNode, *nodes.BindingNode:
        str = fmt.Sprintf("%v%v", str, visitor.Visit(column, visitor))
      default:
        str = fmt.Sprintf("%v%v", str, visitor.QuoteColumnName(column))
      }

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      } else {
        str = fmt.Sprintf("%v)%v", str, SPACE)
      }
    }
  }

  str = fmt.Sprintf("%v%v", str, visitor.Visit(o.Values, visitor))

  if nil != o.Returning {
    str = fmt.Sprintf("%v RETURNING %v", str, visitor.Visit(o.Returning, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitUpdateStatement(o *nodes.UpdateStatementNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("UPDATE %v%v", visitor.Visit(o.Relation, visitor), SPACE)

  if length := len(o.Values) - 1; 0 <= length {
    str = fmt.Sprintf("%vSET%v", str, SPACE)
    for index, assignment := range o.Values {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(assignment, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      } else {
        str = fmt.Sprintf("%v%v", str, SPACE)
      }
    }
  }

  if length := len(o.Wheres) - 1; 0 <= length {
    str = fmt.Sprintf("%vWHERE%v", str, SPACE)
    for index, filter := range o.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(filter, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  return str
}

func (self *ToSqlVisitor) VisitDeleteStatement(o *nodes.DeleteStatementNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("DELETE FROM %v%v", visitor.Visit(o.Relation, visitor), SPACE)

  if length := len(o.Wheres) - 1; 0 <= length {
    str = fmt.Sprintf("%vWHERE%v", str, SPACE)
    for index, filter := range o.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(filter, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  return str
}

func (self *ToSqlVisitor) VisitAlterStatement(o *nodes.AlterStatementNode, visitor VisitorInterface) string {
  str := ""

  if o.Drop {
    str = fmt.Sprintf("DROP TABLE IF EXISTS %v;\n", visitor.Visit(o.Relation, visitor))
  }

  if o.Create {
    str = fmt.Sprintf("CREATE TABLE %v ();\n", visitor.Visit(o.Relation, visitor))
  }

  columns := len(o.RemovedColumns) - 1

  for i, column := range o.RemovedColumns {
    str = fmt.Sprintf("%vALTER TABLE %v DROP COLUMN %v;", visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

    if i != columns {
      str = fmt.Sprintf("%v\n", str)
    }
  }

  indicies := len(o.RemovedIndicies) - 1

  for i, index := range o.RemovedIndicies {
    str = fmt.Sprintf("%vALTER TABLE %v DROP INDEX %v;", visitor.Visit(o.Relation, visitor), visitor.Visit(index, visitor))

    if i != indicies {
      str = fmt.Sprintf("%v\n", str)
    }
  }

  columns = len(o.UnexistingColumns) - 1

  for i, column := range o.UnexistingColumns {
    str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

    if i != columns {
      str = fmt.Sprintf("%v\n", str)
    }
  }

  if 0 <= columns {
    str = fmt.Sprintf("%v\n", str)
  }

  columns = len(o.ModifiedColumns) - 1

  for i, column := range o.ModifiedColumns {
    str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

    if i != columns {
      str = fmt.Sprintf("%v\n", str)
    }
  }

  if 0 <= columns {
    str = fmt.Sprintf("%v\n", str)
  }

  constraints := len(o.Constraints) - 1

  for i, constraint := range o.Constraints {
    str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(constraint, visitor))

    if i != constraints {
      str = fmt.Sprintf("%v\n", str)
    }
  }

  return str
}

// End Nary node visitors.

// Begin Function node visitors.

// End Function node visitors.

func (self *ToSqlVisitor) VisitCount(o *nodes.CountNode, visitor VisitorInterface) string {
  str := "COUNT("

  if o.Distinct {
    str = fmt.Sprintf("%vDISTINCT%v", str, SPACE)
  }

  if length := len(o.Expressions) - 1; 0 <= length {
    for index, expression := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(expression, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  if nil != o.Alias {
    str = fmt.Sprintf("%s AS %s", str, visitor.Visit(o.Alias, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitAverage(o *nodes.AverageNode, visitor VisitorInterface) string {
  str := "AVG("

  if o.Distinct {
    str = fmt.Sprintf("%vDISTINCT%v", str, SPACE)
  }

  if length := len(o.Expressions) - 1; 0 <= length {
    for index, expression := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(expression, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  if nil != o.Alias {
    str = fmt.Sprintf("%s AS %s", str, visitor.Visit(o.Alias, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitSum(o *nodes.SumNode, visitor VisitorInterface) string {
  str := "SUM("

  if o.Distinct {
    str = fmt.Sprintf("%vDISTINCT%v", str, SPACE)
  }

  if length := len(o.Expressions) - 1; 0 <= length {
    for index, expression := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(expression, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  if nil != o.Alias {
    str = fmt.Sprintf("%s AS %s", str, visitor.Visit(o.Alias, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitMaximum(o *nodes.MaximumNode, visitor VisitorInterface) string {
  str := "MAX("

  if o.Distinct {
    str = fmt.Sprintf("%vDISTINCT%v", str, SPACE)
  }

  if length := len(o.Expressions) - 1; 0 <= length {
    for index, expression := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(expression, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  if nil != o.Alias {
    str = fmt.Sprintf("%s AS %s", str, visitor.Visit(o.Alias, visitor))
  }

  return str
}

func (self *ToSqlVisitor) VisitMinimum(o *nodes.MinimumNode, visitor VisitorInterface) string {
  str := "MIN("

  if o.Distinct {
    str = fmt.Sprintf("%vDISTINCT%v", str, SPACE)
  }

  if length := len(o.Expressions) - 1; 0 <= length {
    for index, expression := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(expression, visitor))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  if nil != o.Alias {
    str = fmt.Sprintf("%s AS %s", str, visitor.Visit(o.Alias, visitor))
  }

  return str
}

// Begin SQL constant visitors.

func (self *ToSqlVisitor) VisitSqlType(o sql.Type, visitor VisitorInterface) string {
  switch o {
  case sql.STRING:
    return "varchar(255)"
  case sql.TEXT:
    return "text"
  case sql.BOOLEAN:
    return "boolean"
  case sql.INTEGER:
    return "integer"
  case sql.FLOAT:
    return "float"
  case sql.DECIMAL:
    return "decimal"
  case sql.DATE:
    return "date"
  case sql.TIME:
    return "time"
  case sql.DATETIME:
    return "datetime"
  case sql.TIMESTAMP:
    return "timestamp"
  default:
    panic("Unkown SQL Type constant.")
  }
}

// End SQL constant visitors.

// Begin Base visitors.

func (self *ToSqlVisitor) VisitString(o interface{}) string {
  return fmt.Sprintf(`'%s'`, o)
}

func (self *ToSqlVisitor) VisitInteger(o interface{}) string {
  return fmt.Sprintf(`%d`, o)
}

func (self *ToSqlVisitor) VisitFloat(o interface{}) string {
  return strings.TrimRight(fmt.Sprintf(`%v`, o), `0`)
}

func (self *ToSqlVisitor) VisitBool(o interface{}) string {
  return fmt.Sprintf(`'%t'`, o)
}

// End Base visitors.

// Begin Helpers.

func (self *ToSqlVisitor) QuoteTableName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

func (self *ToSqlVisitor) QuoteColumnName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

// End Helpers.
