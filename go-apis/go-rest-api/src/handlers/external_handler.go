package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/src/config"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// ErrorResponse represents an error response structure
type ErrorResponse struct {
	Error       string `json:"error"`
	Details     string `json:"details,omitempty"`
	RequestID   string `json:"request_id,omitempty"`
	RawResponse string `json:"raw_response,omitempty"`
}

type PingResponse struct {
	Message string `json:"message"`
}

func GetExternalHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	requestID := r.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = fmt.Sprintf("req-%d", time.Now().UnixNano())
	}

	fullURL := cfg.ExternalAPIURL + "/"
	log.Printf("[%s] Attempting to call external API: %s", requestID, fullURL)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(fullURL)
	if err != nil {
		handleError(w, err, "Failed to reach external API", requestID, http.StatusInternalServerError, "")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		handleError(w, fmt.Errorf("status code: %d", resp.StatusCode),
			"External API returned non-200 status code",
			requestID, resp.StatusCode, "")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handleError(w, err, "Failed to read response body", requestID, http.StatusInternalServerError, "")
		return
	}

	// Log raw response for debugging
	log.Printf("[%s] Raw response body: %s", requestID, string(body))

	if len(body) == 0 {
		handleError(w, fmt.Errorf("empty response"),
			"External API returned empty response",
			requestID, http.StatusInternalServerError, "")
		return
	}

	// Validate content type and response format
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		handleError(w, fmt.Errorf("invalid content type: %s", contentType),
			"External API returned non-JSON response",
			requestID, http.StatusBadGateway, string(body))
		return
	}

	var pingResp PingResponse
	if err := json.Unmarshal(body, &pingResp); err != nil {
		handleError(w, err, "Failed to parse external API response",
			requestID, http.StatusInternalServerError, string(body))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", requestID)
	if err := json.NewEncoder(w).Encode(pingResp); err != nil {
		handleError(w, err, "Error encoding response",
			requestID, http.StatusInternalServerError, "")
		return
	}
}

func handleError(w http.ResponseWriter, err error, message string, requestID string, statusCode int, rawResponse string) {
	errMsg := fmt.Sprintf("%s: %v", message, err)
	log.Printf("[%s] Error: %s", requestID, errMsg)
	if rawResponse != "" {
		log.Printf("[%s] Raw response: %s", requestID, rawResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", requestID)
	w.WriteHeader(statusCode)

	errorResp := ErrorResponse{
		Error:       message,
		Details:     err.Error(),
		RequestID:   requestID,
		RawResponse: rawResponse,
	}
	json.NewEncoder(w).Encode(errorResp)
}
