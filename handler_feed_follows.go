package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amr9876/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user *database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	decoder.Decode(&params)

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldnt create feed follow: %v", err))
		return
	}

	respondWithJson(w, http.StatusCreated, databaseFeedFollowToFeedFollow(&feed))

}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user *database.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldnt get feed follows: %v", err))
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedFollowsToFeedFollows(&feedFollow))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user *database.User) {

	feedFollowId := chi.URLParam(r, "feedFollowID")

	err := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     uuid.MustParse(feedFollowId),
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldnt delete feed follow: %v", err))
		return
	}

	respondWithJson(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "Deleted Successfully"})

}
