package utils

import (
  "fmt"
  "strings"
)

func Quote(value interface{}) string {
  return fmt.Sprintf(`"%v"`, value)
}

func Tag(value interface{}) string {
  return fmt.Sprintf(`'%v'`, value)
}

func Trim(value interface{}) string {
  return strings.Trim(strings.TrimRight(fmt.Sprintf("%v", value), "0"), " ")
}
