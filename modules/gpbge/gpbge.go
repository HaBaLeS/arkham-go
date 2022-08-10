package gpbge

import (
	"log"
)

type PhaseEngine struct {
	GameStart  Start
	StartPhase Phase
	Running    bool
}

type Start interface {
	StartGame()
	Callback()
}

type Phase interface {
	Name() string
	Next() Phase
	Execute()
}

func NewPhaseEngine(phase Phase, start Start) *PhaseEngine {
	return &PhaseEngine{
		Running:    true,
		StartPhase: phase,
		GameStart:  start,
	}
}

func (pe *PhaseEngine) Start() {
	pe.GameStart.StartGame()
	current := pe.StartPhase
	for pe.Running {
		log.Printf("Start Execution for Phase: %s", current.Name())
		current.Execute()
		current = current.Next()
	}
}
