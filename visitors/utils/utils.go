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

func Map(fn func(interface{}) interface{}, items interface{}) []interface{} {
  array := []interface{}{items}
  result := make([]interface{}, len(array))
  for index, value := range array {
    result[index] = fn(value)
  }
  return result
}

func Stringer(array []interface{}) []string {
  result := make([]string, len(array))
  for index, value := range array {
    result[index] = fmt.Sprintf("%v", value)
  }
  return result
}

func Join(delimiter string, items interface{}) string {
  return strings.Join(Stringer(Map(func(item interface{}) interface{} {
    return fmt.Sprintf("%v", item)
  }, items)), delimiter)
}
