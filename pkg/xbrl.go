package pkg

import "time"

type XBRL struct {
	Contexts map[string]Context `json:"ContextosPorId"`
	Facts    map[string]Fact    `json:"HechosPorId"`
	Units    map[string]Unit    `json:"UnidadesPorId"`
}

func (x XBRL) QuarterContexts() []Context {
	var contexts []Context
	for _, context := range x.Contexts {

		// We say that this is a quarter if the duration is between 80 amd 100 days
		if context.Period.Duration() >= 80*24*time.Hour && context.Period.Duration() < 100*24*time.Hour {
			contexts = append(contexts, context)
		}
	}
	return contexts
}
