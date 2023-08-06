package server

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunHandler(t *testing.T) {
	mockEchoHandler := func(body []byte) ([]byte, error) {
		return body, nil
	}

	RunHandler(mockEchoHandler, "/", "8080")

	time.Sleep(100 * time.Millisecond)

	rr, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	defer rr.Body.Close()

	assert.Equal(t, http.StatusOK, rr.StatusCode, "Check 200")
}

func TestEchoServer(t *testing.T) {
	mockEchoHandler := createHttpHandler(func(body []byte) ([]byte, error) {
		return body, nil
	})

	server := httptest.NewServer(mockEchoHandler)
	defer server.Close()

	message := "Ping"
	req := bytes.NewBuffer([]byte(message))

	res, err := http.Post(server.URL, "text/plain", req)

	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode, "Check 200")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	assert.Equal(t, message, string(body), "Check echo server")
}

func TestBrokenServ(t *testing.T) {
	mockBrokenHandler := createHttpHandler(func(body []byte) ([]byte, error) {
		return nil, errors.New("Oops! The server gremlins struck again")
	})

	server := httptest.NewServer(mockBrokenHandler)
	defer server.Close()

	message := "Hello there!"
	r := bytes.NewReader([]byte(message))

	response, err := http.Post(server.URL, "application/json", r)
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer response.Body.Close()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode, "Check 500")

	gotBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	assert.Equal(t, http.StatusText(http.StatusInternalServerError) + "\n", string(gotBody), "Check body")
}

func TestCreateHttpHandler(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		givenHandler LimeHandler
		givenReader  io.Reader
		thenStatus   int
		thenBody     string
	}{
		{
			name: "Success echo handelr",
			path: "/echo",
			givenHandler: func(body []byte) ([]byte, error) {
				return body, nil
			},
			givenReader: strings.NewReader("Ping"),
			thenStatus:  http.StatusOK,
			thenBody:    "Ping",
		},
		{
			name: "Failure echo handelr",
			path: "/echo",
			givenHandler: func(body []byte) ([]byte, error) {
				return nil, errors.New("Gremlins were in handler")
			},
			givenReader: strings.NewReader("Ping"),
			thenStatus:  http.StatusInternalServerError,
			thenBody:    http.StatusText(http.StatusInternalServerError) + "\n",
		},
		{
			name: "Success echo handelr",
			path: "/echo",
			givenHandler: func(body []byte) ([]byte, error) {
				return body, nil
			},
			givenReader: iotest.ErrReader(errors.New("Gremlins broke request")),
			thenStatus:  http.StatusBadRequest,
			thenBody:    http.StatusText(http.StatusBadRequest) + "\n",
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("POST", tt.path, tt.givenReader)

		recorder := httptest.NewRecorder()

		httpHandler := createHttpHandler(tt.givenHandler)

		http.HandlerFunc(httpHandler).ServeHTTP(recorder, req)

		assert.Equal(t, tt.thenStatus, recorder.Code, "Check code")

		assert.Equal(t, tt.thenBody, recorder.Body.String(), "Check body")
	}
}
