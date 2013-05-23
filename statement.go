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
		limit       int
		count       bool
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

/**
 * Find will inititilze a Librarian SelectStatement with the assumption of
 * finding a record by an `id` column, returning a pointer to the
 * SelectStatement for continued chaning.
 *
 * Ex: SELECT "users".* FROM "users" WHERE "users"."id" = 3
 *	Search(users).Find(3).Run(session).Query() // #All or #First can be used here.
 *
 * @params int
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Find(i int) *SelectStatement {
	s.Where(s.a("id").Eq(i)).Limit(1)
	return s
}

/**
 * Select allows for SQL projections, accepting an interface type to
 * allow for type switching.  Actual accepted types are strings that
 * will be used to generate columns with the statements accessor,
 * or columns generated previously, returning a pointer to the
 * SelectStatement for continued chaning.
 *
 * Ex: SELECT "users"."id", "users"."email" FROM "users"
 *	Search(users).Select("id", "email").Run(session).Query() // #All or #First can be used here.
 *
 * @params ...interface{}
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Select(c ...interface{}) *SelectStatement {
	for _, cc := range c {
		switch cc.(type) {
		case column:
			s.projections = append(s.projections, cc.(column))
		case string:
			s.projections = append(s.projections, s.a(cc.(string)))
		default:
			panic(BadArgsError)
		}
	}
	return s
}

/**
 * Where provides filtering options for Librarian SelectStatements,
 * taking an `expression` type as a parameter and returning a pointer to the
 * SelectStatement for continued chaning.
 *
 * Ex. SELECT "users".* FROM "users" WHERE "users"."email" = 'test@example.com'
 *	Search(users).Where(users("email").Eq("test@example.com")).Run(session).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

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

func (s *SelectStatement) Limit(l int) *SelectStatement {
	s.limit = l
	return s
}

func (s *SelectStatement) ToSQL() string {
	q := "SELECT"
	if 0 == len(s.projections) {
		q += fmt.Sprintf(" %s ", s.a("*"))
	} else {
		q += fmt.Sprintf(" %s ", s.projections.join(", "))
	}
	q += fmt.Sprintf("FROM %s ", s.reference)
	if 0 < len(s.joins) {
		q += fmt.Sprintf("%s ", s.joins.join(" "))
	}
	if 0 < len(s.filters) {
		q += fmt.Sprintf("WHERE %s ", s.filters.join(" AND "))
	}
	if 0 != s.limit {
		q += fmt.Sprintf("LIMIT %d", s.limit)
	}
	return q
}

func (s *SelectStatement) Run(session *sql.DB) *SelectStatement {
	s.session = session
	return s
}

func (s *SelectStatement) Count() (int, error) {
	s.count = true
	s.projections = []column{column(fmt.Sprintf("COUNT(%s)", s.a("*")))}
	result, err := s.process()
	if 1 == len(result) && nil == err {
		count, ok := result[0]["count"]
		if ok {
			return int(count.(int64)), err
		} else {
			return -1, BadResultsError
		}
	}
	return -1, BadResultsError
}

func (s *SelectStatement) First() (result, error) {
	s.limit = 1
	result, err := s.process()
	if 1 > len(result) {
		return nil, err
	}
	return result[0], err
}

func (s *SelectStatement) Query() (results, error) {
	return s.process()
}

func (s *SelectStatement) clone() *SelectStatement {
	clone := &SelectStatement{
		a:           s.a,
		projections: s.projections,
		reference:   s.reference,
		filters:     s.filters,
		joins:       s.joins,
		session:     s.session,
		limit:       s.limit,
	}
	return clone
}

func (s *SelectStatement) process() (results, error) {
	if nil == s.session {
		return nil, NoSessionError
	}
	if 0 == s.limit && !s.count {
		sqlExpectedRowCount, _ := s.clone().Count()
		s.limit = int(sqlExpectedRowCount)
	}
	sqlQuery := s.ToSQL()
	sqlStatment, err := s.session.Prepare(sqlQuery)
	defer sqlStatment.Close()
	sqlRows, err := sqlStatment.Query()
	if nil != err {
		return nil, err
	}
	sqlColumns, err := sqlRows.Columns()
	if err != nil {
		return nil, err
	}
	sqlReturnSize := s.limit | 1
	sqlResultsArray := make(results, sqlReturnSize)
	sqlCurrentResultIndex := 0
	defer logQueryInformation(time.Now(), sqlQuery)
	for sqlRows.Next() {
		var sqlResultsBuffer = generateResultsBuffer(len(sqlColumns))
		err = sqlRows.Scan(sqlResultsBuffer...)
		sqlResultsArray[sqlCurrentResultIndex] = generateResultMap(sqlColumns, sqlResultsBuffer)
		sqlCurrentResultIndex += 1
	}
	return sqlResultsArray, err
}

func generateResultMap(c []string, p []interface{}) result {
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

func generateResultsBuffer(l int) []interface{} {
	p := make([]interface{}, l)
	for i := 0; i < len(p); i++ {
		var buf interface{}
		p[i] = &buf
	}
	return p
}

func logQueryInformation(t time.Time, q string) {
	fmt.Printf("(%v) - %s\n", time.Now().Sub(t), q)
}
