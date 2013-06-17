package utils

import (
  "fmt"
)

type Stringified []string

func (str *Stringified) Push(str string) *Stringified {
  *str = append(*str, items...)
  return str
}

func Stringified() *Stringified {
  return &Stringified{}
}
