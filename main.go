package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

type Message struct {
	Text string
	UserName string
}

// This example demonstrates a trivial echo server.
func main() {
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1)/gochat")
	if err != nil {
		panic(err)
	}

	http.Handle("/echo", websocket.Handler(func(ws *websocket.Conn) {
		var message Message

		err = websocket.JSON.Receive(ws, &message) // ? = placeholder
		if err != nil {
			panic(err)
		}
		stmtIns, err := db.Prepare("INSERT INTO message (`date`, `username`, `text`)  VALUES( ?, ?, ?)")
		if err != nil {
			panic(err)
		}

		effected, err := stmtIns.Exec(time.Now().Unix(), message.UserName, message.Text)
		fmt.Println(effected)
		if err != nil {
			panic(err)
		}

		websocket.Message.Send(ws, "Success")
		fmt.Println("Message from client " + message.Text + " " + message.UserName)
	}))


	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err.Error())
		panic("ListenAndServe: " + err.Error())
	}
}