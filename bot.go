package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kiliankoe/openmensa"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var canteen *openmensa.Canteen

func main() {
	if len(os.Args) <= 0 {
		println("Please add token as cli argument.")
		return
	}
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", os.Args[1]))
	canteens, err := openmensa.GetCanteens(175)
	if err != nil {
		fmt.Println("error creating discord bot: ", err)
		return
	}
	dg.AddHandler(messageCreate)
	canteen = canteens[0]

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!mensa" {
		sendMealsForDate(s, time.Now(), m.ChannelID)
	}
}

func sendMealsForDate(s *discordgo.Session, t time.Time, channel string) {
	msg, err := s.ChannelMessageSend(channel, "Just a second...")
	if err != nil {
		println("error sending message: ", err)
		return
	}
	meals, err := canteen.GetMeals(t)
	if err != nil {
		s.ChannelMessageEdit(channel, msg.ID, "An error occured :angry:")
		return
	}
	var messages = make([]*discordgo.MessageEmbedField, len(meals))
	var footer = ""
	for i, meal := range meals {
		messages[i] = &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s: %s", meal.Category, meal.Name),
			Value: fmt.Sprintf("%.2f€", *meal.Prices.Students),
		}
		if len(meal.Notes) > 0 {
			footer += fmt.Sprintf("[%d] %s\n", i + 1, strings.Join(meal.Notes, ", "))
		}
	}
	empty := ""
	s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("** :spaghetti: Menü vom %s**", t.Format("02.01.2006")),
			Fields: messages,
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: "https://pbs.twimg.com/profile_images/643755515118505984/xzZMK7fU_400x400.png",
				Text: footer,
			},
		},
		ID: msg.ID,
		Channel: msg.ChannelID,
		Content: &empty,
	})
}