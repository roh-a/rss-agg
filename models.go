package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/roh-a/rss-agg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID : dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url    string `json:"url"`
	UserID uuid.UUID  `json:"user_id"`
	
}


func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID : feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
	}
} 


func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	
	feedRes := []Feed{}

	for _,feed := range feeds {
		feedRes = append(feedRes, databaseFeedToFeed(feed))
	}

	return feedRes
} 