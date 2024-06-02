package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VSM1le/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldn't get feed %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
		Url:      params.URL,
		UserID:   user.ID,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldnt create user %v", err))
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
