package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Player struct {
	JerseyNo      int    `json:"jersey_no"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	NoOfCenturies int    `json:"no_of_centuries"`
}

func main() {
	fmt.Println("Welcome to API Client \n")
	homePage := "http://localhost:8080/"

	GetRequest(homePage) //Getting homepage

	player1 := `
	{
		"jersey_no":7, 
		"name":"MSD", 
		"age":43, 
		"no_of_centuries":34
	}
	`
	PostRequest(homePage+"addPlayer", player1) //adding player 1

	player2 := Player{18, "Virat Kohli", 34, 79}
	b, err := json.Marshal(player2)
	if err != nil {
		panic(err)
	}
	PostRequest(homePage+"addPlayer", string(b)) //adding player 2

	GetRequest(homePage + "getAllPlayers")

	GetRequest(homePage + "findPlayerByJerseyNo/18")

	updatedPlayer := Player{18, "Virat Kohli", 34, 80}
	res, err := json.Marshal(updatedPlayer)
	if err != nil {
		panic(err)
	}

	PutRequest(homePage+"updatePlayer/18", string(res))

	GetRequest(homePage + "findPlayerByJerseyNo/18")

	DeleteRequest(homePage + "deletePlayer/18")

	GetRequest(homePage + "getAllPlayers")

}

func GetRequest(url string) {
	fmt.Println("Hitting " + url + " GET method")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content) + "\n")
}

func PostRequest(url string, body string) {
	fmt.Println("Hitting " + url + " POST method")
	requestbody := strings.NewReader(body)
	resp, err := http.Post(url, "application/json", requestbody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content) + "\n")
}

func PutRequest(url string, body string) {
	fmt.Println("Hitting " + url + " PUT method")
	requestbody := strings.NewReader(body)
	req, err := http.NewRequest("PUT", url, requestbody) // creating request
	//resp, err := http.Put(url, "application/json", requestbody)
	if err != nil {
		panic(err)
	}
	client := &http.Client{} //creating client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content) + "\n")
}

func DeleteRequest(url string) {
	fmt.Println("Hitting " + url + " DELETE method")
	req, err := http.NewRequest("DELETE", url, nil) // creating request
	if err != nil {
		panic(err)
	}
	client := &http.Client{} //creating client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content) + "\n")
}
