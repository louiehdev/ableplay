package api

import (
	"context"
	"net/http"

	"github.com/louiehdev/ableplay/internal/auth"
	"github.com/louiehdev/ableplay/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

func (api *apiConfig) RequireRole(minRole string) func(http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			// 1. Extract API key
			rawKey, prefixKey, err := auth.GetAPIKey(r)
			if err != nil {
				data.RespondWithError(w, http.StatusInternalServerError, "Invalid API Key, "+err.Error())
				return
			}
			hashedKey, err := api.DB.GetKeyByID(r.Context(), prefixKey)
			if err != nil {
				data.RespondWithError(w, http.StatusInternalServerError, "Invalid API Key, "+err.Error())
				return
			}
			match, err := auth.CheckHash(rawKey, hashedKey)
			if !match || err != nil {
				data.RespondWithError(w, http.StatusInternalServerError, "Invalid API Key, "+err.Error())
				return
			}

			// 2. Look up user
			user, err := api.DB.GetUserByAPIKey(r.Context(), hashedKey)
			if err != nil {
				data.RespondWithError(w, http.StatusInternalServerError, "No user attached to provided key, "+err.Error())
				return
			}

			// 3. Check role
			if auth.RoleLevel[user.Role] < auth.RoleLevel[minRole] {
				data.RespondWithError(w, http.StatusForbidden, "User does not have permission")
				return
			}

			// 4. Add user to context
			ctx := context.WithValue(r.Context(), userContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
