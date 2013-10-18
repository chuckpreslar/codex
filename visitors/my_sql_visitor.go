// Package visitors provides AST visitors for the codex package.
package visitors

import (
	"errors"
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/nodes"
)

type MySqlVisitor struct {
	*ToSqlVisitor
}

func (self *MySqlVisitor) Accept(o interface{}) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()

	result = self.Visit(o, self)

	return
}

// Begin Unary node visitors.

func (self *MySqlVisitor) VisitIndexName(o *nodes.IndexNameNode, visitor VisitorInterface) string {
	return fmt.Sprintf("`%v`", o.Expr)
}

// End Unary node visitors.

// Being Binary node visitors.

func (self *MySqlVisitor) VisitExistingColumn(o *nodes.ExistingColumnNode, visitor VisitorInterface) string {
	return fmt.Sprintf("MODIFY COLUMN %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
}

// End Binary node visitors.

// Begin Nary node visitors.

func (self *MySqlVisitor) VisitNotNull(o *nodes.NotNullNode, visitor VisitorInterface) string {
	var typ interface{}

	if 0 >= len(o.Options) {
		panic("Missing column type definition for MySql NOT NULL constraint.")
	} else {
		typ = o.Options[0]
	}

	return fmt.Sprintf("MODIFY %v %v NOT NULL", visitor.Visit(o.Columns, visitor), visitor.Visit(typ, visitor))
}

// End Nary node visitors.

// Begin Helpers.

func (self *MySqlVisitor) QuoteTableName(o interface{}) string {
	return fmt.Sprintf("`%v`", o)
}

func (self *MySqlVisitor) QuoteColumnName(o interface{}) string {
	return fmt.Sprintf("`%v`", o)
}

// End Helpers.
