package utils

import (
  "fmt"
  "strings"
)

func Join(slice []interface{}, delimiter string) string {
  joined := make([]string, len(slice))
  for index, value := range slice {
    joined[index] = fmt.Sprintf("%v", value)
  }
  return strings.Join(joined, delimiter)
}
