package main

import (
	"sync"
)

type Player struct {
	JerseyNo      int `gorm:"primaryKey"`
	Name          string
	Age           int
	NoOfCenturies int
}

func main() {
	CreateLoggers()
	OpenDBConnection()

	wg := new(sync.WaitGroup)
	wg.Add(3)

	//Add player
	addPlayers := GetAddPlayerList()
	addPlayerChannel := make(chan *Player)
	go AddPlayer(addPlayers, addPlayerChannel, wg)

	//Update player
	updatePlayers := GetUpdatedPlayerList()
	updatePlayerChannel := make(chan *Player)
	go UpdatePlayer(updatePlayers, updatePlayerChannel, wg)

	//Delete player
	deletePlayerChannel := make(chan *Player)
	go DeletePlayer(addPlayers[4:6], deletePlayerChannel, wg)

	//Print whatever player is added, updated or deleted
	go PrintPlayer(addPlayerChannel, updatePlayerChannel, deletePlayerChannel)

	wg.Wait()
}
