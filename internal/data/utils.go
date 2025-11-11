package data

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"
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

func ParseQueryParams(values url.Values) (limit int32) {
	// TO-DO: Allow for greater amount of params and return correct value types

	// Example: queries := make(map[string]interface{}, len(values))

	limit = 50

	if limitStr := values.Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = int32(l)
		}
	}

	if limit > 100 {
		limit = 100
	}

	return limit
}

func ToNullableText(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ToNullableInt(stringInt string) *int32 {
	i, err := strconv.Atoi(stringInt)
	if err != nil {
		return nil
	}
	nInt := int32(i)
	return &nInt
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
