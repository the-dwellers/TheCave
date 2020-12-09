package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/wsxiaoys/terminal"
)

type logType int

const (
	debug     logType = iota
	info      logType = iota
	warning   logType = iota
	exception logType = iota
)

// Log to console
func Log(t logType, message string) {

	switch t {
	case debug:
		terminal.Stdout.Colorf("@{.w}")
	case info:
		terminal.Stdout.Colorf("@c")
	case warning:
		terminal.Stdout.Colorf("@y")
	case exception:
		terminal.Stdout.Colorf("@r")
	}
	terminal.Stdout.Print(message).Nl().Reset()
}

// LogSentMessage : Log a message sent by the bot
func LogSentMessage(session *discordgo.Session, message *discordgo.Message) {
	prefix := getGuildName(session, message.GuildID) + "#" + getChannelName(session, message.ChannelID) + ":< "
	terminal.Stdout.Colorf("@m" + prefix + message.Content).Nl().Reset()
}

func getGuildName(session *discordgo.Session, guildID string) string {
	guild, err := session.Guild(guildID)
	if err != nil {
		return guildID
	}
	return guild.Name
}

func getChannelName(session *discordgo.Session, channelID string) string {
	channel, err := session.Channel(channelID)
	if err != nil {
		return channelID
	}
	return channel.Name
}

// LogReceivedMessage : Log a message received by the bot
func LogReceivedMessage(session *discordgo.Session, message *discordgo.Message, interest bool) {
	prefix := getGuildName(session, message.GuildID) + "#" + getChannelName(session, message.ChannelID) + "@" + message.Author.Username + ":> "

	if interest {
		Log(info, prefix+message.Content)
	}
}
