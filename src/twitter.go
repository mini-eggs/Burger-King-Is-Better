package BkIsBetter

import (
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	defaultTweetText = "Just saying... Burger King's 10pc. chicken nugget is only $1.49..."
	fileName         = "replies.json"
)

type saveIDs struct {
	TweetIDs []int64
}

func getClient() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}

func createTweet(tweetText *string, params *twitter.StatusUpdateParams) error {
	if *tweetText == "false" {
		return errors.New("No tweet was set")
	}

	_, _, err := getClient().Statuses.Update(*tweetText, params)
	return err
}

func handleSingleTweet(aTweet twitter.Tweet) error {
	hasRow, checkErr := hasRow(aTweet.ID)
	if checkErr != nil {
		return checkErr
	}

	if !hasRow {

		log.Println("Beginning to reply to tweet " + strconv.FormatInt(aTweet.ID, 10))

		insertErr := insertRow(aTweet.ID)
		if insertErr != nil {
			return insertErr
		}

		tweet := "@" + aTweet.User.ScreenName + " " + defaultTweetText
		createTweetErr := createTweet(&tweet, &twitter.StatusUpdateParams{
			InReplyToStatusID: aTweet.ID,
		})
		if createTweetErr != nil {
			return createTweetErr
		}

		log.Println("Replied to tweet: " + strconv.FormatInt(aTweet.ID, 10) + " with `" + tweet + "`")

	} else {

		log.Println("Tweet " + strconv.FormatInt(aTweet.ID, 10) + " has already been replied to")

	}

	return nil
}

func iterateTweets(statuses []twitter.Tweet) error {
	var err = make(chan error)
	errOccurred := false
	var wg sync.WaitGroup
	wg.Add(len(statuses))

	log.Println("Iterating Tweets with goroutine")

	go func(aTweets []twitter.Tweet) {
		for _, singleTweet := range aTweets {
			go func(aTweet twitter.Tweet) {
				anErr := handleSingleTweet(aTweet)
				if anErr != nil {
					log.Println("Error replying to tweet: " + strconv.FormatInt(aTweet.ID, 10))
					errOccurred = true
					err <- anErr
				}
				wg.Done()
			}(singleTweet)
		}
	}(statuses)

	go func() {
		wg.Wait()
		close(err)
	}()

	if errOccurred {
		return <-err
	}

	return nil
}

func queryHashtagAndReply(hashtag *string) error {
	client := getClient()
	if *hashtag == "false" {
		return errors.New("No hashtag was set")
	}
	log.Println("Querying Twitter for hashtag of " + *hashtag)
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#" + *hashtag,
	})
	if err != nil {
		return err
	}

	return iterateTweets(search.Statuses)
}

func queryHashtagLoop(hashtag *string, infinite bool) error {
	if infinite {
		var count int
		count = 0

		for {
			log.Println("Pass number: " + strconv.Itoa(count))
			count = count + 1

			err := queryHashtagAndReply(hashtag)
			if err != nil {
				return err
			}

			time.Sleep(time.Minute * 3)
		}
	} else {
		return queryHashtagAndReply(hashtag)
	}
}
