package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/6a/blade2/src/enums"
	"github.com/6a/blade2/src/templates"
	"github.com/gorilla/websocket"
)

const queueTickRate = (1000 / 1) * time.Millisecond
const lobbySize = 10240

// Use a channel to prevent race conditions - new clients are added to the lobby channel, which is then
// checked and dumped into the matchmakingQueue slice at the end of the matchmaking queue loop
// This is done instead of a single channel to reduce complexity and memory thrashing that may come about from
// frequently removing and replacing things in a channel
var lobby = make(chan *Client, lobbySize)
var matchmakingQueue []*Client

var running = true

func processMatchmakingQueue() {
	for {
		// Exit early if the server is being shut down or something
		if !running {
			break
		}

		// If some more clients connected, add them to the queue
		// If there are no client to add, the loop breaks immediately
		processLobby := true
		for processLobby {
			select {
			case client := <-lobby:
				matchmakingQueue = append(matchmakingQueue, client)
			default:
				processLobby = false
			}
		}

		// If there are at least 2 clients in the matchmaking queue
		if len(matchmakingQueue) > 1 {
			// Find two live clients
			var clients []*Client
			queuePos := 0
			for {
				// If the current client being checked is alive, add it to the list of clients
				// Otherwise drop it and keep searching
				if matchmakingQueue[queuePos].IsAlive() {
					clients = append(clients, matchmakingQueue[queuePos])
				} else {
					matchmakingQueue[queuePos].Drop(templates.MakeJSON(templates.Information{Code: e.Drop, Message: ""}))
				}

				queuePos++

				// If two live clients are found, generate a match and add it to the game host
				if len(clients) > 1 {
					matchmakingQueue = matchmakingQueue[queuePos:]
					for index := 0; index < 2; index++ {
						clients[index].sendMessage(templates.MakeJSON(templates.Information{Code: e.MatchFound, Message: strconv.Itoa(index)}))
					}

					game := CreateGame(clients[0], clients[1])
					AddGame(&game)
					break
				}

				// If the number of clients that are needed exceeds the number of unchecked clients in the queue, exit early
				if len(matchmakingQueue[queuePos:]) < (2 - len(clients)) {
					matchmakingQueue = matchmakingQueue[queuePos:]
					break
				}
			}
		} else {
			time.Sleep(queueTickRate)
		}
	}
}

// JoinQueue creates a Client object and adds it to the matchmaking queue
func JoinQueue(c *websocket.Conn) {
	client := NewClient(c)
	client.activate()
	client.sendMessage(templates.MakeJSON(templates.Information{Code: e.Connected, Message: client.ID}))
	fmt.Printf("Added client [%s] to the matchmaking queue\n", client.ID)
	lobby <- &client
}

// InitMatchMakingQueue initializes the matchmaking queue
func InitMatchMakingQueue() {
	go processMatchmakingQueue()
}
