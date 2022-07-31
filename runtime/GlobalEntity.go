package runtime

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GlobalGameController struct {
}

func (controller *GlobalGameController) HandleEvent(ps *PlaySession, event GameEvent) {

}

func (controller *GlobalGameController) FetchCommands(ps *PlaySession) []GameCommand {
	if ps.CurrentState == StateVoid {
		return []GameCommand{
			NewCommand("list", "list all available scenarios", listScenarios),
			NewCommand("scenario", "Load a specific scenaria", loadScenario),
			NewCommand("script", "load and execute a script", executeScript),
		}
	}
	return []GameCommand{}
}

func listScenarios(ps *PlaySession, args []string) {
	fmt.Printf("Available Scenarios:\n")
	for _, name := range ScenarioList() {
		fmt.Printf(" - %s\n", name)
	}
	println()
}

func loadScenario(ps *PlaySession, args []string) {
	scenario, err := LoadScenario(args[0])
	if err != nil {
		fmt.Printf("Unable to load scenario: %s\n", err.Error())
	}
	ps.Scenario = scenario
	scenario.Activate(ps)
}

func executeScript(ps *PlaySession, args []string) {
	f, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("Error loading script %s\n", args[0])
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
OUTER:
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			cmd := ps.fetchUserStateCommands()
			results := strings.Split(line, " ")
			for _, c := range cmd {
				if c.Command() == results[0] {
					fmt.Printf("Executing: \"%s\"\n", line)
					c.Call(ps, results[1:])
					continue OUTER
				}
			}
			println("Command not found !")
		}
	}
}
