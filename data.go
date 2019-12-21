package main

type botInfo struct {
	name      string
	token     string
	responses []response
}

type query int

const (
	greet query = iota
	quit  query = iota
)

type response struct {
	state     query
	responses []string
}

// Watcher : Notification bot
var Watcher = botInfo{
	name:  "Watcher",
	token: WatcherToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want?", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	}}

// Wyrm : Bot manager
var Wyrm = botInfo{name: "Wyrm", token: WyrmToken}

// Manager : User Manager
var Manager = botInfo{name: "Manager", token: ManagerToken}
