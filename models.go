package main

import (
	"time"

	"github.com/amr9876/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser *database.User) User {
	return User{
		ID:        dbUser.ID,
		UpdatedAt: dbUser.UpdatedAt,
		CreatedAt: dbUser.CreatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed *database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		UpdatedAt: dbFeed.UpdatedAt,
		CreatedAt: dbFeed.CreatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds *[]database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range *dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(&dbFeed))
	}

	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow *database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		CreatedAt: dbFeedFollow.CreatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows *[]database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range *dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(&dbFeedFollow))
	}

	return feedFollows
}
