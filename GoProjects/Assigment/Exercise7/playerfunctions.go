package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("HOME PAGE")
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	player := &Player{}
	err := json.NewDecoder(r.Body).Decode(player)
	debugLog.Println(err)
	debugLog.Println("Create ", *player)
	result := db.Create(player)
	if result.Error != nil || result.RowsAffected == 0 {
		errorLog.Printf("Error while adding player %v", *player)
		//fix error response
	}
	json.NewEncoder(w).Encode("Player with jersey no. - " + strconv.Itoa(player.JerseyNo) + " is added succesfully!")

}

func FindAllPlayers(w http.ResponseWriter, r *http.Request) {
	var players []Player
	db.Find(&players)
	//db.First(&players)
	json.NewEncoder(w).Encode(players)
}

func FindPlayerById(w http.ResponseWriter, r *http.Request) {
	jerseyNo := mux.Vars(r)["jerseyNo"]
	if jerseyNo == "" {
		json.NewEncoder(w).Encode("Enter some jersey No.")
		return
	}
	player := &Player{}
	db.First(player, jerseyNo)
	if player.JerseyNo == 0 {
		json.NewEncoder(w).Encode("Player with given id doesn't exist")
		return
	}
	debugLog.Println("FindPLayerByID ", *player)
	json.NewEncoder(w).Encode(*player)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	jerseyNo := mux.Vars(r)["jerseyNo"]
	if jerseyNo == "" {
		json.NewEncoder(w).Encode("Enter some jersey No.")
		return
	}

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	player := &Player{}
	updatedPlayer := &Player{}
	db.First(player, jerseyNo)
	debugLog.Println(jerseyNo)
	debugLog.Println("updated player ", *updatedPlayer)
	debugLog.Println("player ", player)
	if player.JerseyNo == 0 {
		json.NewEncoder(w).Encode("Player with given id doesn't exist")
		return
	}

	json.NewDecoder(r.Body).Decode(updatedPlayer)

	if updatedPlayer.Name == "" {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}
	debugLog.Println("update ", *player)
	player.Age = updatedPlayer.Age
	player.Name = updatedPlayer.Name
	player.NoOfCenturies = updatedPlayer.NoOfCenturies
	result := db.Save(player)
	if result.Error != nil || result.RowsAffected == 0 {
		errorLog.Printf("Error while adding player %v", *player)
		json.NewEncoder(w).Encode("ERROR!")
		return
	}
	json.NewEncoder(w).Encode("Player with jersey no. - " + strconv.Itoa(player.JerseyNo) + " is updated succesfully!")

}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	jerseyNo := mux.Vars(r)["jerseyNo"]
	if jerseyNo == "" {
		json.NewEncoder(w).Encode("Enter some jersey No.")
		return
	}
	player := &Player{}
	db.First(player, jerseyNo)
	if player.JerseyNo == 0 {
		json.NewEncoder(w).Encode("Player with given id doesn't exist")
		return
	}
	debugLog.Println(" delete ", *player)
	result := db.Delete(player)
	if result.Error != nil || result.RowsAffected == 0 {
		errorLog.Printf("Error while adding player %v", *player)
		json.NewEncoder(w).Encode("ERROR!")
		return
	}
	json.NewEncoder(w).Encode("Player with jersey no. - " + strconv.Itoa(player.JerseyNo) + " is deleted succesfully!")
}
