package players_killed

type PlayerState interface {
	ComputeKills(amount int)
	GetKills() int
}
