package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	bot := Wyrm
	discord, err := discordgo.New("Bot " + bot.token)
	if err != nil {
		fmt.Printf("! Error starting discord session: %s", err)
	}

	// Handlers
	discord.AddHandler(ready)
	discord.AddHandler(messageReceived)

	// Intent
	discord.Identify.Intents = discordgo.MakeIntent(bot.intent)

	// Start
	discord.Open()
	fmt.Printf("%s is now running. Press Ctrl-C to exit.\n", bot.name)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.ChannelMessageSend(Debug.notificationChannel, bot.getResponse(quit))
	discord.Close()
}

func messageReceived(s *discordgo.Session, event *discordgo.MessageCreate) {
	// Ignore messages from self
	if event.Author.ID == s.State.User.ID {
		return
	}
	LogReceivedMessage(s, event.Message, true)
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Started...")
	message, err := s.ChannelMessageSend(Debug.notificationChannel, Wyrm.getResponse(greet))
	if err != nil {
		Log(exception, "Failed to send startup message")
		return
	}
	LogSentMessage(s, message)
}
