package main

import "math/rand"

type caveBot struct {
	name      string
	token     string
	responses []response
}

type query int

func (bot caveBot) getResponse(queryType query) string {
	for _, response := range bot.responses {
		if response.state == queryType {
			return response.responses[rand.Intn(len(response.responses))]
		}
	}
	return "ok"
}

const (
	greet query = iota
	quit  query = iota
)

type response struct {
	state     query
	responses []string
}

type debugValues struct {
	notificationChannel string
}

// Watcher : Notification bot
var Watcher = caveBot{
	name:  "Watcher",
	token: WatcherToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want?", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	}}

// Wyrm : Bot manager
var Wyrm = caveBot{
	name:  "Wyrm",
	token: WyrmToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want?", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	}}

// Manager : User Manager
var Manager = caveBot{name: "Manager", token: ManagerToken}

// Debug : Values used for debugging bots
var Debug = debugValues{notificationChannel: "540531918769225738"}
