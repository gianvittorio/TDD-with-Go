package store

import "example.com/build_an_application/src/domain"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []domain.Player
}
