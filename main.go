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
	discord.AddHandler(ready)

	// Start
	discord.Open()
	fmt.Printf("%s is now running. Press Ctrl-C to exit.", bot.name)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}

func tick(s *discordgo.Session) {
	//
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Started...")
}
