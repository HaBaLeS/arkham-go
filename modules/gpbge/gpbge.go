package gpbge

import (
	"log"
)

type PhaseEngine struct {
	StartPhase Phase
	Running    bool
}

type Phase interface {
	Name() string
	Next() Phase
	Execute()
}

func NewPhaseEngine(start Phase) *PhaseEngine {
	return &PhaseEngine{
		Running:    true,
		StartPhase: start,
	}
}

func (pe *PhaseEngine) Start() {
	current := pe.StartPhase
	for pe.Running {
		log.Printf("Start Execution for Phase: %s", current.Name())
		current.Execute()
		current = current.Next()
	}
}
