package data

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(jsonData)
}

func RespondWithError(w http.ResponseWriter, code int, errMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	if len(errMessage) == 0 {
		w.WriteHeader(code)
		return
	}
	errData, err := json.Marshal(errorResponse{Error: errMessage})
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(errData)
}

func GetRequestUUID(r *http.Request, idType string) (uuid.UUID, error) {
	id := r.PathValue(idType)
	reqUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return reqUUID, nil
}

func ToPgtypeText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}

func ToPgtypeInt4(stringInt string) pgtype.Int4 {
	i, err := strconv.Atoi(stringInt)
	if err != nil {
		return pgtype.Int4{Valid: false}
	}

	return pgtype.Int4{Int32: int32(i), Valid: true}
}

func RemoveEmptyValues(stringSlice []string) []string {
	var cleanSlice []string
	for _, value := range stringSlice {
		if value != "" {
			cleanSlice = append(cleanSlice, value)
		}
	}
	return cleanSlice
}

func IsChecked(s string) bool {
	return s != ""
}
