package store_test

import (
	"testing"

	"example.com/build_an_application/src/domain"
	"example.com/build_an_application/src/store"
	"example.com/build_an_application/test/assertions"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := store.NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []domain.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		assertions.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertions.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := store.NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33
		assertions.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := store.NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertions.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := store.NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertions.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := store.NewFileSystemPlayerStore(database)

		assertions.AssertNoError(t, err)
	})

	//file_system_store_test.go
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := store.CreateTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemPlayerStore(database)

		assertions.AssertNoError(t, err)

		got := store.GetLeague()

		want := domain.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		assertions.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertions.AssertLeague(t, got, want)
	})
}
