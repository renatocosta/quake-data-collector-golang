package domain

type AggregateRoot struct {
	Id     uint64
	Events []Event `json:"-"`
}

func (a *AggregateRoot) RecordThat(event Event) {
	a.Events = append(a.Events, event)
}
