package matchmaking

var matches = make(chan *Match)

func InitMatchManager() {
	go func() {
		match := <-matches
		match.ExecuteNext()
	}()
}
