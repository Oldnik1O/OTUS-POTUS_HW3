\\Механизм обработки исключений в игре "Космическая битва" на языке Go. 
\\ т.к. Go использует отличный от многих других языков механизм обработки ошибок 
\\ вместо исключений и блоков try-catch, Go использует возвращаемые значения ошибок
\\ нужно явно проверять, была ли ошибка возвращена, и затем решать, что делать с этой ошибкой. 

\\ Command - интерфейс, описывающий команды, которые мы можем выполнить.
\\ Queue - структура, представляющая очередь команд. Мы добавляем команды в эту очередь и обрабатываем их.
\\  handleError - функция, которая решает, что делать при возникновении ошибки при выполнении команды.
\\ RetryCommand - команда, которая пытается повторно выполнить оригинальную команду. Если она не удается после нескольких попыток, вызывается LogCommand для записи ошибки.
\\ LogCommand - команда, которая записывает ошибку в журнал.

\\ Создание базовых структур и интерфейсов:

package main

import (
  "fmt"
  "log"
)

type Command interface {
  Execute() error
}

type Queue struct {
  commands []Command
}

func (q *Queue) AddCommand(c Command) {
  q.commands = append(q.commands, c)
}

func (q *Queue) ProcessCommands() {
  for len(q.commands) > 0 {
    command := q.commands[0]
    q.commands = q.commands[1:]
    if err := command.Execute(); err != nil {
      handleError(command, err)
    }
  }
}

func handleError(command Command, err error) {
  switch cmd := command.(type) {
  case *RetryCommand:
    q.AddCommand(&LogCommand{message: err.Error()})
  case *LogCommand:
    log.Println("Failed to log error:", err.Error())
  default:
    q.AddCommand(&RetryCommand{originalCommand: command})
  }
}

type RetryCommand struct {
  originalCommand Command
  attempts        int
}

func (rc *RetryCommand) Execute() error {
  rc.attempts++
  if rc.attempts > 2 {
    return fmt.Errorf("failed after %d attempts", rc.attempts)
  }
  return rc.originalCommand.Execute()
}

type LogCommand struct {
  message string
}

func (lc *LogCommand) Execute() error {
  log.Println(lc.message)
  return nil
}

var q = &Queue{}

func main() {
  // Testing the mechanism
  q.AddCommand(&RetryCommand{})
  q.ProcessCommands()
}

