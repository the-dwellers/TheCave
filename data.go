package main

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// CaveBot : Simple Bot Structure
type CaveBot struct {
	name      string
	token     string
	responses []response
	intent    discordgo.Intent
}

type query int

func (bot CaveBot) getResponse(queryType query) string {
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
var Watcher = CaveBot{
	name:  "Watcher",
	token: WatcherToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want?", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	},
	intent: discordgo.IntentsGuildMessages,
}

// Wyrm : Bot manager
var Wyrm = CaveBot{
	name:  "Wyrm",
	token: WyrmToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want?", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	},
	intent: discordgo.IntentsGuildMessages,
}

// Manager : User Manager
var Manager = CaveBot{name: "Manager", token: ManagerToken}

// CarolineBot : Broadcasting music bot
var CarolineBot = CaveBot{
	name:  "Caroline",
	token: CarolineToken,
	responses: []response{
		response{state: greet, responses: []string{"What do you want", "Yes?", "Anything you need?"}},
		response{state: quit, responses: []string{"Cya!", "Somebody take over for me?", "Gone for tea"}},
	},
	intent: discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates,
}

// Debug : Values used for debugging bots
var Debug = debugValues{notificationChannel: "540531918769225738"}
