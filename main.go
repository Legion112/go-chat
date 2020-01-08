package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

type Message struct {
	Id	int
	Date	int
	Text     string
	Username string
}

// This example demonstrates a trivial echo server.
func main() {
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1)/gochat")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, username, text, date FROM message ORDER BY id DESC LIMIT 5")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var messages = make(map[int]Message);
	for rows.Next() {
		var message = Message{};

		if err := rows.Scan(&message.Id, &message.Username, &message.Text, &message.Date);
		err != nil {
			panic(err)
		}
		messages[message.Id] = message
	}

	//if !rows.NextResultSet() {
	//	panic( rows.Err())
	//}
	var jsonString []byte;
	jsonString, err = json.Marshal(messages)
	if err != nil {
		panic(err)
	}
	print(string(jsonString));
	return;

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

		effected, err := stmtIns.Exec(time.Now().Unix(), message.Username, message.Text)
		fmt.Println(effected)
		if err != nil {
			panic(err)
		}

		websocket.Message.Send(ws, "sddsf");
		fmt.Println("Message from client " + message.Text + " " + message.Username)
	}))


	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err.Error())
		panic("ListenAndServe: " + err.Error())
	}
}