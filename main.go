package main

import (
	//"fmt"
	"os"
	"log"

	bot "dmAssistGo/bot"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	discordKey := os.Getenv("Discord_Key")
	//dmChannel := os.Getenv("DM_Channel")

	bot.BotToken = discordKey
	bot.Run()
	
	
}
