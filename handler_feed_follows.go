package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VSM1le/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreatFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FEEDID uuid.UUID `json:"feedid"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		FeedID:   params.FEEDID,
		UserID:   user.ID,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldnt create feed follow %v", err))
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg *apiConfig) handlerSelectFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feed, err := apiCfg.DB.SelectFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldnt create feed follow %v", err))
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feed))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldnt delete feed follow %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithERROR(w, 400, fmt.Sprintf("couldnt delete feed follow %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
