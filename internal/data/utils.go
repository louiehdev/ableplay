package data

import (
	"context"
	"encoding/json"
	"fmt"
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

func ParseQueryParams(values url.Values) map[string]interface{} {
	queries := make(map[string]interface{}, len(values))

	for key := range values {
		if value := values.Get(key); value != "" {
			if i, err := strconv.Atoi(value); err == nil {
				queries[key] = int32(i)
			} else {
				queries[key] = string(value)
			}
		}
	}

	if limit, ok := queries["limit"]; !ok {
		queries["limit"] = int32(50) // Set limit to default if none provided
	} else {
		if num, ok := limit.(int32); ok {
			if num < 0 {
				queries["limit"] = int32(50)
			} else if num > 100 {
				queries["limit"] = int32(100)
			}
		}
	}

	return queries
}

func GetContextValue[T any](ctx context.Context, contextKey any) (T, error) {
	value, ok := ctx.Value(contextKey).(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("no values found, check key or data type")
	}

	return value, nil
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
