package main

import (
	_ "fmt"
	"time"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELETOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Handle("/twitter", func(m *tb.Message) {
		config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
		token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		client := twitter.NewClient(httpClient)

		// Send a Tweet
		//tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)
		_, _, err := client.Statuses.Update("tweet from go-twitter", nil)

		if err != nil {
			log.Fatal(err)
			b.Send(m.Sender, "tweet failed")
			return
		}

		b.Send(m.Sender, "tweet sent")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
	})
	
	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
	})
	
	b.Handle(tb.OnChannelPost, func (m *tb.Message) {
		// channel posts only
	})
	

	b.Start()
}