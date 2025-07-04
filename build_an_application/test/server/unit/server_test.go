package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build_an_application/src/domain"
	"example.com/build_an_application/src/server"
	"example.com/build_an_application/test/server/assertions"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []domain.Player
}

func (s *StubPlayerStore) GetLeague() []domain.Player {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// server_test.go
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	webserver := server.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := server.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		assertions.AssertStatus(t, response.Code, http.StatusOK)
		assertions.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := server.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		assertions.AssertStatus(t, response.Code, http.StatusOK)
		assertions.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := server.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertions.AssertStatus(t, got, want)
	})
}

// server_test.go
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	webserver := server.NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := server.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		assertions.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []domain.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		webserver := server.NewPlayerServer(&store)

		request := server.NewLeagueRequest()
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		got := server.GetLeagueFromResponse(t, response.Body)
		assertions.AssertStatus(t, response.Code, http.StatusOK)
		assertions.AssertLeague(t, got, wantedLeague)
		assertions.AssertContentType(t, response, server.JsonContentType)
	})
}
