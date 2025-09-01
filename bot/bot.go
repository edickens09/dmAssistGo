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
	discord.AddHandler(RollInitative)

	discord.Open()

	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
func RollInitative(discord *discordgo.Session, message *discordgo.MessageCreate){

	if message.Author.ID == discord.State.User.ID {
		return
	}
	switch {

	case strings.Contains(message.Content, "!initative"): 

		discord.ChannelMessageSend(message.ChannelID, "How many combatants?")
		
		responseChan := make(chan *discordgo.MessageCreate)
		var handler func (*discordgo.Session, *discordgo.MessageCreate)
		handler = func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
			if msg.Author.Bot {
				return
			}

			if msg.Author.ID == message.Author.ID {
				responseChan <- msg
			}
		}

		removeHandler := discord.AddHandler(handler)

		reply := <- responseChan
		discord.ChannelMessageSend(message.ChannelID, "Succesful integration "+ reply.Content)

		removeHandler()
	}
}
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	
	case strings.Contains(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "Hello here are some options")

	case strings.Contains(message.Content, "!standin"):
		discord.ChannelMessageSend(message.ChannelID, "This is a standin option")
	}
}
