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
	state    query
	response string
}

// Watcher : Notification bot
var Watcher = botInfo{
	name:  "Watcher",
	token: WatcherToken,
	responses: []response{
		response{state: greet, response: "What do you want?"},
		response{state: greet, response: "Yes?"},
		response{state: greet, response: "Anything you need?"},
		response{state: quit, response: "Cya!"},
		response{state: quit, response: "Somebody take over for me?"},
		response{state: quit, response: "Gone for tea"},
	}}

// Wyrm : Bot manager
var Wyrm = botInfo{name: "Wyrm", token: WyrmToken}

// Manager : User Manager
var Manager = botInfo{name: "Manager", token: ManagerToken}
