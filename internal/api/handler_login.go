package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/louiehdev/ableplay/internal/auth"
	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var params data.AddUserParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if params.FirstName == "" || params.LastName == "" || params.Email == "" || params.Password == "" {
		data.RespondWithError(w, http.StatusBadRequest, "Missing essential user metadata")
		return
	}

	hashedPassword, err := auth.CreateHash(params.Password)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Cannot hash password, "+err.Error())
		return
	}
	params.Password = hashedPassword

	user, err := api.DB.AddUser(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to add user to database, "+err.Error())
		return
	}
	apiKey, err := generateAPIKey(r.Context(), api.DB, user)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "could not create API Key, "+err.Error())
		return
	}
	userData := struct{
		User data.User 
		APIKey string}{
			User: user, 
			APIKey: apiKey}

	data.RespondWithJSON(w, http.StatusOK, userData)
}

func (api *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	var params data.UpdateUserParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if params.Password != "" {
		hashedPassword, err := auth.CreateHash(params.Password)
		if err != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Cannot hash password, "+err.Error())
			return
		}
		params.Password = hashedPassword
	}

	user, err := api.DB.UpdateUser(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to update user in database, "+err.Error())
		return
	}

	data.RespondWithJSON(w, http.StatusOK, user)
}

func (api *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) { // Add raw API key to cookie for frontend use, make sure to set SameSite and HTTPOnly and Secure
	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := api.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		data.RespondWithError(w, http.StatusUnauthorized, "Incorrect username or password, "+err.Error())
		return
	}

	if match, err := auth.CheckHash(params.Password, user.Password); !match {
		data.RespondWithError(w, http.StatusUnauthorized, "Incorrect password, "+err.Error())
		return
	}

	userData := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
		Email     string `json:"email"`
	}{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Email:     user.Email,
	}

	data.RespondWithJSON(w, http.StatusOK, userData)
}

func (api *apiConfig) handlerCreateAPIKey(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetContextValue[data.User](r.Context(), userContextKey)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Incorrect user data provided, could not create API Key, "+err.Error())
		return
	}

	apiKey, err := generateAPIKey(r.Context(), api.DB, user)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "could not create API Key, "+err.Error())
		return
	}

	data.RespondWithJSON(w, http.StatusCreated, apiKey)
}

func generateAPIKey(ctx context.Context, db *data.Queries, user data.User) (string, error) {
	apiKey, keyPrefix := auth.CreateAPIKey()
	hashedKey, err := auth.CreateHash(apiKey)
	if err != nil {
		return "", err
	}

	if _, err := db.CreateKey(ctx, data.CreateKeyParams{ID: keyPrefix, ApiKey: hashedKey, UserID: user.ID}); err != nil {
		return "", err
	}

	return apiKey, nil
}