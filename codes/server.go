package main

import "net"
import "fmt"
import "bufio"
import "strings" // требуется только ниже для обработки примера

func main() {

  fmt.Println("Launching server...")

  // Устанавливаем прослушивание порта , _
  ln := net.Listen("tcp", ":8081")

  // Открываем порт , _
  conn := ln.Accept()

  // Запускаем цикл
  for {
    // Будем прослушивать все сообщения разделенные \n , _
    message := bufio.NewReader(conn).ReadString('\n')
    if message == "1" {
        fmt.Print("You pidor")
    }
    // Распечатываем полученое сообщение
    fmt.Print("Message Received:", string(message))
    // Процесс выборки для полученной строки
    newmessage := strings.ToUpper(message)
    // Отправить новую строку обратно клиенту
    conn.Write([]byte(newmessage + "\n"))
  }
}
