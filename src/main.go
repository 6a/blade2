// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/0110101001110011/blade2/src/lib"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

var upgrader = websocket.Upgrader{} // use default options

var sessions = make(map[string]lib.GameState)

type SessionResponse struct {
	SessionID string
}

func session(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	u4 := uuid.Must(uuid.NewV4()).String()
	sessions[u4] = lib.NewGameState()
	response := SessionResponse{u4}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.WriteJSON(js)
	if err != nil {
		log.Println("write:", err)
		close(u4)
	}

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			close(u4)
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			close(u4)
		}
	}
}

func close(sessionID string) {
	delete(sessions, sessionID)
}

func main() {
	var addr = flag.String("addr", "localhost:"+os.Args[1], "http service address")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/v1/session", session)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
