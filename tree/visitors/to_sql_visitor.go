package visitors

import (
  "codex/tree/nodes"
  "fmt"
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

func (visitor *ToSqlVisitor) Accept(o interface{}) string {
  return visitor.Visit(o)
}

func (visitor *ToSqlVisitor) Visit(o interface{}) string {
  switch o.(type) {
  // Comparison visitors.
  case *nodes.Assignment:
    return visitor.VisitAssignment(o.(*nodes.Assignment))
  case *nodes.Equal:
    return visitor.VisitEqual(o.(*nodes.Equal))
  case *nodes.NotEqual:
    return visitor.VisitNotEqual(o.(*nodes.NotEqual))
  case *nodes.GreaterThan:
    return visitor.VisitGreaterThan(o.(*nodes.GreaterThan))
  case *nodes.GreaterThanOrEqual:
    return visitor.VisitGreaterThanOrEqual(o.(*nodes.GreaterThanOrEqual))
  case *nodes.LessThan:
    return visitor.VisitLessThan(o.(*nodes.LessThan))
  case *nodes.LessThanOrEqual:
    return visitor.VisitLessThanOrEqual(o.(*nodes.LessThanOrEqual))
  case *nodes.Like:
    return visitor.VisitLike(o.(*nodes.Like))
  case *nodes.Unlike:
    return visitor.VisitUnlike(o.(*nodes.Unlike))
  case *nodes.Or:
    return visitor.VisitOr(o.(*nodes.Or))
  case *nodes.And:
    return visitor.VisitAnd(o.(*nodes.And))
  // Begin SQL node visitors.
  case *nodes.Relation:
    return visitor.VisitRelation(o.(*nodes.Relation))
  case *nodes.Attribute:
    return visitor.VisitAttribute(o.(*nodes.Attribute))
  case *nodes.Grouping:
    return visitor.VisitGrouping(o.(*nodes.Grouping))
  case *nodes.Not:
    return visitor.VisitNot(o.(*nodes.Not))
  case *nodes.Literal:
    return visitor.VisitLiteral(o.(*nodes.Literal))
  case *nodes.InnerJoin:
    return visitor.VisitInnerJoin(o.(*nodes.InnerJoin))
  case *nodes.OuterJoin:
    return visitor.VisitOuterJoin(o.(*nodes.OuterJoin))
  case *nodes.On:
    return visitor.VisitOn(o.(*nodes.On))
  case *nodes.UnqualifiedColumn:
    return visitor.VisitUnqualifiedColumn(o.(*nodes.UnqualifiedColumn))
  case *nodes.Limit:
    return visitor.VisitLimit(o.(*nodes.Limit))
  case *nodes.Offset:
    return visitor.VisitOffset(o.(*nodes.Offset))
  case *nodes.JoinSource:
    return visitor.VisitJoinSource(o.(*nodes.JoinSource))
  case *nodes.SelectCore:
    return visitor.VisitSelectCore(o.(*nodes.SelectCore))
  case *nodes.SelectStatement:
    return visitor.VisitSelectStatement(o.(*nodes.SelectStatement))
  case *nodes.Values:
    return visitor.VisitValues(o.(*nodes.Values))
  case *nodes.InsertStatement:
    return visitor.VisitInsertStatement(o.(*nodes.InsertStatement))
  case *nodes.UpdateStatement:
    return visitor.VisitUpdateStatement(o.(*nodes.UpdateStatement))
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

func (visitor *ToSqlVisitor) VisitAssignment(o *nodes.Assignment) string {
  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

// Being comparison node visitors.

func (visitor *ToSqlVisitor) VisitEqual(o *nodes.Equal) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NULL", visitor.Visit(o.Left))
  }

  return fmt.Sprintf("%v = %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitNotEqual(o *nodes.NotEqual) string {
  if nil == o.Right {
    return fmt.Sprintf("%v IS NOT NULL", visitor.Visit(o.Left))
  }

  return fmt.Sprintf("%v != %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThan(o *nodes.GreaterThan) string {
  return fmt.Sprintf("%v > %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitGreaterThanOrEqual(o *nodes.GreaterThanOrEqual) string {
  return fmt.Sprintf("%v >= %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLessThan(o *nodes.LessThan) string {
  return fmt.Sprintf("%v < %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLessThanOrEqual(o *nodes.LessThanOrEqual) string {
  return fmt.Sprintf("%v <= %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitLike(o *nodes.Like) string {
  return fmt.Sprintf("%v LIKE %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitUnlike(o *nodes.Unlike) string {
  return fmt.Sprintf("%v NOT LIKE %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitOr(o *nodes.Or) string {
  return fmt.Sprintf("%v OR %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitAnd(o *nodes.And) string {
  return fmt.Sprintf("%v AND %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

// End comparison node visitors.

// Begin SQL node visitors.

func (visitor *ToSqlVisitor) VisitRelation(o *nodes.Relation) string {
  if 0 < len(o.Alias) {
    return visitor.QuoteTableName(o.Alias)
  }

  return visitor.QuoteTableName(o.Name)
}

func (visitor *ToSqlVisitor) VisitAttribute(o *nodes.Attribute) string {
  return fmt.Sprintf("%v.%v", visitor.Visit(o.Relation), visitor.QuoteColumnName(o.Name))
}

func (visitor *ToSqlVisitor) VisitGrouping(o *nodes.Grouping) string {
  return fmt.Sprintf("(%v)", visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitNot(o *nodes.Not) string {
  return fmt.Sprintf("NOT (%v)", visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitLiteral(o *nodes.Literal) string {
  return fmt.Sprintf("%v", o.Expr)
}

func (visitor *ToSqlVisitor) VisitInnerJoin(o *nodes.InnerJoin) string {
  str := fmt.Sprintf("INNER JOIN %v", visitor.Visit(o.Left))
  if nil != o.Right {
    str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Right))
  }
  return str
}

func (visitor *ToSqlVisitor) VisitOuterJoin(o *nodes.OuterJoin) string {
  return fmt.Sprintf("LEFT OUTER JOIN %v %v", visitor.Visit(o.Left), visitor.Visit(o.Right))
}

func (visitor *ToSqlVisitor) VisitOn(o *nodes.On) string {
  return fmt.Sprintf("ON %v", visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitUnqualifiedColumn(o *nodes.UnqualifiedColumn) string {
  return visitor.QuoteColumnName(o.Expr)
}

func (visitor *ToSqlVisitor) VisitLimit(o *nodes.Limit) string {
  return fmt.Sprintf("%v%v", LIMIT, visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitOffset(o *nodes.Offset) string {
  return fmt.Sprintf("%v%v", OFFSET, visitor.Visit(o.Expr))
}

func (visitor *ToSqlVisitor) VisitJoinSource(o *nodes.JoinSource) string {
  str := fmt.Sprintf("%v", visitor.Visit(o.Left))
  if length := len(o.Right) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, SPACE)
    for index, join := range o.Right {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(join))
      if index != length {
        str = fmt.Sprintf("%v%v", str, SPACE)
      }
    }
  }
  return strings.Trim(str, " ")
}

func (visitor *ToSqlVisitor) VisitSelectCore(o *nodes.SelectCore) string {
  str := fmt.Sprintf("%v", SELECT)

  if length := len(o.Projections) - 1; 0 <= length {
    for index, projection := range o.Projections {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(projection))

      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
  } else {
    str = fmt.Sprintf("%v%v.%v", str, visitor.Visit(o.Relation), STAR)
  }

  str = fmt.Sprintf("%v%v%v", str, FROM, visitor.Visit(o.Source))

  if length := len(o.Wheres) - 1; 0 <= length {
    str = fmt.Sprintf("%v%v", str, WHERE)
    for index, where := range o.Wheres {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(where))
      if index != length {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  return strings.Trim(str, " ")
}

func (visitor *ToSqlVisitor) VisitSelectStatement(o *nodes.SelectStatement) string {
  str := ""

  for _, core := range o.Cores {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(core))
  }

  if nil != o.Limit {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(o.Limit))
  }

  if nil != o.Offset {
    str = fmt.Sprintf("%v%v", str, visitor.Visit(o.Offset))
  }

  return str
}

func (visitor *ToSqlVisitor) VisitValues(o *nodes.Values) string {
  str := ""

  if length := len(o.Expressions) - 1; 0 <= length {
    str = fmt.Sprintf("%vVALUES (", str)
    for index, value := range o.Expressions {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(value))
      if index != length {
        str = fmt.Sprintf("%v%v", str, COMMA)
      }
    }
    str = fmt.Sprintf("%v)", str)
  }

  return str
}

func (visitor *ToSqlVisitor) VisitInsertStatement(o *nodes.InsertStatement) string {
  str := fmt.Sprintf("INSERT INTO %v%v", visitor.Visit(o.Relation), SPACE)

  if length := len(o.Columns) - 1; 0 <= length {
    str = fmt.Sprintf("%v(", str)
    for index, column := range o.Columns {

      switch column.(type) {
      case *nodes.Literal:
        str = fmt.Sprintf("%v%v", str, visitor.Visit(column))
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

  str = fmt.Sprintf("%v%v", str, visitor.VisitValues(o.Values))
  return str
}

func (visitor *ToSqlVisitor) VisitUpdateStatement(o *nodes.UpdateStatement) string {
  // "UPDATE #{visit o.relation, a}",
  // ("SET #{o.values.map { |value| visit value, a }.join ', '}" unless o.values.empty?),
  // ("WHERE #{wheres.map { |x| visit x, a }.join ' AND '}" unless wheres.empty?),
  str := fmt.Sprintf("UPDATE %v%v", visitor.Visit(o.Relation), SPACE)

  if length := len(o.Values) - 1; 0 <= length {
    str = fmt.Sprintf("%vSET%v", str, SPACE)
    for index, assignment := range o.Values {
      str = fmt.Sprintf("%v%v", str, visitor.Visit(assignment))
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
      str = fmt.Sprintf("%v%v", str, visitor.Visit(filter))
      if index != length {
        str = fmt.Sprintf("%v%v", str, AND)
      }
    }
  }

  return str
}

// End SQL node visitors.

// Begin standard type visitors.

func (visitor *ToSqlVisitor) VisitString(o interface{}) string {
  return fmt.Sprintf(`'%s'`, o)
}

func (visitor *ToSqlVisitor) VisitInteger(o interface{}) string {
  return fmt.Sprintf(`%d`, o)
}

func (visitor *ToSqlVisitor) VisitFloat(o interface{}) string {
  return strings.TrimRight(fmt.Sprintf(`%v`, o), `0`)
}

func (visitor *ToSqlVisitor) VisitBool(o interface{}) string {
  return fmt.Sprintf(`'%t'`, o)
}

// End standard Type visitors.

// Begin helper visitors.

func (visitor *ToSqlVisitor) QuoteTableName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

func (visitor *ToSqlVisitor) QuoteColumnName(o interface{}) string {
  return fmt.Sprintf(`"%v"`, o)
}

// End helper visitors.
