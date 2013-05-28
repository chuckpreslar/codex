/*
 *  The MIT License (MIT)
 *
 *  Copyright (c) 2013 Chuck Preslar
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 */

package librarian

import (
	"database/sql"
)

type StatementInitializer struct {
	session *sql.DB
}

func With(session *sql.DB) StatementInitializer {
	return StatementInitializer{session}
}

func (s StatementInitializer) Search(a accessor) *SelectStatement {
	return &SelectStatement{
		a:           a,
		projections: []column{},
		reference:   table(a("")),
		ordering:    []string{},
		session:     s.session,
	}
}

func (s StatementInitializer) Insert(a accessor) *InsertStatement {
	return &InsertStatement{a: a}
}

func (s StatementInitializer) Update(a accessor) *UpdateStatement {
	return &UpdateStatement{a: a}
}

func (s StatementInitializer) Delete(a accessor) *DeleteStatement {
	return &DeleteStatement{a: a}
}
