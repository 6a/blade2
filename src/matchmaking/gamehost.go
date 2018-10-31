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
const queueSize = 2048

var queueCount = tools.Max(int(float32(runtime.NumCPU())*0.75), 1)
var gamePipelines = make([]chan *Game, queueCount)

// InitGameHost initializes the game management queue
func InitGameHost() {
	for i := range gamePipelines {
		gamePipelines[i] = make(chan *Game, queueSize)
	}

	for index := 0; index < queueCount; index++ {
		go func(index int) {
			executed := false
			for {
				select {
				case game := <-gamePipelines[index]:
					if game.Client[0].IsAlive() && game.Client[1].IsAlive() {
						executed = true
						game.ExecuteNext()
						gamePipelines[index] <- game
					} else {
						for _, client := range game.Client {
							if !client.IsAlive() {
								log.Printf("Client [%s] in game [%d] dropped connection", client.ID, game.ID)
							}

							client.Drop(templates.Information{Status: e.OponentDroppedConnection, Message: ""})
						}
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
		l := len(pipeline)
		if l < lowestCount {
			lowestCount = l
			nextPipe = pipeline
		}
	}

	nextPipe <- game
}
