package matchmaking

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/0110101001110011/blade2/src/game"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/gorilla/websocket"
)

const MaxChannelSize = 10240
const MaxPollWait int64 = 1000

var matchmakingQueue = make(chan *Client, MaxChannelSize)
var running = true
var initTime = int64(0)

func poll() {
	for {
		if len(matchmakingQueue) > 1 {
			c1, c2 := <-matchmakingQueue, <-matchmakingQueue

			js, err := json.Marshal(templates.Instruction{Instruction: game.GameFound})
			if err != nil {
				log.Println("Failed to marshal JSON object when attempting to alert client [c1] that a game was found")
				c1.sendMessage(templates.GenericError)
			} else {
				c1.sendMessage(js)
			}

			js, err = json.Marshal(templates.Instruction{Instruction: game.GameFound})
			if err != nil {
				log.Println("Failed to marshal JSON object when attempting to alert client [c2] that a game was found")
				c2.sendMessage(templates.GenericError)
			} else {
				c2.sendMessage(js)
			}
		} else {
			time.Sleep(time.Millisecond * time.Duration(MaxPollWait))

			if !running {
				break
			}
		}
	}
}

func JoinQueue(c *websocket.Conn) {
	client := NewClient(c)
	fmt.Printf("Added client [%s] to the matchmaking queue\n", client.ID)
	client.run()

	matchmakingQueue <- &client

	js, err := json.Marshal(templates.StandardJSON{Status: game.OK, Message: client.ID})
	if err != nil {
		log.Println("Failed to marshal JSON object when attempting to return the client ID after joining the matchmaking queue")
		client.sendMessage(templates.GenericError)
	} else {
		log.Printf("Returning JSON object [%s]\n", js)
		client.sendMessage(js)
	}
}

func Init() {
	initTime = time.Now().Unix()
	go poll()
}
