// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/0110101001110011/blade2/src/game"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var data map[string]interface{}

type StandardJSON struct {
	Status  game.EStatus
	Message string
}

type Instruction struct {
	Instruction game.EInstruction
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/v1/session"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			//log.Printf("Error: [%s]", err)
		} else {
			log.Printf("Received a message: [%d] %s", mt, message)
		}

		// time.Sleep(1 * time.Second)
	}
}
