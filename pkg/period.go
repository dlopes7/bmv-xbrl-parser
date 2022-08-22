package pkg

import "time"

type PeriodType int

const (
	Instant  PeriodType = 1
	Duration            = 2
)

func (p PeriodType) String() string {
	switch p {
	case Instant:
		return "Instant"
	case Duration:
		return "Duration"
	}
	return "Unknown"
}

type Period struct {
	Type      PeriodType `json:"Tipo"`
	Instant   *time.Time `json:"FechaInstante"`
	StartDate *time.Time `json:"FechaInicio"`
	EndDate   *time.Time `json:"FechaFin"`
}

func (p Period) Duration() time.Duration {
	if p.Type != Duration {
		return 0
	}
	return p.EndDate.Sub(*p.StartDate)
}
