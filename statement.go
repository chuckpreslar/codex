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

func (s *SelectStatement) Select(c ...string) *SelectStatement {
	for _, cc := range c {
		s.projections = append(s.projections, s.a(cc))
	}
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
	return q
}
