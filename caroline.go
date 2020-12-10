package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/mikkyang/id3-go"
)

type carolineBot struct {
	bot CaveBot
}

type fullFileInfo struct {
	file   os.FileInfo
	folder string
}

var buffer = make([][]byte, 0)

// Caroline : stuff
var Caroline = carolineBot{CarolineBot}

func playJetSet(s *discordgo.Session, vc *discordgo.VoiceConnection) {
	// todo: convert to generic folders playing function.
	// todo: define playlists in data.go
	// todo: load playlist from file
	fileinfos1, err := ioutil.ReadDir("E:/Audio/Music/2_Mello/Memories Of Tokyo-To An Ode To Jet Set Radio")
	if err != nil {
		Log(warning, "Unable to read files in folder"+err.Error())
	}
	var files []fullFileInfo
	for _, f := range fileinfos1 {
		files = append(files, fullFileInfo{f, "E:/Audio/Music/2_Mello/Memories Of Tokyo-To An Ode To Jet Set Radio"})
	}

	filesinfos2, err := ioutil.ReadDir("E:/Audio/Music/Jet Set Radio/Jet Set Radio OST")
	if err != nil {
		Log(warning, "Unable to read files in folder"+err.Error())
	}

	for _, f := range filesinfos2 {
		files = append(files, fullFileInfo{f, "E:/Audio/Music/Jet Set Radio/Jet Set Radio OST"})
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })

	playFiles(s, vc, files)
}

func playFiles(s *discordgo.Session, vc *discordgo.VoiceConnection, files []fullFileInfo) {
	for _, f := range files {
		if f.file.Name()[len(f.file.Name())-3:len(f.file.Name())] != "mp3" {
			Log(info, "Not mp3 file: "+f.file.Name())
			continue
		}
		fpath := filepath.Clean(f.folder + "/" + f.file.Name())
		playFile(s, vc, fpath)
	}
}

func playFile(s *discordgo.Session, vc *discordgo.VoiceConnection, fpath string) {
	mp3file, err := id3.Open(fpath)
	Song := filepath.Base(fpath)
	if err != nil {
		Log(warning, "No ID3 tag in "+fpath)
	} else if mp3file.Title() != "" {
		Song = mp3file.Title()[0 : len(mp3file.Title())-1]
	}
	Log(info, "Playing "+Song+"("+fpath+")")
	s.ChannelMessageSend(Debug.notificationChannel, "Now Playing: **"+Song+"**")
	if !vc.Ready {
		return
	}
	dgvoice.PlayAudioFile(vc, fpath, make(chan bool))
}

func playFolder(s *discordgo.Session, vc *discordgo.VoiceConnection, folder string) {
	filesinfo, err := ioutil.ReadDir(folder)
	if err != nil {
		Log(warning, "Unable to read files in folder"+err.Error())
	}

	var files []fullFileInfo
	for _, f := range filesinfo {
		files = append(files, fullFileInfo{f, folder})
	}

	playFiles(s, vc, files)

}

func (bot carolineBot) messageReceived(s *discordgo.Session, event *discordgo.MessageCreate) {
	// Ignore messages from self
	if event.Author.ID == s.State.User.ID {
		return
	}
	LogReceivedMessage(s, event.Message, true)
	if strings.Contains(event.Message.Content, "car") {
		vs, err := s.State.VoiceState(event.GuildID, event.Message.Author.ID)
		if err != nil {
			Log(warning, "Unable to get user voice state:"+err.Error())
		}

		vc, err := s.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
		if err != nil {
			Log(warning, "Unable to connect to voice channel "+err.Error())
			return
		}

		if strings.Contains(event.Message.Content, "jet") {
			playJetSet(s, vc)
		} else if strings.Contains(event.Message.Content, "wii") {
			playFile(s, vc, "E:/Development/Gits/TheCave/wii.mp3")
		}

		vc.Disconnect()
	}
}
