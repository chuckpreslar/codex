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
	"database/sql"
	"fmt"
	"time"
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
		session     *sql.DB
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

func (s *SelectStatement) Run(session *sql.DB) *SelectStatement {
	s.session = session
	return s
}

func (s *SelectStatement) Count() (int64, error) {
	if s.session == nil {
		return -1, NoSessionError
	}
	s.projections = []column{column(fmt.Sprintf("COUNT(%s)", s.a("*")))}
	q := s.ToSQL()
	statment, err := s.session.Prepare(q)
	defer statment.Close()
	if err != nil {
		return -1, err
	}
	rows, err := statment.Query()
	if err != nil {
		return -1, err
	}
	var count int64
	rows.Next()
	defer logQueryInformation(time.Now(), q)
	err = rows.Scan(&count)
	return count, err
}

func (s *SelectStatement) First() (result, error) {
	if s.session == nil {
		return nil, NoSessionError
	}
	q := fmt.Sprintf("%s LIMIT 1", s.ToSQL())
	statment, err := s.session.Prepare(q)
	defer statment.Close()
	rows, err := statment.Query()
	if err != nil {
		return nil, err
	}
	col, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	defer logQueryInformation(time.Now(), q)
	var ptr = make([]interface{}, len(col))
	for i := 0; i < len(ptr); i++ {
		var buf interface{}
		ptr[i] = &buf
	}
	rows.Next()
	err = rows.Scan(ptr...)
	r := generateResult(col, ptr)
	return r, err
}

func (s *SelectStatement) All() (results, error) {
	if s.session == nil {
		return nil, NoSessionError
	}
	q := s.ToSQL()
	statment, err := s.session.Prepare(q)
	defer statment.Close()
	rows, err := statment.Query()
	if err != nil {
		return nil, err
	}
	col, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count, _ := s.Count()
	r := make(results, count)
	o := 0
	defer logQueryInformation(time.Now(), q)
	for rows.Next() {
		var ptr = make([]interface{}, len(col))
		for i := 0; i < len(ptr); i++ {
			var buf interface{}
			ptr[i] = &buf
		}
		err = rows.Scan(ptr...)
		r[o] = generateResult(col, ptr)
		o += 1
	}
	return r, err
}

func generateResult(c []string, p []interface{}) result {
	r := make(result)
	for i, v := range p {
		switch x := (*v.(*interface{})); x.(type) {
		case int64:
			p[i] = x.(int64)
		case time.Time:
			p[i] = x.(time.Time)
		case float64:
			p[i] = x.(float64)
		case bool:
			p[i] = x.(bool)
		case []uint8:
			p[i] = fmt.Sprintf("%s", x)
		default:
			p[i] = nil
		}
		r[c[i]] = p[i]
	}
	return r
}

func logQueryInformation(t time.Time, q string) {
	fmt.Printf("(%v) - %s\n", time.Now().Sub(t), q)
}
