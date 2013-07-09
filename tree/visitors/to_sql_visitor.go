package visitors

import (
  "fmt"
  "github.com/chuckpreslar/codex/tree/nodes"
  "strings"
)

const (
  SPACE = ` `
  COMMA = `, `
  STAR  = `*`

  // Keywords
  SELECT = ` SELECT `
  FROM   = ` FROM `
  WHERE  = ` WHERE `
  LIMIT  = ` LIMIT `
  OFFSET = ` OFFSET `
  AND    = ` AND `
)

type ToSqlVisitor struct{}

func (sql *ToSqlVisitor) Accept(o interface{}) string {
  return sql.Visit(o, sql)
}

func (sql *ToSqlVisitor) Visit(o interface{}, visitor VisitorInterface) string {
  switch o.(type) {
  // Comparison visitors.
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
  // Begin SQL node visitors.
  case *nodes.RelationNode:
    return visitor.VisitRelation(o.(*nodes.RelationNode), visitor)
  case *nodes.AttributeNode:
    return visitor.VisitAttribute(o.(*nodes.AttributeNode), visitor)
  case *nodes.GroupingNode:
    return visitor.VisitGrouping(o.(*nodes.GroupingNode), visitor)
  case *nodes.NotNode:
    return visitor.VisitNot(o.(*nodes.NotNode), visitor)
  case *nodes.LiteralNode:
    return visitor.VisitLiteral(o.(*nodes.LiteralNode), visitor)
  case *nodes.InnerJoinNode:
    return visitor.VisitInnerJoin(o.(*nodes.InnerJoinNode), visitor)
  case *nodes.OuterJoinNode:
    return visitor.VisitOuterJoin(o.(*nodes.OuterJoinNode), visitor)
  case *nodes.OnNode:
    return visitor.VisitOn(o.(*nodes.OnNode), visitor)
  case *nodes.UnqualifiedColumnNode:
    return visitor.VisitUnqualifiedColumn(o.(*nodes.UnqualifiedColumnNode), visitor)
  case *nodes.LimitNode:
    return visitor.VisitLimit(o.(*nodes.LimitNode), visitor)
  case *nodes.OffsetNode:
    return visitor.VisitOffset(o.(*nodes.OffsetNode), visitor)
  case *nodes.JoinSourceNode:
    return visitor.VisitJoinSource(o.(*nodes.JoinSourceNode), visitor)
  case *nodes.SelectCoreNode:
    return visitor.VisitSelectCore(o.(*nodes.SelectCoreNode), visitor)
  case *nodes.SelectStatementNode:
    return visitor.VisitSelectStatement(o.(*nodes.SelectStatementNode), visitor)
  case *nodes.ValuesNode:
    return visitor.VisitValues(o.(*nodes.ValuesNode), visitor)
  case *nodes.InsertStatementNode:
    return visitor.VisitInsertStatement(o.(*nodes.InsertStatementNode), visitor)
  case *nodes.UpdateStatementNode:
    return visitor.VisitUpdateStatement(o.(*nodes.UpdateStatementNode), visitor)
  case *nodes.DeleteStatementNode:
    return visitor.VisitDeleteStatement(o.(*nodes.DeleteStatementNode), visitor)
  // SQL function visitors.
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
  // Standard type visitors.
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

func (sql *ToSqlVisitor) VisitAssignment(o *nodes.AssignmentNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// Being comparison node visitors.

func (sql *ToSqlVisitor) VisitEqual(o *nodes.EqualNode, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitNotEqual(o *nodes.NotEqualNode, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v != %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitGreaterThan(o *nodes.GreaterThanNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitGreaterThanOrEqual(o *nodes.GreaterThanOrEqualNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLessThan(o *nodes.LessThanNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLessThanOrEqual(o *nodes.LessThanOrEqualNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLike(o *nodes.LikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitUnlike(o *nodes.UnlikeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitOr(o *nodes.OrNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitAnd(o *nodes.AndNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// End comparison node visitors.

// Begin SQL node visitors.

func (sql *ToSqlVisitor) VisitRelation(o *nodes.RelationNode, visitor VisitorInterface) string {
  if 0 < len(o.Alias) {
    return visitor.QuoteTableName(o.Alias)
  }

  return visitor.QuoteTableName(o.Name)
}

func (sql *ToSqlVisitor) VisitAttribute(o *nodes.AttributeNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v.%v", visitor.Visit(o.Relation, visitor), visitor.QuoteColumnName(o.Name))
}

func (sql *ToSqlVisitor) VisitGrouping(o *nodes.GroupingNode, visitor VisitorInterface) string {
  return fmt.Sprintf("(%v)", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitNot(o *nodes.NotNode, visitor VisitorInterface) string {
  return fmt.Sprintf("NOT (%v)", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitLiteral(o *nodes.LiteralNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v", o.Expr)
}

func (sql *ToSqlVisitor) VisitInnerJoin(o *nodes.InnerJoinNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("INNER JOIN %v", visitor.Visit(o.Left, visitor))
  if nil != o.Right {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Right, visitor))
  }
  return str
}

func (sql *ToSqlVisitor) VisitOuterJoin(o *nodes.OuterJoinNode, visitor VisitorInterface) string {
  return fmt.Sprintf("LEFT OUTER JOIN %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitOn(o *nodes.OnNode, visitor VisitorInterface) string {
  return fmt.Sprintf("ON %v", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitUnqualifiedColumn(o *nodes.UnqualifiedColumnNode, visitor VisitorInterface) string {
  return visitor.QuoteColumnName(o.Expr)
}

func (sql *ToSqlVisitor) VisitLimit(o *nodes.LimitNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v%v", LIMIT, visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitOffset(o *nodes.OffsetNode, visitor VisitorInterface) string {
  return fmt.Sprintf("%v%v", OFFSET, visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitJoinSource(o *nodes.JoinSourceNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitSelectCore(o *nodes.SelectCoreNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("%v", SELECT)

  if length := len(o.Projections) - 1; 0 <= length {
    for index, projection := range o.Projections {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(projection, visitor))

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  } else {
    str = fmt.Sprintf("%v%v.%v", str, visitor.Visit(o.Relation, visitor), STAR)
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

  return strings.Trim(str, " ")
}

func (sql *ToSqlVisitor) VisitSelectStatement(o *nodes.SelectStatementNode, visitor VisitorInterface) string {
  str := ""

  for _, core := range o.Cores {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(core, visitor))
  }

  if nil != o.Limit {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(o.Limit, visitor))
  }

  if nil != o.Offset {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(o.Offset, visitor))
  }

  return str
}

func (sql *ToSqlVisitor) VisitValues(o *nodes.ValuesNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitInsertStatement(o *nodes.InsertStatementNode, visitor VisitorInterface) string {
  str := fmt.Sprintf("INSERT INTO %v%v", visitor.Visit(o.Relation, visitor), SPACE)

  if length := len(o.Columns) - 1; 0 <= length {
    str = fmt.Sprintf("%v(", str)
    for index, column := range o.Columns {

      switch column.(type) {
      case *nodes.LiteralNode:
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
  return str
}

func (sql *ToSqlVisitor) VisitUpdateStatement(o *nodes.UpdateStatementNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitDeleteStatement(o *nodes.DeleteStatementNode, visitor VisitorInterface) string {
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

// End SQL node visitors.

// Begin SQL Function visitors.

func (sql *ToSqlVisitor) VisitCount(o *nodes.CountNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitAverage(o *nodes.AverageNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitSum(o *nodes.SumNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitMaximum(o *nodes.MaximumNode, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitMinimum(o *nodes.MinimumNode, visitor VisitorInterface) string {
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

// End SQL Function visitors.

// Begin standard type visitors.

func (sql *ToSqlVisitor) VisitString(o interface{}) string {
  return fmt.Sprintf(`'%s'`, o)
}

func (sql *ToSqlVisitor) VisitInteger(o interface{}) string {
  return fmt.Sprintf(`%d`, o)
}

func (sql *ToSqlVisitor) VisitFloat(o interface{}) string {
  return strings.TrimRight(fmt.Sprintf(`%v`, o), `0`)
}

func (sql *ToSqlVisitor) VisitBool(o interface{}) string {
  return fmt.Sprintf(`'%t'`, o)
}

// End standard Type visitors.

// Begin helper visitors.

func (sql *ToSqlVisitor) QuoteTableName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

func (sql *ToSqlVisitor) QuoteColumnName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

// End helper visitors.
