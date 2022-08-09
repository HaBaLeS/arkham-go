package arkham_game

import (
	"github.com/HaBaLeS/arkham-go/modules/gpbge"
	"log"
	"time"
)

type ExecFunc func()

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

func BuildArkhamPhases() gpbge.Phase {
	mp := CreateMythosPhase()

	ip := CreateInvestigationPhase()

	ep := CreateEnemyPhase()

	up := CreateUpkeepPhase()

	mp.SetNext(ip)
	ip.SetNext(ep)
	ep.SetNext(up)
	up.SetNext(mp)

	return ip //Game starts with Investigation, first mythos is skipped
}

func DefaultExecFunc() {
	log.Printf("Executing Phase internally")
	time.Sleep(500 * time.Millisecond)
}
