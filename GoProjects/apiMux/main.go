package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	JerseyNo      int    `gorm:"primaryKey" json:"jersey_no"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	NoOfCenturies int    `json:"no_of_centuries"`
}

func init() {
	CreateLoggers()
	OpenDBConnection()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/getAllPlayers", FindAllPlayers).Methods("GET")
	r.HandleFunc("/addPlayer", AddPlayer).Methods("POST")
	r.HandleFunc("/findPlayerByJerseyNo/{jerseyNo}", FindPlayerById).Methods("GET")
	r.HandleFunc("/updatePlayer/{jerseyNo}", UpdatePlayer).Methods("PUT")
	r.HandleFunc("/deletePlayer/{jerseyNo}", DeletePlayer).Methods("DELETE")

	infoLog.Fatal(http.ListenAndServe(":8080", r))
}
