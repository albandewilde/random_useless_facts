package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

var discordToken string

func init() {
	// Read the discord token
	discordToken = os.Getenv("discord")
}

func main() {
	// Create discord bot instance
	bot, err := dgo.New("Bot " + discordToken)

	if err != nil {
		log.Fatal(err)
	}

	// Register function callback
	bot.AddHandler(fact)
	bot.AddHandler(help)

	// Open bot connection
	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func getFactFromBody(body []byte) string {
	content := strings.Split(string(body), "\n")

	return content[0]
}

func getFact(url string) string {

	res, err := http.Get(url)

	var fact string
	if err != nil {
		fact = err.Error()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fact = err.Error()
	} else {
		var err error
		fact = getFactFromBody(body)
		if err != nil {
			fact = err.Error()
		}
	}

	return fact
}

func fact(s *dgo.Session, m *dgo.MessageCreate) {

	// Check if the content of a message is a command for us
	if !strings.HasPrefix(m.Content, "°") {
		return
	}

	// Possible url
	urlRandomFact := "https://uselessfacts.jsph.pl/random.txt?language=en"
	urlTodayFact := "https://uselessfacts.jsph.pl/today.txt?language=en"
	// Choose the url depending on the command
	var url string
	if m.Content == "°random" || m.Content == "°fact" {
		url = urlRandomFact
	} else if m.Content == "°today" {
		url = urlTodayFact
	}

	// May the command is unknow, so we don't have any url
	if url != "" {
		s.ChannelMessageSend(m.ChannelID, getFact(url))
	}
}

func help(s *dgo.Session, m *dgo.MessageCreate) {
	if m.Content == "°help" {
		content := "```\n" +
			"- °help → Display this help message\n" +
			"- °random or °fact → Display a random useless fact\n" +
			"- °today → Display the random useless fact of the day\n" +
			"```"

		s.ChannelMessageSend(m.ChannelID, content)
	}
}
