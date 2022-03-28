package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println(ExitMessage())
  os.Exit(0)
}

func ExitMessage() string {
  return "ok"
}
