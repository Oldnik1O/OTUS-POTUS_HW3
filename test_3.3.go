// TestLogCommand
// Проверяет, что команда для записи в лог выполняется корректно и не добавляет других команд в очередь

func TestLogCommand(t *testing.T) {
  q := &Queue{}

  // Добавляем команду LogCommand в очередь
  q.AddCommand(&LogCommand{message: "Test message"})
  q.ProcessCommands()

  if len(q.commands) != 0 {
    t.Fatalf("Expected no commands in the queue, got %d", len(q.commands))
  }
}