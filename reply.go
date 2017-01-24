package twcli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/mitchellh/cli"
	"github.com/yyoshiki41/go-twitter/twitter"
)

type replyCommand struct {
	ui cli.Ui
}

func (c *replyCommand) Run(args []string) int {
	var query string
	var count int

	f := flag.NewFlagSet("search", flag.ContinueOnError)
	f.StringVar(&query, "query", "", "query")
	f.StringVar(&query, "q", "", "query")
	f.IntVar(&count, "count", 15, "count")
	f.IntVar(&count, "c", 15, "count")
	f.Usage = func() { c.ui.Error(c.Help()) }
	if err := f.Parse(args); err != nil {
		return 1
	}

	if query == "" {
		c.ui.Error("query is empty.")
		return 1
	}

	var userID int64 = 751227608755499008

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: query,
		Count: count,
	})
	if err != nil {
		c.ui.Error(fmt.Sprintf(
			"Failed to search: %s", err))
	}
	if resp.StatusCode != 200 {
		c.ui.Error(fmt.Sprintf(
			"search: StatusCode %d", resp.StatusCode))
	}

	followers, resp, err := client.Followers.IDs(&twitter.FollowerIDParams{
		UserID: userID,
		Count:  5000,
	})
	if err != nil {
		c.ui.Error(fmt.Sprintf(
			"Failed to followers/ids: %s", err))
	}
	if resp.StatusCode != 200 {
		c.ui.Error(fmt.Sprintf(
			"followers/ids: StatusCode %d", resp.StatusCode))
	}

	for _, t := range search.Statuses {
		// TODO: 送信済みユーザーは弾く

		for _, i := range followers.IDs {
			// follower の場合のみ、DM送信
			if i == t.User.ID {
				/*
								dm, resp, err := client.DirectMessages.New(&twitter.DirectMessageNewParams{
									UserID: t.User.ID,
									Text:   "",
								})

								if err != nil {
									log.Fatalf("%s", err)
								}
								if resp.StatusCode != 200 {
									log.Fatalf("status code %d", resp.StatusCode)
								}
					// TODO: 今日送信したユーザーをファイル書き出し
				*/
				break
			}
		}
	}

	return 0
}

func (c *replyCommand) Synopsis() string {
	return "search twitter"
}

func (c *replyCommand) Help() string {
	return strings.TrimSpace(`
Usage: twcli search [options]
  search API
Options:
  -query,q=    search query
  -count,c=    count
`)
}
