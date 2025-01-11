package main

import (
	// "net"
	"net/http"
	// "strings"
	// "encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}
	defer c.Close()

	//First Wait for the password:
	_, message, err2 := c.ReadMessage() //ReadMessage blocks until message received
	if err2 != nil {
		fmt.Println("readPassErr:", err2)
	}

	if string(message) != "password" {
		return  //if password wrong don't let gamer connect
	}

	// receive drawing commands
	for {
		_, message, err := c.ReadMessage() //ReadMessage blocks until message received
		if err != nil {
			fmt.Println("readDrawCommandErr:", err)
		}

		fmt.Println(message)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/echo", echo) //this request comes from webrtc.html
	http.Handle("/", fileServer)

	err := http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
		panic(err)
	}
}