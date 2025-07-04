package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, `[]`)
	defer cleanDatabase()
	store, _ := NewFileSystemPlayerStore(database)
	
	webserver := NewPlayerServer(store)
	player := "Pepper"

	webserver.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	webserver.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		webserver.ServeHTTP(response, NewGetScoreRequest(player))
		AssertStatus(t, response.Code, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		webserver.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response.Code, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := League{
			{Name: "Pepper", Wins: 3},
		}
		AssertLeague(t, got, want)
	})
}
