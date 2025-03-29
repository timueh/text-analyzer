package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewLettersRoute(t *testing.T) {
	data := `{"data": "eehe"}`
	req := httptest.NewRequest(http.MethodPost, "/foo", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	HandleLetterFrequency(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got %d, want %d", w.Code, http.StatusOK)
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	wantJson := `[{"name":"e","value":3},{"name":"h","value":1}]`
	if strings.TrimSpace(string(body)) != strings.TrimSpace(wantJson) {
		t.Errorf("got %s, want %s", string(body), wantJson)
	}

}
