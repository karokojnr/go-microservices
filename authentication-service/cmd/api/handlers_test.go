package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authenticate(t *testing.T) {
	jsonToReturn := `
{
	"error": false,
	"message": "message"
}
`
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client
	postBody := map[string]interface{}{
		"email":    "test@test.com",
		"password": "",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)

	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted but got %d", responseRecorder.Code)
	}
}
