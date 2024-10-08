package main

import (
	"fmt"
	"net/http"

	"github.com/DoctrineAlanK/rssagg/internal/database"
	"github.com/DoctrineAlanK/rssagg/internal/database/auth"
)

type authedHandle func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApikey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
