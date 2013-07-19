package visitors

import (
  "errors"
  "fmt"
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

// Begin Helpers.

func (self *MySqlVisitor) QuoteTableName(o interface{}) string {
  return fmt.Sprintf("`%v`", o)
}

func (self *MySqlVisitor) QuoteColumnName(o interface{}) string {
  return fmt.Sprintf("`%v`", o)
}

// End Helpers.
