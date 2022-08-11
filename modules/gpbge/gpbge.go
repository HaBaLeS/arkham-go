package gpbge

import (
	"fmt"
	"github.com/HaBaLeS/arkham-go/runtime"
	"log"
)

type PhaseEngine struct {
	Players    []*runtime.PlayerDeck
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
		Players:    make([]*runtime.PlayerDeck, 0),
	}
}

func (pe *PhaseEngine) Start() {
	err := pe.validate()
	if err != nil {
		log.Panicf("Phase Engine not Set up correctly! %v", err)
	}
	pe.GameStart.StartGame()
	current := pe.StartPhase
	for pe.Running {
		log.Printf("Start Execution for Phase: %s", current.Name())
		current.Execute()
		current = current.Next()
	}
}

func (pe *PhaseEngine) AddPlayer(player *runtime.PlayerDeck) {
	pe.Players = append(pe.Players, player)
}

func (pe *PhaseEngine) validate() error {
	if len(pe.Players) == 0 {
		return fmt.Errorf("0 players added to PhaseEngine")
	}
	return nil
}
