package main

import "testing"

func TestExitMessage(t *testing.T) {
  result := ExitMessage()

  if result != "ok" {
    t.Errorf("function should return 'ok' but got %s", result)
  }
}
