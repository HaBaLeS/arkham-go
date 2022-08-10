package arkham_game

import (
	"github.com/HaBaLeS/arkham-go/card"
	"github.com/HaBaLeS/arkham-go/modules/gpbge"
	"github.com/HaBaLeS/arkham-go/runtime"
	"log"
	"sync"
)

type ExecFunc func()

type ArkhamStart struct {
	wg            *sync.WaitGroup
	wgResolveFunc ExecFunc
	startLocation *card.Location
}

func (as *ArkhamStart) StartGame() {
	log.Printf("Game Start. Waiting for Player to start the Game")
	as.wg.Add(1)
	as.wg.Wait()
}

func (as *ArkhamStart) Callback() {
	//enable player
	//flip card startlocation
	as.startLocation.ActivateLocation()
	as.wg.Done()
}

type ArkhamPhase struct {
	name     string
	next     gpbge.Phase
	execfunc ExecFunc
}

func (ap *ArkhamPhase) Name() string {
	return ap.name
}

func (ap *ArkhamPhase) Next() gpbge.Phase {
	return ap.next
}

func (ap *ArkhamPhase) Execute() {
	ap.execfunc()
}

func (ap *ArkhamPhase) SetNext(next gpbge.Phase) {
	ap.next = next
}

func BuildArkhamGame(scnData *runtime.ScenarioData) *gpbge.PhaseEngine {

	//Start of Game
	start := &ArkhamStart{
		wg:            &sync.WaitGroup{},
		startLocation: scnData.StartLocation,
	}
	start.wgResolveFunc = start.Callback

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
	engine := gpbge.NewPhaseEngine(ip, start)

	return engine //Game starts with Investigation, first mythos is skipped
}
