package runtime

import (
	"arkham-go/card"
	"arkham-go/runtime/scenarios"
	"strings"
)

type Scenario struct {
	card   *card.Scenario
	script *Script
}

func (scene *Scenario) Activate(ps *PlaySession) {
	scene.script.CallEventIfExists(ps, NewGameEvent("Activate", map[string]interface{}{
		// Umm
	}))
}

func ScenarioList() []string {
	entries, err := scenarios.Scenarios.ReadDir(".")
	if err != nil {
		panic(err)
	}
	result := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := strings.TrimSuffix(entry.Name(), ".scenario")
		result = append(result, name)
	}
	return result
}

func LoadScenario(ps *PlaySession, name string) (*Scenario, error) {
	data, err := scenarios.Scenarios.ReadFile(name + ".scenario")
	if err != nil {
		return nil, err
	}
	script, err := NewScript(string(data))
	if err != nil {
		return nil, err
	}

	idValue, err := script.ReadConstant("Card")
	if err != nil {
		return nil, err
	}
	var crd card.ArkhamCard = nil
	switch id := idValue.Interface().(type) {
	case string:
		crd = ps.CardDB.cards[id]
	}
	return &Scenario{
		crd.(*card.Scenario),
		script,
	}, nil
}
