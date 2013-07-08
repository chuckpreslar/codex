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
  case *nodes.Assignment:
    return visitor.VisitAssignment(o.(*nodes.Assignment), visitor)
  case *nodes.Equal:
    return visitor.VisitEqual(o.(*nodes.Equal), visitor)
  case *nodes.NotEqual:
    return visitor.VisitNotEqual(o.(*nodes.NotEqual), visitor)
  case *nodes.GreaterThan:
    return visitor.VisitGreaterThan(o.(*nodes.GreaterThan), visitor)
  case *nodes.GreaterThanOrEqual:
    return visitor.VisitGreaterThanOrEqual(o.(*nodes.GreaterThanOrEqual), visitor)
  case *nodes.LessThan:
    return visitor.VisitLessThan(o.(*nodes.LessThan), visitor)
  case *nodes.LessThanOrEqual:
    return visitor.VisitLessThanOrEqual(o.(*nodes.LessThanOrEqual), visitor)
  case *nodes.Like:
    return visitor.VisitLike(o.(*nodes.Like), visitor)
  case *nodes.Unlike:
    return visitor.VisitUnlike(o.(*nodes.Unlike), visitor)
  case *nodes.Or:
    return visitor.VisitOr(o.(*nodes.Or), visitor)
  case *nodes.And:
    return visitor.VisitAnd(o.(*nodes.And), visitor)
  // Begin SQL node visitors.
  case *nodes.Relation:
    return visitor.VisitRelation(o.(*nodes.Relation), visitor)
  case *nodes.Attribute:
    return visitor.VisitAttribute(o.(*nodes.Attribute), visitor)
  case *nodes.Grouping:
    return visitor.VisitGrouping(o.(*nodes.Grouping), visitor)
  case *nodes.Not:
    return visitor.VisitNot(o.(*nodes.Not), visitor)
  case *nodes.Literal:
    return visitor.VisitLiteral(o.(*nodes.Literal), visitor)
  case *nodes.InnerJoin:
    return visitor.VisitInnerJoin(o.(*nodes.InnerJoin), visitor)
  case *nodes.OuterJoin:
    return visitor.VisitOuterJoin(o.(*nodes.OuterJoin), visitor)
  case *nodes.On:
    return visitor.VisitOn(o.(*nodes.On), visitor)
  case *nodes.UnqualifiedColumn:
    return visitor.VisitUnqualifiedColumn(o.(*nodes.UnqualifiedColumn), visitor)
  case *nodes.Limit:
    return visitor.VisitLimit(o.(*nodes.Limit), visitor)
  case *nodes.Offset:
    return visitor.VisitOffset(o.(*nodes.Offset), visitor)
  case *nodes.JoinSource:
    return visitor.VisitJoinSource(o.(*nodes.JoinSource), visitor)
  case *nodes.SelectCore:
    return visitor.VisitSelectCore(o.(*nodes.SelectCore), visitor)
  case *nodes.SelectStatement:
    return visitor.VisitSelectStatement(o.(*nodes.SelectStatement), visitor)
  case *nodes.Values:
    return visitor.VisitValues(o.(*nodes.Values), visitor)
  case *nodes.InsertStatement:
    return visitor.VisitInsertStatement(o.(*nodes.InsertStatement), visitor)
  case *nodes.UpdateStatement:
    return visitor.VisitUpdateStatement(o.(*nodes.UpdateStatement), visitor)
  case *nodes.DeleteStatement:
    return visitor.VisitDeleteStatement(o.(*nodes.DeleteStatement), visitor)
  // SQL function visitors.
  case *nodes.Count:
    return visitor.VisitCount(o.(*nodes.Count), visitor)
  case *nodes.Average:
    return visitor.VisitAverage(o.(*nodes.Average), visitor)
  case *nodes.Sum:
    return visitor.VisitSum(o.(*nodes.Sum), visitor)
  case *nodes.Maximum:
    return visitor.VisitMaximum(o.(*nodes.Maximum), visitor)
  case *nodes.Minimum:
    return visitor.VisitMinimum(o.(*nodes.Minimum), visitor)
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

func (sql *ToSqlVisitor) VisitAssignment(o *nodes.Assignment, visitor VisitorInterface) string {
  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// Being comparison node visitors.

func (sql *ToSqlVisitor) VisitEqual(o *nodes.Equal, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitNotEqual(o *nodes.NotEqual, visitor VisitorInterface) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(o.Left, visitor))
  }

  return fmt.Sprintf("%v != %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitGreaterThan(o *nodes.GreaterThan, visitor VisitorInterface) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitGreaterThanOrEqual(o *nodes.GreaterThanOrEqual, visitor VisitorInterface) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLessThan(o *nodes.LessThan, visitor VisitorInterface) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLessThanOrEqual(o *nodes.LessThanOrEqual, visitor VisitorInterface) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitLike(o *nodes.Like, visitor VisitorInterface) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitUnlike(o *nodes.Unlike, visitor VisitorInterface) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitOr(o *nodes.Or, visitor VisitorInterface) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitAnd(o *nodes.And, visitor VisitorInterface) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// End comparison node visitors.

// Begin SQL node visitors.

func (sql *ToSqlVisitor) VisitRelation(o *nodes.Relation, visitor VisitorInterface) string {
  if 0 < len(o.Alias) {
    return visitor.QuoteTableName(o.Alias)
  }

  return visitor.QuoteTableName(o.Name)
}

func (sql *ToSqlVisitor) VisitAttribute(o *nodes.Attribute, visitor VisitorInterface) string {
  return fmt.Sprintf("%v.%v", visitor.Visit(o.Relation, visitor), visitor.QuoteColumnName(o.Name))
}

func (sql *ToSqlVisitor) VisitGrouping(o *nodes.Grouping, visitor VisitorInterface) string {
  return fmt.Sprintf("(%v)", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitNot(o *nodes.Not, visitor VisitorInterface) string {
  return fmt.Sprintf("NOT (%v)", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitLiteral(o *nodes.Literal, visitor VisitorInterface) string {
  return fmt.Sprintf("%v", o.Expr)
}

func (sql *ToSqlVisitor) VisitInnerJoin(o *nodes.InnerJoin, visitor VisitorInterface) string {
  str := fmt.Sprintf("INNER JOIN %v", visitor.Visit(o.Left, visitor))
  if nil != o.Right {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Right, visitor))
  }
  return str
}

func (sql *ToSqlVisitor) VisitOuterJoin(o *nodes.OuterJoin, visitor VisitorInterface) string {
  return fmt.Sprintf("LEFT OUTER JOIN %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

func (sql *ToSqlVisitor) VisitOn(o *nodes.On, visitor VisitorInterface) string {
  return fmt.Sprintf("ON %v", visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitUnqualifiedColumn(o *nodes.UnqualifiedColumn, visitor VisitorInterface) string {
  return visitor.QuoteColumnName(o.Expr)
}

func (sql *ToSqlVisitor) VisitLimit(o *nodes.Limit, visitor VisitorInterface) string {
  return fmt.Sprintf("%v%v", LIMIT, visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitOffset(o *nodes.Offset, visitor VisitorInterface) string {
  return fmt.Sprintf("%v%v", OFFSET, visitor.Visit(o.Expr, visitor))
}

func (sql *ToSqlVisitor) VisitJoinSource(o *nodes.JoinSource, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitSelectCore(o *nodes.SelectCore, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitSelectStatement(o *nodes.SelectStatement, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitValues(o *nodes.Values, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitInsertStatement(o *nodes.InsertStatement, visitor VisitorInterface) string {
  str := fmt.Sprintf("INSERT INTO %v%v", visitor.Visit(o.Relation, visitor), SPACE)

  if length := len(o.Columns) - 1; 0 <= length {
    str = fmt.Sprintf("%v(", str)
    for index, column := range o.Columns {

      switch column.(type) {
      case *nodes.Literal:
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

func (sql *ToSqlVisitor) VisitUpdateStatement(o *nodes.UpdateStatement, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitDeleteStatement(o *nodes.DeleteStatement, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitCount(o *nodes.Count, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitAverage(o *nodes.Average, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitSum(o *nodes.Sum, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitMaximum(o *nodes.Maximum, visitor VisitorInterface) string {
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

func (sql *ToSqlVisitor) VisitMinimum(o *nodes.Minimum, visitor VisitorInterface) string {
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
