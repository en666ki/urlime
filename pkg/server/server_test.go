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

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

type ServerSuite struct {
	suite.Suite
	server Server
}

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) TestNoRouter() {
	s.Error(s.server.Start("8080"), "Runinng server w/o router")
}

func (s *ServerSuite) TestAddHandler() {
	tests := map[string]struct {
		method  string
		path    string
		handler LimeHandler
	}{
		"/get": {
			method: "GET",
			path:   "/get",
			handler: func(body []byte) ([]byte, int, error) {
				return body, http.StatusOK, nil
			},
		},
		"/post": {
			method: "POST",
			path:   "/post",
			handler: func(body []byte) ([]byte, int, error) {
				return nil, 0, errors.New("error")
			},
		},
	}

	for _, tt := range tests {
		s.server.AddHandler(tt.method, tt.path, tt.handler)
		walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			route = strings.ReplaceAll(route, "/*/", "/")
			s.Equal(tests[route].path, route)
			s.Equal(tests[route].method, method)
			return nil
		}
		err := chi.Walk(s.server.router, walkFunc)
		s.NoError(err)
	}
	s.server.router = nil
}

func (s *ServerSuite) TestRunHandler() {
	var server Server
	mockEchoHandler := func(body []byte) ([]byte, int, error) {
		return body, http.StatusOK, nil
	}

	server.AddHandler("GET", "/", mockEchoHandler)
	go server.Start("8080")

	time.Sleep(100 * time.Millisecond)

	rr, err := http.Get("http://localhost:8080")
	s.NoError(err, "GET request")
	defer rr.Body.Close()

	s.Equal(http.StatusOK, rr.StatusCode, "Check 200")
}

func TestEchoServer(t *testing.T) {
	mockEchoHandler := createHttpHandler(func(body []byte) ([]byte, int, error) {
		return body, http.StatusOK, nil
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

func TestBrokenServer(t *testing.T) {
	mockBrokenHandler := createHttpHandler(func(body []byte) ([]byte, int, error) {
		return nil, 0, errors.New("Oops! The server gremlins struck again")
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
	assert.Equal(t, http.StatusText(http.StatusInternalServerError)+"\n", string(gotBody), "Check body")
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
			path: "/echo_s",
			givenHandler: func(body []byte) ([]byte, int, error) {
				return body, http.StatusOK, nil
			},
			givenReader: strings.NewReader("Ping"),
			thenStatus:  http.StatusOK,
			thenBody:    "Ping",
		},
		{
			name: "Failure echo handelr",
			path: "/echo_ise",
			givenHandler: func(body []byte) ([]byte, int, error) {
				return nil, 0, errors.New("Gremlins were in handler")
			},
			givenReader: strings.NewReader("Ping"),
			thenStatus:  http.StatusInternalServerError,
			thenBody:    http.StatusText(http.StatusInternalServerError) + "\n",
		},
		{
			name: "Reader failure echo handelr",
			path: "/echo_br",
			givenHandler: func(body []byte) ([]byte, int, error) {
				return body, 0, nil
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
