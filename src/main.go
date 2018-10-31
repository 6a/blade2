// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/0110101001110011/blade2/src/matchmaking"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	matchmaking.JoinQueue(c)
}

func main() {
	var addr = flag.String("addr", "localhost:"+os.Args[1], "http service address")
	matchmaking.InitMatchMakingQueue()
	matchmaking.InitGameHost()

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/v1/session", connect)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
