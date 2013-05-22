package librarian

import (
	"fmt"
)

type (
	Statement interface {
		ToSQL() string
	}
)

type (
	SelectStatement struct {
		a           accessor
		projections columns
		reference   table
		filters     expressions
		joins       expressions
	}

	InsertStatement struct {
		a accessor
	}
	UpdateStatement struct {
		a accessor
	}
	DeleteStatement struct {
		a accessor
	}
)

func (s *SelectStatement) Select(c ...interface{}) *SelectStatement {
	for _, cc := range c {
		switch cc.(type) {
		case column:
			s.projections = append(s.projections, cc.(column))
		default:
			s.projections = append(s.projections, s.a(cc.(string)))
		}
	}
	return s
}

func (s *SelectStatement) Where(e expression) *SelectStatement {
	s.filters = append(s.filters, e)
	return s
}

func (s *SelectStatement) Join(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("JOIN %s", e)))
	return s
}

func (s *SelectStatement) InnerJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("INNER JOIN %s", e)))
	return s
}

func (s *SelectStatement) OuterJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("OUTER JOIN %s", e)))
	return s
}

func (s *SelectStatement) LeftJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("LEFT JOIN %s", e)))
	return s
}

func (s *SelectStatement) RightJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("RIGHT JOIN %s", e)))
	return s
}

func (s *SelectStatement) FullJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("FULL JOIN %s", e)))
	return s
}

func (s *SelectStatement) ToSQL() string {
	q := "SELECT"
	if len(s.projections) == 0 {
		q += fmt.Sprintf(" %s ", s.a("*"))
	} else {
		q += fmt.Sprintf(" %s ", s.projections.join(", "))
	}
	q += fmt.Sprintf("FROM %s ", s.reference)
	if len(s.joins) > 0 {
		q += fmt.Sprintf("%s ", s.joins.join(" "))
	}
	if len(s.filters) > 0 {
		q += fmt.Sprintf("WHERE %s ", s.filters.join(" AND "))
	}
	return q
}
