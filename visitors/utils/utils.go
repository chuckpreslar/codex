package utils

import (
  "fmt"
  "regexp"
  "strings"
)

func Quote(value interface{}) string {
  return fmt.Sprintf(`"%v"`, value)
}

func Tag(value interface{}) string {
  if IsNumber(value) || IsSqlFunction(value) {
    return strings.TrimRight(fmt.Sprintf(`%v`, value), "0")
  }
  return fmt.Sprintf(`'%v'`, value)
}

func IsSqlFunction(value interface{}) bool {
  matcher, _ := regexp.Compile(`^\w+\((\w+)?\)`)
  return matcher.MatchString(value.(string))
}

func IsNumber(value interface{}) bool {
  matcher, _ := regexp.Compile(`\d`)
  return matcher.MatchString(value.(string))
}
