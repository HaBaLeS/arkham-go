package gamelogic

import (
	"github.com/HaBaLeS/arkham-go/runtime"
)

func StartCardActivated() {
	i1 := runtime.ScenarioSession().Player[0].Investigator
	i2 := runtime.ScenarioSession().Player[1].Investigator
	runtime.ScenarioSession().StartLocation.ActivateLocation(i1, i2)
}
