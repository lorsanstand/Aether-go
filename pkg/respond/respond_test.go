package respond

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestRespondJSON(t *testing.T) {
	r := Respond{}
	w := httptest.NewRecorder()
	status := 201
	data := map[string]string{"test": "hi"}

	r.RespondJSON(w, status, data)

	if w.Code != status {
		t.Errorf("Invalid status code inital: %v final: %v", status, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Error("Format is not application/json")
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed parsing JSON: %v", err)
	}
}
