package server

import (
	"log"
	"math"
	"runtime"
	"time"

	"github.com/6a/blade2/src/enums"
	"github.com/6a/blade2/src/templates"
	"github.com/6a/blade2/src/tools"
)

const idleTick = 2000 * time.Millisecond
const gameTick = (1000 / 50) * time.Millisecond

var queueCount = tools.MaxInt(int(float32(runtime.NumCPU())*0.75), 1)
var gameLobbies = make([]chan *Game, queueCount)
var gamePipelines = make([][]*Game, queueCount)

// InitGameHost initializes the game management queue
func InitGameHost() {
	for i := range gameLobbies {
		gameLobbies[i] = make(chan *Game, lobbySize)
	}

	// Indices:
	// - pli (Pipeline Index)
	// - gi (Game Index)
	// - ci (Client Index)
	for pli := 0; pli < queueCount; pli++ {
		// Each pipeline gets it own goroutine
		go func(pli int) {
			for {
				// If some more games are queued to be added to the pipeline, add them
				// If there are no games to add, the loop breaks immediately
				processLobby := true
				for processLobby {
					select {
					case game := <-gameLobbies[pli]:
						gamePipelines[pli] = append(gamePipelines[pli], game)
					default:
						processLobby = false
					}
				}

				// - Count how many games are in this pipeline. This is done incase another game is added during iteration. If
				//   more games are added during this time, the following code is unaffected as the extra games are never accessed.
				// - Check every game to ensure both clients are alive.
				//   - Both alive: Prompt game to relay any stored updates between each client
				//	 - 1 or both dead: Add the games index (position in the pipeline) to hte list of games to remove
				iterations := len(gamePipelines[pli])

				// A list of games that have dead clients (and therefore need to be removed once iteration has finished)
				toRemove := []int{}
				if iterations > 0 {
					for gi := 0; gi < iterations; gi++ {
						if gamePipelines[pli][gi].Client[0].IsAlive() && gamePipelines[pli][gi].Client[1].IsAlive() {
							gamePipelines[pli][gi].RelayUpdates()
						} else {
							toRemove = append(toRemove, gi)
						}
					}

					// If the toRemove array contains indices, it means that there are some games to remove from the pipeline
					// This is done backwards to avoid indexing errors
					for tri := len(toRemove) - 1; tri >= 0; tri-- {
						// Calculate the index of the next game to remove
						egi := toRemove[tri]

						// For each client, check if it is dead. If it is dead, send a drop command (just in case) to the dead client
						// Also send a drop command to the other client. If its dead, it doesnt matter, but if it is alive it receives an
						// appropriate message
						for ci := 1; ci >= 0; ci-- {
							if !gamePipelines[pli][egi].Client[ci].IsAlive() {
								log.Printf("Client [%s] in game [%d] dropped connection", gamePipelines[pli][egi].Client[ci].ID, gamePipelines[pli][egi].ID)

								gamePipelines[pli][egi].Client[ci].Drop(templates.Information{Code: e.Drop, Message: ""})
								gamePipelines[pli][egi].Client[1-ci].Drop(templates.Information{Code: e.OponentDroppedConnection, Message: ""})
							}
						}

						gamePipelines[pli] = append(gamePipelines[pli][:egi], gamePipelines[pli][egi+1:]...)
					}

					time.Sleep(gameTick)
				} else {
					time.Sleep(idleTick)
				}
			}
		}(pli)
	}
}

// AddGame adds a Game (pointer) to the gamehost queue, so that it can start forwarding updates between hosts
func AddGame(game *Game) {
	lowestCount := math.MaxInt64
	nextPipeIndex := 0
	for ppi, pipeline := range gamePipelines {
		l := len(pipeline) // This may be a bit of a bottleneck. If that is the case, then it would be efficient to keep track of pipe utilization
		if l < lowestCount {
			lowestCount = l
			nextPipeIndex = ppi
		}
	}

	gameLobbies[nextPipeIndex] <- game
}
