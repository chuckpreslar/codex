package utils

import (
  "fmt"
  "regexp"
)

func Quote(value interface{}) string {
  return fmt.Sprintf(`"%v"`, value)
}

func Tag(value interface{}) string {
  if IsNumber(value) || IsSqlFunction(value) {
    return fmt.Sprintf(`%v`, value)
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
