package runtime

import (
	"arkham-go/runtime/scenarios"
	"strings"
)

type Scenario struct {
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

func LoadScenario(name string) (*Scenario, error) {
	data, err := scenarios.Scenarios.ReadFile(name + ".scenario")
	if err != nil {
		return nil, err
	}
	script, err := NewScript(string(data))
	if err != nil {
		return nil, err
	}
	return &Scenario{
		script,
	}, nil
}
