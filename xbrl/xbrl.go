package xbrl

import (
	"encoding/json"
	"time"
)

type XBRL struct {
	Contexts map[string]*Context `json:"ContextosPorId"`
	Facts    map[string]*Fact    `json:"HechosPorId"`
	Units    map[string]*Unit    `json:"UnidadesPorId"`
}

func ParseXBRLData(data []byte) (*XBRL, error) {
	var x XBRL
	if err := json.Unmarshal(data, &x); err != nil {
		return &x, err
	}
	for _, context := range x.Contexts {
		context.Facts = x.factsByContext(context.ID)
	}

	return &x, nil
}

func (x XBRL) factsByContext(contextID string) []Fact {
	var facts []Fact
	for _, fact := range x.Facts {
		if fact.ContextID == contextID {
			facts = append(facts, *fact)
		}
	}
	return facts
}

func (x XBRL) QuarterContexts() []Context {
	var contexts []Context
	for _, context := range x.Contexts {

		// We say that this is a quarter if the duration is between 80 amd 100 days
		if context.Period.Duration() >= 80*24*time.Hour && context.Period.Duration() < 100*24*time.Hour {
			contexts = append(contexts, *context)
		}
	}
	return contexts
}
