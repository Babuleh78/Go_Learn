package main

import (
    "html/template"
    "log"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // подключенные клиенты
var broadcasts = make(chan Message)           // канал для отправки сообщений
var mu sync.Mutex                             // мьютекс для защиты доступа к clients

// Сообщение
type Message struct {
    Text string `json:"text"`
}

// WebSocket соединения
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// Обработка WebSocket соединений
func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Ошибка при подключении:", err)
        return
    }
    defer conn.Close()

    mu.Lock()
    clients[conn] = true
    mu.Unlock()

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Println("Ошибка при чтении JSON:", err)
            mu.Lock()
            delete(clients, conn)
            mu.Unlock()
            break
        }
        broadcasts <- msg
    }
}

// Отправка сообщений всем подключенным клиентам
func handleMessages() {
    for {
        msg := <-broadcasts
        mu.Lock()
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Println("Ошибка при отправке JSON:", err)
                client.Close()
                delete(clients, client)
            }
        }
        mu.Unlock()
    }
}

// Главная страница
func serveHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func main() {
    go handleMessages()

    http.HandleFunc("/", serveHome)
    http.HandleFunc("/ws", handleConnections)
    log.Println("Сервер запущен на :8080")
    http.ListenAndServe(":8080", nil)
}