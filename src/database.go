package BkIsBetter

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type twitterReplies struct {
	TweetID int64 `gorm:"size:255"`
}

func connect() (*gorm.DB, error) {
	var connectionInfo string
	connectionInfo = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD")
	connectionInfo = connectionInfo + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	connectionInfo = connectionInfo + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"

	return gorm.Open("mysql", connectionInfo)
}

func hasRow(tweetID int64) (bool, error) {
	db, connectErr := connect()
	if connectErr != nil {
		return true, connectErr
	}

	var count int
	queryErr := db.Table("twitter_replies").Where("tweet_id = ?", tweetID).Count(&count).Error
	if queryErr != nil {
		return true, queryErr
	}

	defer db.Close()

	return count != 0, nil
}

func insertRow(tweetID int64) error {
	db, connectErr := connect()
	if connectErr != nil {
		return connectErr
	}

	queryErr := db.Create(&twitterReplies{TweetID: tweetID}).Error
	if queryErr != nil {
		return queryErr
	}

	defer db.Close()

	return nil
}
