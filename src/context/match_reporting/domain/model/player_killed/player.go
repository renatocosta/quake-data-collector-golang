package players_killed

type Player interface {
	KillUp(match Matchable)
	KillDown(match Matchable)
	GetPlayers() []string
}
