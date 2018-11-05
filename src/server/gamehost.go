package server

import (
	"log"
	"math"
	"runtime"
	"time"

	"github.com/0110101001110011/blade2/src/e"
	"github.com/0110101001110011/blade2/src/templates"
	"github.com/0110101001110011/blade2/src/tools"
)

const idleTick = 2000 * time.Millisecond
const gameTick = (1000 / 50) * time.Millisecond

var queueCount = tools.MaxInt(int(float32(runtime.NumCPU())*0.75), 1)
var gamePipelines = make([][]*Game, queueCount)

// InitGameHost initializes the game management queue
func InitGameHost() {
	// Initialize
	// for i := range gamePipelines {
	// 	gamePipelines[i] = []*Game{}
	// }

	// Indices:
	// - pli (Pipeline Index)
	// - gi (Game Index)
	// - ci (Client Index)
	for pli := 0; pli < queueCount; pli++ {
		go func(pli int) {
			for {
				iterations := len(gamePipelines[pli])
				if iterations > 0 {
					for gi := 0; gi < iterations; gi++ {
						if gamePipelines[pli][gi].Client[0].IsAlive() && gamePipelines[pli][gi].Client[1].IsAlive() {
							gamePipelines[pli][gi].RelayUpdates()
						} else {
							for ci := 1; ci > 0; ci-- {
								if !gamePipelines[pli][gi].Client[ci].IsAlive() {
									log.Printf("Client [%s] in game [%d] dropped connection", gamePipelines[pli][gi].Client[ci].ID, gamePipelines[pli][gi].ID)
									gamePipelines[pli][gi].Client[ci].Drop(templates.Information{Code: e.Drop, Message: ""})
									gamePipelines[pli][gi].Client[1-ci].Drop(templates.Information{Code: e.OponentDroppedConnection, Message: ""})
								}
							}

							gamePipelines[pli] = append(gamePipelines[pli][:gi], gamePipelines[pli][gi+1:]...)
						}
					}

					time.Sleep(gameTick * 10)
				} else {
					time.Sleep(idleTick)
				}
			}
		}(pli)
	}
}

// AddGame will add a pointer to a game object to the least populated game pipeline
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

	gamePipelines[nextPipeIndex] = append(gamePipelines[nextPipeIndex], game)
}
