package BkIsBetter

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

// Initialize - start the application.
func Initialize() error {

	// Get command line variables
	programType := flag.String("type", "default", "What work is there to be done.")
	tweetText := flag.String("text", "false", "Tweet text")
	hashtag := flag.String("hashtag", "mcdonalds", "Hashtag")
	infinite := flag.String("infinite", "false", "Hashtag")
	flag.Parse()

	// Get environment port
	port := os.Getenv("PORT")

	// Run program
	switch *programType {

	case "tweet":
		return createTweet(tweetText, &twitter.StatusUpdateParams{})

	case "query":
		return queryHashtagLoop(hashtag, *infinite != "false")

	case "default":
		if port == "" {
			return errors.New("no program type has been specified")
		}
		// We're on a heroku server.
		// Execute default application.
		go func() {
			serverErr := queryHashtagLoop(hashtag, true)
			if serverErr != nil {
				log.Fatal(serverErr)
			}
		}()
		// Listen on port
		return http.ListenAndServe(":"+port, nil)

	default:
		return errors.New("no program type has been specified")
	}
}
