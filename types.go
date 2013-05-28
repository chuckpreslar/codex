/*
 *	The MIT License (MIT)
 *
 *	Copyright (c) 2013 Chuck Preslar
 *
 *	Permission is hereby granted, free of charge, to any person obtaining a copy
 *	of this software and associated documentation files (the "Software"), to deal
 *	in the Software without restriction, including without limitation the rights
 *	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *	copies of the Software, and to permit persons to whom the Software is
 *	furnished to do so, subject to the following conditions:
 *
 *	The above copyright notice and this permission notice shall be included in
 *	all copies or substantial portions of the Software.
 *
 *	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *	THE SOFTWARE.
 */

package librarian

import (
	"fmt"
	"strings"
)

type accessor func(string) column
type column string
type columns []column
type expression string
type expressions []expression
type joinable []interface{}
type result map[string]interface{}
type results []result
type table string

func Table(t string) accessor {
	return func(c string) column {
		switch c {
		case ``:
			return column(fmt.Sprintf(`"%s"`, t))
		case `*`:
			return column(fmt.Sprintf(`"%s".%s`, t, c))
		default:
			return column(fmt.Sprintf(`"%s"."%s"`, t, c))
		}
	}
}

func (a accessor) On(s string) column {
	return column(fmt.Sprintf(`%s ON %s`, a(``), a(s)))
}

func (c column) Eq(v interface{}) expression {
	if nil == v {
		return expression(fmt.Sprintf(`%s IS NULL`, c))
	}
	return expression(fmt.Sprintf(`%s = %s`, c, formatExpressionValue(v)))
}

func (c column) Neq(v interface{}) expression {
	if nil == v {
		return expression(fmt.Sprintf(`%s IS NOT NULL`, c))
	}
	return expression(fmt.Sprintf(`%s != %s`, c, formatExpressionValue(v)))
}

func (c column) Gt(v interface{}) expression {
	return expression(fmt.Sprintf(`%s > %s`, c, formatExpressionValue(v)))
}

func (c column) Gte(v interface{}) expression {
	return expression(fmt.Sprintf(`%s >= %s`, c, formatExpressionValue(v)))
}

func (c column) Lt(v interface{}) expression {
	return expression(fmt.Sprintf(`%s < %s`, c, formatExpressionValue(v)))
}

func (c column) Lte(v interface{}) expression {
	return expression(fmt.Sprintf(`%s <= %s`, c, formatExpressionValue(v)))
}

func (c column) As(s string) column {
	return column(fmt.Sprintf(`%s AS "%s"`, c, s))
}

func (e expression) Or(ee expression) expression {
	return expression(fmt.Sprintf(`%s OR %s`, e, ee))
}

/**
 * Helper methods.
 */

func formatExpressionValue(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case column:
		return fmt.Sprintf("%s", v)
	default:
		return fmt.Sprintf("'%v'", v)
	}
}

func (c columns) join(d string) string {
	s := make([]string, len(c))
	for i, cc := range c {
		s[i] = string(cc)
	}
	return strings.Join(s, d)
}

func (e expressions) join(d string) string {
	s := make([]string, len(e))
	for i, ee := range e {
		s[i] = string(ee)
	}
	return strings.Join(s, d)
}

func (j joinable) join(d string) string {
	s := make([]string, len(j))
	for i, jj := range j {
		s[i] = fmt.Sprintf("%v", jj)
	}
	return strings.Join(s, d)
}
