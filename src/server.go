package main

import (
	"log"
	"net/http"
	//"github.com/stianeikeland/go-rpio"
)

type Team uint

const (
	Atleti Team = iota
	Madrid
)

var (
	goals = [2]uint{0, 0}

//	pin = rpio.Pin(14)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func resetGoals() {
	for i := range goals {
		goals[i] = 0
	}
}

func incrementGoals(team Team) {
	goals[team]++
}

func decrementGoals(team Team) {
	if goals[team] > 0 {
		goals[team]--
	}
}

func setGoals(team Team, howMany uint) {
	goals[team] = howMany
}

func updateGoals(team Team) {
}

func main() {
	resetGoals()
	log.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", http.FileServer(http.Dir("public/")))
}
