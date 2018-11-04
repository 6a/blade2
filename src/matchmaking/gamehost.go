package matchmaking

import (
	"log"
	"runtime"
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/0110101001110011/blade2/src/tools"
)

const idleTick = 2000 * time.Millisecond
const gameTick = (1000 / 50) * time.Millisecond
const queueSize = 2048

var queueCount = tools.MaxInt(int(float32(runtime.NumCPU())*0.75), 1)
var gamePipelines = make([]chan *Game, queueCount)

// InitGameHost initializes the game management queue
func InitGameHost() {
	for i := range gamePipelines {
		gamePipelines[i] = make(chan *Game, queueSize)
	}

	for index := 0; index < queueCount; index++ {
		go func(index int) {
			executed := false
			currentGameIndex := 0
			for {
				select {
				case game := <-gamePipelines[index]:
					if game.Client[0].IsAlive() && game.Client[1].IsAlive() {
						executed = true
						game.RelayUpdates()
						gamePipelines[index] <- game
					} else {
						for _, client := range game.Client {
							if !client.IsAlive() {
								log.Printf("Client [%s] in game [%d] dropped connection", client.ID, game.ID)
							}

							client.Drop(templates.Information{Code: e.OponentDroppedConnection, Message: ""})
						}
					}

					currentGameIndex++
					if len(gamePipelines[index]) >= currentGameIndex {
						currentGameIndex = 0
						time.Sleep(gameTick)
					}
				default:
					if !executed {
						time.Sleep(idleTick)
					}

					executed = false
				}
			}
		}(index)
	}
}

// AddGame will add a pointer to a game object to the least populated game pipeline
func AddGame(game *Game) {
	lowestCount := queueSize
	nextPipe := gamePipelines[0]
	for _, pipeline := range gamePipelines {
		l := len(pipeline) // This may be a bit of a bottleneck. If that is the case, then it would be efficient to keep track of pipe utilization
		if l < lowestCount {
			lowestCount = l
			nextPipe = pipeline
		}
	}

	nextPipe <- game
}
