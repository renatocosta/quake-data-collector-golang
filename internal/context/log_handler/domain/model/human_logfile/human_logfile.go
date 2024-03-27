package human_logfile

type HumanLogFile interface {
	AddRow(row HumanLogFileRowable)
	GetTotalKills() int
	GetRows() []HumanLogFileRowable
}

func NewHumanLogFile() HumanLogFile {
	return &HumanLogFileEntity{}
}

type HumanLogFileEntity struct {
	Rows []HumanLogFileRowable
}

func (h *HumanLogFileEntity) AddRow(row HumanLogFileRowable) {

	row.Validation()
	h.Rows = append(h.Rows, HumanLogFileRow{
		playerWhoKilled: row.GetPlayerWhoKilled(),
		playerWhoDied:   row.GetPlayerWhoDied(),
		meanOfDeath:     row.GetMeanOfDeath(),
	})
}

func (h *HumanLogFileEntity) GetTotalKills() int {
	return len(h.Rows)
}

func (h *HumanLogFileEntity) GetRows() []HumanLogFileRowable {
	return h.Rows
}
