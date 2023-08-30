package main

import (
	"fmt"
	"net/http"

	"github.com/amr9876/rss-aggregator/internal/auth"
	"github.com/amr9876/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, *database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Invalid api key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Couldnt get user: %v", err))
			return
		}

		handler(w, r, &user)

	}

}
