package players_killed

type DeathPlayer struct {
	amount int
}

func NewDeathPlayer() PlayerState {
	return &DeathPlayer{}
}

func (k *DeathPlayer) ComputeKills(amount int) {
	k.amount = amount - 1
}

func (k *DeathPlayer) GetKills() int {
	return k.amount
}
