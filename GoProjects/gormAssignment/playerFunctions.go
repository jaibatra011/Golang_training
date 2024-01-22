package main

import (
	"sync"
	"time"
)

func AddPlayer(players []*Player, p chan *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	//defer close(p)

	for i, player := range players {
		debugLog.Println(i, " add ", *player)
		result := db.Create(player)
		//debugLog.Println(result.Error)
		//debugLog.Println(result.RowsAffected)
		if result.Error != nil || result.RowsAffected == 0 {
			errorLog.Printf("Error while adding player %v", *player)
			continue
		}
		time.Sleep(1 * time.Second)
		p <- player
	}
	time.Sleep(time.Second)
}

func UpdatePlayer(players []*Player, p chan *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	//defer close(p)

	for i, updatedPlayer := range players {
		debugLog.Println(i, " update ", *updatedPlayer)
		time.Sleep(2 * time.Second)
		var player *Player
		db.First(&player, updatedPlayer.JerseyNo)
		player.Age = updatedPlayer.Age
		player.Name = updatedPlayer.Name
		player.NoOfCenturies = updatedPlayer.NoOfCenturies
		result := db.Save(player)
		if result.Error != nil || result.RowsAffected == 0 {
			errorLog.Printf("Error while adding player %v", *updatedPlayer)
			continue
		}
		p <- player
	}
	time.Sleep(time.Second)
}

func DeletePlayer(players []*Player, p chan *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	//defer close(p)

	for i, player := range players {
		debugLog.Println(i, " delete ", *player)
		time.Sleep(5 * time.Second)
		result := db.Delete(player)
		if result.Error != nil || result.RowsAffected == 0 {
			errorLog.Printf("Error while adding player %v", *player)
			continue
		}
		p <- player
	}
	time.Sleep(time.Second)
}

func PrintPlayer(addPlayerChannel chan *Player, updatePlayerChannel chan *Player, deletePlayerChannel chan *Player) {
	for {
		select {
		case p1 := <-addPlayerChannel:
			infoLog.Printf("%+v Player is added! \n", *p1)
		case p2 := <-updatePlayerChannel:
			infoLog.Printf("%+v Player is updated! \n", *p2)
		case p3 := <-deletePlayerChannel:
			infoLog.Printf("%+v Player is deleted! \n", *p3)
		}
	}
}
