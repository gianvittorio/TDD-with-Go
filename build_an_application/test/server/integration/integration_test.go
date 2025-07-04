package integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build_an_application/src/domain"
	"example.com/build_an_application/src/server"
	"example.com/build_an_application/src/store"
	"example.com/build_an_application/test/server/assertions"
)

// server_integration_test.go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	webserver := server.NewPlayerServer(store)
	player := "Pepper"

	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		webserver.ServeHTTP(response, server.NewGetScoreRequest(player))
		assertions.AssertStatus(t, response.Code, http.StatusOK)

		assertions.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		webserver.ServeHTTP(response, server.NewLeagueRequest())
		assertions.AssertStatus(t, response.Code, http.StatusOK)

		got := server.GetLeagueFromResponse(t, response.Body)
		want := []domain.Player{
			{Name: "Pepper", Wins: 3},
		}
		assertions.AssertLeague(t, got, want)
	})
}
