package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/build_an_application/src/server"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
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
		[]string{},
	}
	webserver := &server.PlayerServer{Store: &store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := server.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		server.AssertStatus(t, response.Code, http.StatusOK)
		server.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := server.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		server.AssertStatus(t, response.Code, http.StatusOK)
		server.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := server.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		server.AssertStatus(t, got, want)
	})
}

// server_test.go
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	webserver := &server.PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := server.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		webserver.ServeHTTP(response, request)

		server.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}
