package engine

import (
	"log"
	"time"
)

func CreateInvestigationPhase() *ArkhamPhase {
	return &ArkhamPhase{
		name:     "Investigation",
		execfunc: DoInvestigate,
	}
}

func DoInvestigate() {
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
	time.Sleep(500 * time.Minute)

}
