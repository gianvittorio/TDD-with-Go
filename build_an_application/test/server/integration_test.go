package integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build_an_application/src/server"
	"example.com/build_an_application/src/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	webserver := server.PlayerServer{store}
	player := "Pepper"

	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	webserver.ServeHTTP(response, server.NewGetScoreRequest(player))
	server.AssertStatus(t, response.Code, http.StatusOK)
	server.AssertResponseBody(t, response.Body.String(), "3")
}
