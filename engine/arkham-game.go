package engine

import (
	"github.com/HaBaLeS/arkham-go/command"
	"github.com/HaBaLeS/arkham-go/gamelogic"
	"log"
	"sync"
)

type ExecFunc func()

type ArkhamStart struct {
	wg            *sync.WaitGroup
	wgResolveFunc ExecFunc
}

func (as *ArkhamStart) StartGame() {
	log.Printf("Game Start. Waiting for Player to start the Game")
	as.wg.Add(1)
	as.wg.Wait()
}

func (as *ArkhamStart) Callback() {
	as.wgResolveFunc()
	as.wg.Done()
}

type ArkhamPhase struct {
	name       string
	next       Phase
	execfunc   ExecFunc
	engineChan chan command.EngineCommand
}

func (ap *ArkhamPhase) Name() string {
	return ap.name
}

func (ap *ArkhamPhase) Next() Phase {
	return ap.next
}

func (ap *ArkhamPhase) Execute() {
	ap.execfunc()
}

func (ap *ArkhamPhase) SetNext(next Phase) {
	ap.next = next
}

func BuildArkhamGame() *PhaseEngine {

	//Start of Game
	start := &ArkhamStart{
		wg: &sync.WaitGroup{},
	}
	start.wgResolveFunc = gamelogic.StartCardActivated

	//Phases

	mp := CreateMythosPhase()
	ip := CreateInvestigationPhase()
	ep := CreateEnemyPhase()
	up := CreateUpkeepPhase()

	mp.SetNext(ip)
	ip.SetNext(ep)
	ep.SetNext(up)
	up.SetNext(mp)

	//Put together
	engine := NewPhaseEngine(ip, start)

	return engine //Game starts with Investigation, first mythos is skipped
}
