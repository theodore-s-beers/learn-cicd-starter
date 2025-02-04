package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	// Capture original log output
	originalLogOutput := log.Default().Writer()
	defer log.Default().SetOutput(originalLogOutput) // Restore log output after test

	// Redirect log output to buffer to suppress printing
	log.Default().SetOutput(&bytes.Buffer{})

	// Use a test HTTP response recorder
	recorder := httptest.NewRecorder()

	respondWithError(recorder, http.StatusInternalServerError, "Something went wrong")
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, recorder.Code)
	}

	var response map[string]string
	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	expectedMessage := "Something went wrong"
	if response["error"] != expectedMessage {
		t.Errorf("Expected error message %q, got %q", expectedMessage, response["error"])
	}
}
