package librarian

import (
	"fmt"
	"strings"
)

type column string
type columns []column
type table string
type expression string
type expressions []expression
type accessor func(string) column

func Table(t string) accessor {
	return func(c string) column {
		switch c {
		case "":
			return column(fmt.Sprintf("\"%s\"", t))
		case "*":
			return column(fmt.Sprintf("\"%s\".%s", t, c))
		default:
			return column(fmt.Sprintf("\"%s\".\"%s\"", t, c))
		}
	}
}

func (c column) Eq(v interface{}) expression {
	if v == nil {
		return expression(fmt.Sprintf("%s IS NULL", c))
	}
	return expression(fmt.Sprintf("%s = %s", c, format_expression_value(v)))
}

func (c column) Neq(v interface{}) expression {
	if v == nil {
		return expression(fmt.Sprintf("%s IS NOT NULL", c))
	}
	return expression(fmt.Sprintf("%s != %s", c, format_expression_value(v)))
}

func (c column) Gt(v interface{}) expression {
	return expression(fmt.Sprintf("%s > %s", c, format_expression_value(v)))
}

func (c column) Gte(v interface{}) expression {
	return expression(fmt.Sprintf("%s >= %s", c, format_expression_value(v)))
}

func (c column) Lt(v interface{}) expression {
	return expression(fmt.Sprintf("%s < %s", c, format_expression_value(v)))
}

func (c column) Lte(v interface{}) expression {
	return expression(fmt.Sprintf("%s <= %s", c, format_expression_value(v)))
}

func (e expression) Or(ee expression) expression {
	return expression(fmt.Sprintf("%s OR %s", e, ee))
}

/**
 * Helper methods.
 */

func format_expression_value(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	default:
		return fmt.Sprintf("'%s'", v)
	}
}

func (c columns) join(d string) string {
	s := []string{}
	for _, cc := range c {
		s = append(s, string(cc))
	}
	return strings.Join(s, d)
}

func (e expressions) join(d string) string {
	s := []string{}
	for _, ee := range e {
		s = append(s, string(ee))
	}
	return strings.Join(s, d)
}
