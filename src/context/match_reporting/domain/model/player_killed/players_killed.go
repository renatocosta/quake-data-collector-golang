package players_killed

var (
	WorldKillerLabel = "world"
)

type PlayerKilled interface {
	Find()
	GetTotal() int64
	Compute(Matchable)
	IsEligibleToBeAPlayer(Matchable) bool
	Consolidate()
	GetPlayers() []string
}

type PlayerKilledEntity struct {
	Total       int64
	PlayerState Player
	Players     []string
}

func NewPlayerKilled(playerState Player) PlayerKilled {
	return &PlayerKilledEntity{
		PlayerState: playerState,
	}
}

func (p *PlayerKilledEntity) Find() {

}

func (p *PlayerKilledEntity) GetTotal() int64 {
	return p.Total
}

func (p *PlayerKilledEntity) Compute(match Matchable) {
	if err := match.Validation(); err != nil {
		panic(err.Error())
	}

	if !p.IsEligibleToBeAPlayer(match) {
		p.PlayerState.KillDown(match)
	}

	p.PlayerState.KillUp(match)
}

func (p *PlayerKilledEntity) IsEligibleToBeAPlayer(match Matchable) bool {
	return match.GetPlayerWhoKilled() != WorldKillerLabel
}

func (p *PlayerKilledEntity) Consolidate() {

}

func (p *PlayerKilledEntity) GetPlayers() []string {
	return p.Players
}
