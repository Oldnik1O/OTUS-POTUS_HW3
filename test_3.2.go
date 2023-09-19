// TestLogAfterMultipleFailures
// Проверяет, что после двух неудачных попыток выполнения RetryCommand, команда для записи в лог будет добавлена в очередь

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



func TestLogAfterMultipleFailures(t *testing.T) {
  q := &Queue{}

  // Добавляем команду RetryCommand, которая уже имеет 2 попытки
  q.AddCommand(&RetryCommand{
    originalCommand: &AlwaysFailCommand{},
    attempts:        2,
  })
  q.ProcessCommands()

  if len(q.commands) != 1 {
    t.Fatalf("Expected 1 command in the queue, got %d", len(q.commands))
  }

  _, ok := q.commands[0].(*LogCommand)
  if !ok {
    t.Fatal("Expected a LogCommand in the queue")
  }
}
