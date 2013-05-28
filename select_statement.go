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
	"strings"
	"time"
)

type SelectStatement struct {
	a           accessor
	projections columns
	reference   table
	filters     expressions
	joins       expressions
	ordering    joinable
	session     *sql.DB
	limit       int
	offset      int
	count       bool
}

/**
 * `Select` allows for SQL projections, accepting an `interface` type to
 * allow for type switching.  Actual accepted types are strings that
 * will be used to generate `columns` with the statements `accessor`,
 * or `columns` generated previously, returning a pointer to the
 * SelectStatement for continued chaning.
 *
 * Ex: SELECT "users"."id", "users"."email" FROM "users"
 *	With(session).Search(users).Select("id", "email").Query() // #Query or #First can be used here.
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
 * `Where` provides filtering options for Librarian `SelectStatements`,
 * taking an `expression` type as a parameter and returning a pointer to the
 * `SelectStatement` for continued chaning.
 *
 * Ex. SELECT "users".* FROM "users" WHERE "users"."email" = 'test@example.com'
 *	With(session).Search(users).Where(users("email").Eq("test@example.com")).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Where(e expression) *SelectStatement {
	s.filters = append(s.filters, e)
	return s
}

/**
 * `Join` takes an `expression` like `Where`, except the expression is generated
 * trough a chained call to `On` which operates on an `accessor`. A pointer to
 * the `SelectStatement` is returned for chaining.
 *
 * Ex. SELECT "users".* FROM "users" JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).Join(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Join(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("JOIN %s", e)))
	return s
}

/**
 * Refer to `Join`
 *
 * Ex. SELECT "users".* FROM "users" INNER JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).InnerJoin(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) InnerJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("INNER JOIN %s", e)))
	return s
}

/**
 * Refer to `Join`
 *
 * Ex. SELECT "users".* FROM "users" OUTER JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).OuterJoin(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) OuterJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("OUTER JOIN %s", e)))
	return s
}

/**
 * Refer to `Join`
 *
 * Ex. SELECT "users".* FROM "users" LEFT JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).LeftJoin(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) LeftJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("LEFT JOIN %s", e)))
	return s
}

/**
 * Refer to `Join`
 *
 * Ex. SELECT "users".* FROM "users" RIGHT JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).Right(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) RightJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("RIGHT JOIN %s", e)))
	return s
}

/**
 * Refer to `Join`
 *
 * Ex. SELECT "users".* FROM "users" FULL JOIN "orders" ON "orders"."user_id" = "user"."id"
 *	With(session).Search(users).FullJoin(orders.On("user_id").Eq(users("id"))).Query()
 *
 * @params expression
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) FullJoin(e expression) *SelectStatement {
	s.joins = append(s.joins, expression(fmt.Sprintf("FULL JOIN %s", e)))
	return s
}

/**
 * `Limit` restricts the number of results returned from
 * executing the query.
 *
 * Ex. SELECT "users".* FROM "users" LIMIT 1
 *	With(session).Search(users).Limit(1).Query()
 *
 * @params int
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Limit(l int) *SelectStatement {
	s.limit = l
	return s
}

/**
 * `Offset` skips a the number of records given by the function parameter.
 *
 * Ex. SELECT "users".* FROM "users" OFFSET 1
 *	With(session).Search(users).Offset(1).Query()
 *
 * @params int
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Offset(o int) *SelectStatement {
	s.offset = o
	return s
}

func (s *SelectStatement) Order(c interface{}, d string) *SelectStatement {
	d = strings.ToUpper(d)
	switch c.(type) {
	case column:
		s.ordering = append(s.ordering, fmt.Sprintf("%s %s", c.(column), d))
	case string:
		s.ordering = append(s.ordering, fmt.Sprintf("%s %s", s.a(c.(string)), d))
	default:
		panic(BadArgsError)
	}
	return s
}

/**
 * `ToSQL` parses the current `SelectStatement`, returning it as a string
 * in SQL syntax.
 *
 * @receiver *SelectStatement
 * @returns string
 */

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
	if 0 < len(s.ordering) {
		q += fmt.Sprintf("ORDER BY %s ", s.ordering.join(", "))
	}
	if 0 != s.limit {
		q += fmt.Sprintf("LIMIT %d ", s.limit)
	}
	if 0 != s.offset {
		q += fmt.Sprintf("OFFSET %d ", s.offset)
	}
	return strings.TrimRight(q, " ")
}

/**
 * `Find` will inititilze a Librarian `SelectStatement` with the assumption of
 * finding a record by an `id` column, returning a pointer to the
 * `SelectStatement` for continued chaning.
 *
 * Ex: SELECT "users".* FROM "users" WHERE "users"."id" = 3
 *	With(session).Search(users).Find(3).Query() // #Query or #First can be used here.
 *
 * @params int
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Find(i int) (result, error) {
	return s.Where(s.a("id").Eq(i)).First()
}

/**
 * `Count` removes a `SelectStatement` projections as they're no longer needed,
 * calling SQLs COUNT instead.
 *
 * @receiver *SelectStatement
 * @returns int
 */

func (s *SelectStatement) Count() (int, error) {
	s.count = true
	s.projections = []column{column(fmt.Sprintf("COUNT(%s)", s.a("*")))}
	res, err := s.process()
	if 1 == len(res) && nil == err {
		count, ok := res[0]["count"]
		if ok {
			return int(count.(int64)), err
		} else {
			return 0, BadResultsError
		}
	}
	return 0, BadResultsError
}

/**
 * `First` returns the first result stored in a database
 * that meet the `SelectStatements` criteria or an error.
 *
 * @receiver *SelectStatement
 * @returns result, error
 */

func (s *SelectStatement) First() (result, error) {
	s.limit = 1
	res, err := s.process()
	if 1 > len(res) {
		return nil, err
	}
	return res[0], err
}

/**
 * `Query` returns a `results` type (array of []result) of records
 * stored in the databse the meet the `SelectStatements` criteria or an
 * error.
 *
 * @receiver *SelectStatement
 * @returns results, error
 */

func (s *SelectStatement) Query() (results, error) {
	return s.process()
}

/**
 * `Clone` creates a new instance or a `SelectStatement` based
 * on the values currently associated with the receiver, returning
 * a pointer to the clone.
 *
 * @receiver *SelectStatement
 * @returns *SelectStatement
 */

func (s *SelectStatement) Clone() *SelectStatement {
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

/**
 * `process` is called interally to generate the results to return
 * when a statement finalizer is executed (ex. #First, #Query, #Count)
 *
 * @receiver *SelectStatement
 * @returns results, error
 */

func (s *SelectStatement) process() (results, error) {
	if nil == s.session || nil != s.session.Ping() {
		return nil, BadSessionError
	}
	sqlQuery := s.ToSQL()
	sqlStatment, err := s.session.Prepare(sqlQuery)
	if nil != err {
		return nil, err
	}
	if 0 == s.limit && !s.count {
		sqlPreflightResult, _ := sqlStatment.Exec()
		sqlPreflightResultCount, _ := sqlPreflightResult.RowsAffected()
		s.limit = int(sqlPreflightResultCount)
	}
	sqlRows, err := sqlStatment.Query()
	if nil != err {
		return nil, err
	}
	sqlColumns, err := sqlRows.Columns()
	if nil != err {
		return nil, err
	}
	sqlReturnSize := s.limit | 1
	sqlResultsArray := make(results, sqlReturnSize)
	sqlCurrentResultIndex := 0
	defer sqlStatment.Close()
	defer sqlRows.Close()
	defer logQueryInformation(time.Now(), sqlQuery)
	for sqlRows.Next() {
		sqlResultsBuffer := generateResultsBuffer(len(sqlColumns))
		err = sqlRows.Scan(sqlResultsBuffer...)
		sqlResultsArray[sqlCurrentResultIndex] = generateResultMap(sqlColumns, sqlResultsBuffer)
		sqlCurrentResultIndex += 1
	}
	return sqlResultsArray, err
}

/**
 * `generateResultMap` is used internally to generate a `result` type.
 * Taking a queries column names, mapping them to values stored in an
 * array used to buffer a query response.
 *
 * @params []string, []interface{}
 * @returns result
 */

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

/**
 * `generateResultsBuffer` is used internally to generate an array
 * of type `interface{}` to store values returned from Rows #Scan
 * method.
 *
 * @params int
 * @returns []interface{}
 */

func generateResultsBuffer(l int) []interface{} {
	p := make([]interface{}, l)
	for i := 0; i < len(p); i++ {
		var buf interface{}
		p[i] = &buf
	}
	return p
}

func logQueryInformation(t time.Time, q string) {
	fmt.Printf("librarian \033[96m::\033[0m (%v) \033[96m::\033[0m %s\n", time.Now().Sub(t), q)
}
