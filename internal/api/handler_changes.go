package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerAddGameChange(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetContextValue[data.User](r.Context(), userContextKey)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Insufficient user data provided, "+err.Error())
		return
	}

	var params data.AddGamesChangeParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	params.UserID = user.ID

	if err := api.DB.AddGamesChange(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerAddFeatureChange(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetContextValue[data.User](r.Context(), userContextKey)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Insufficient user data provided, "+err.Error())
		return
	}

	var params data.AddFeaturesChangeParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	params.UserID = user.ID

	if err := api.DB.AddFeaturesChange(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerAddGameFeatureChange(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetContextValue[data.User](r.Context(), userContextKey)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Insufficient user data provided, "+err.Error())
		return
	}

	var params data.AddGamesFeaturesChangeParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	params.UserID = user.ID

	if err := api.DB.AddGamesFeaturesChange(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerGetGamesChanges(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	gamesChanges, err := api.DB.GetGamesChanges(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve potential games from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, gamesChanges)
}

func (api *apiConfig) handlerGetFeaturesChanges(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	featuresChanges, err := api.DB.GetFeaturesChanges(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve potential features from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, featuresChanges)
}

func (api *apiConfig) handlerGetGamesFeaturesChanges(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	gamesfeaturesChanges, err := api.DB.GetGamesFeaturesChanges(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve potential games features connections from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, gamesfeaturesChanges)
}

func (api *apiConfig) handlerSubmitChanges(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetContextValue[data.User](r.Context(), userContextKey)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Insufficient user data provided, "+err.Error())
		return
	}

	var params []data.SubmitChangeParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong, no changes found to submit")
		return
	}

	for _, change := range params {
		change.UserID = user.ID
		if err := submitChange(r.Context(), api.Conn, api.DB, change); err != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong, could not submit change, "+err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func submitChange(ctx context.Context, db *pgxpool.Pool, queries *data.Queries, params data.SubmitChangeParams) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	switch params.Category {
	case "game":
		if err := submitGameChange(ctx, qtx, params); err != nil {
			return err
		}
	case "feature":
		if err := submitFeatureChange(ctx, qtx, params); err != nil {
			return err
		}
	case "gamefeature":
		if err := submitGameFeatureChange(ctx, qtx, params); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func submitGameChange(ctx context.Context, qtx *data.Queries, params data.SubmitChangeParams) error {
	updateData, ok := params.Data.(data.UpdateGameParams)
	if !ok {
		return errors.New("incorrect data for submission type")
	}
	switch params.Type {
	case "add":
		if err := qtx.SubmitGamesChange(ctx, params.ChangeID); err != nil {
			return err
		}
	case "update":
		if err := qtx.UpdateGame(ctx, updateData); err != nil {
			return err
		}
	case "delete":
		if err := qtx.DeleteGame(ctx, params.ChangeID); err != nil {
			return err
		}
	}

	if err := qtx.UpdateGamesChange(ctx, data.UpdateGamesChangeParams{ID: params.ChangeID, Status: params.Status, ModeratorID: params.UserID}); err != nil {
		return err
	}
	return nil
}

func submitFeatureChange(ctx context.Context, qtx *data.Queries, params data.SubmitChangeParams) error {
	updateData, ok := params.Data.(data.UpdateFeatureParams)
	if !ok {
		return errors.New("incorrect data for submission type")
	}
	switch params.Type {
	case "add":
		if err := qtx.SubmitFeaturesChange(ctx, params.ChangeID); err != nil {
			return err
		}
	case "update":
		if err := qtx.UpdateFeature(ctx, updateData); err != nil {
			return err
		}
	case "delete":
		if err := qtx.DeleteFeature(ctx, params.ChangeID); err != nil {
			return err
		}
	}

	if err := qtx.UpdateFeaturesChange(ctx, data.UpdateFeaturesChangeParams{ID: params.ChangeID, Status: params.Status, ModeratorID: params.UserID}); err != nil {
		return err
	}
	return nil
}

func submitGameFeatureChange(ctx context.Context, qtx *data.Queries, params data.SubmitChangeParams) error {
	updateData, ok := params.Data.(data.UpdateGameFeatureParams)
	if !ok {
		return errors.New("incorrect data for submission type")
	}
	switch params.Type {
	case "add":
		if err := qtx.SubmitGamesFeaturesChange(ctx, params.ChangeID); err != nil {
			return err
		}
	case "update":
		if err := qtx.UpdateGameFeature(ctx, updateData); err != nil {
			return err
		}
	case "delete":
		if err := qtx.DeleteGameFeature(ctx, data.DeleteGameFeatureParams{GameID: updateData.GameID, FeatureID: updateData.FeatureID}); err != nil {
			return err
		}
	}

	if err := qtx.UpdateGamesFeaturesChange(ctx, data.UpdateGamesFeaturesChangeParams{ID: params.ChangeID, Status: params.Status, ModeratorID: params.UserID}); err != nil {
		return err
	}
	return nil
}
