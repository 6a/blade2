// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/0110101001110011/blade2/src/server"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	server.JoinQueue(c)
}

func main() {
	// Define the websocket that will be brought up
	var addr = flag.String("addr", "localhost:"+os.Args[1], "http service address")
	flag.Parse()

	// Set up logging
	log.SetFlags(0)

	// Define the endpoint URL and the function that will handle it
	http.HandleFunc("/v1/session", connect)

	// Seed the random generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Spin up the various async server pipelines
	server.InitMatchMakingQueue()
	server.InitGameHost()

	// Start serving the websocket - Log and exit on error
	log.Fatal(http.ListenAndServe(*addr, nil))
}
