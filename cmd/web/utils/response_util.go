package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func WriteAPIErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	jsonResp, err := json.Marshal(APIResponse{
		Success: false,
		Data:    nil,
		Error:   err.Error(),
	})
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

func WriteAPISuccessResponse(w http.ResponseWriter, r *http.Request, response interface{}) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	logger.Info("success", "method", r.Method, "uri", r.URL.RequestURI())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}
