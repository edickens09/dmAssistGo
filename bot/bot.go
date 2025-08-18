package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Println(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()

	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	
	case strings.Contains(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "Hello World!")

	case strings.Contains(message.Content, "!bye"):
		discord.ChannelMessageSend(message.ChannelID, "Hello World!")
	}
}
