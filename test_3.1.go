// TestRetryCommand:
// Проверяет, что при ошибке команда будет добавлена в очередь с помощью RetryCommand

package main

import (
  "errors"
  "testing"
)

// Mock команда, которая всегда завершается ошибкой
type AlwaysFailCommand struct{}

func (a *AlwaysFailCommand) Execute() error {
  return errors.New("always fail")
}

func TestRetryCommand(t *testing.T) {
  q := &Queue{}

  // Добавляем команду, которая всегда вызывает ошибку, в очередь
  q.AddCommand(&AlwaysFailCommand{})
  q.ProcessCommands()

  if len(q.commands) != 1 {
    t.Fatalf("Expected 1 command in the queue, got %d", len(q.commands))
  }

  switch cmd := q.commands[0].(type) {
  case *RetryCommand:
    if cmd.attempts != 1 {
      t.Fatalf("Expected 1 attempt, got %d", cmd.attempts)
    }
  default:
    t.Fatal("Expected a RetryCommand in the queue")
  }
}

