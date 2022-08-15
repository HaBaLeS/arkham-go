package engine

import (
	"github.com/HaBaLeS/arkham-go/card"
	"github.com/HaBaLeS/arkham-go/command"
	"github.com/HaBaLeS/arkham-go/runtime"
	"log"
)

func CreateInvestigationPhase() *ArkhamPhase {
	ap := &ArkhamPhase{
		name:       "Investigation",
		engineChan: make(chan command.EngineCommand),
	}

	ap.execfunc = ap.DoInvestigate

	return ap
}

func (ap *ArkhamPhase) DoInvestigate() {
	log.Printf("\t Choose who Starts")
	log.Printf("\t Do 3 Actions of:")
	log.Printf("\t\t Draw Card")
	log.Printf("\t\t Get Ressource")
	log.Printf("\t\t Engage Enemy")
	log.Printf("\t\t Investigate")
	log.Printf("\t\t Move")
	log.Printf("\t\t Play a Card")
	log.Printf("\t\t Escape")
	log.Printf("\t\t Fight")
	log.Printf("\t (if engaged with enemy and NOT Escape, Fight (or special Card  'verhandeln', 'aufgeben') then there is a Gelegenheitsangriff)")
	log.Printf("Waiiting for commands to come")

	command.SetEngineChannel(ap.engineChan)
	for true {
		//repeat until we ceveive a command to return
		cmd := <-ap.engineChan
		log.Printf("Recived Engine Command: %v", cmd)
		switch v := cmd.(type) {
		case command.DoInvestigate:
			doInvestigate(v)
		default:
			log.Panicf("Not handling: %s", v)
		}
	}

}

func doInvestigate(v command.DoInvestigate) {
	log.Printf("\t %s is investigating in %s", v.Investigator, v.Location)
	location := card.AcAsLocation(runtime.CardDBG().GetCard(v.Location))
	location.RemoveClueTokensOnCard(1)
}
