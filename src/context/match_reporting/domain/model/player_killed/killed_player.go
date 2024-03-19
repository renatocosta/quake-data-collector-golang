package players_killed

type KilledPlayer struct {
}

func NewKilledPlayer() Player {
	return &KilledPlayer{}
}

func (k *KilledPlayer) KillUp(match Matchable) {

}

func (k *KilledPlayer) KillDown(match Matchable) {

}

func (k *KilledPlayer) GetPlayers() []string {
	return []string{}
}
