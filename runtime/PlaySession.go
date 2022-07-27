package runtime

import "fmt"

type GameState int

const (
	StateStart GameState = iota
	StateMythos
	StateEnd
)

type GameCommand interface {
	command() string

	help() string

	call(*PlaySession, []string)
}

var globalCommands = map[GameState][]GameCommand{
	StateStart: {NewCommand("start",
		"start <scenario> # starts the scenario of the given name",
		startScenario),
		NewCommand("exit",
			"exit the game",
			exitScenario),
	},
}

type PlaySession struct {
	currentState GameState
	stateStack   []GameState
	round        int
}

func NewGame() *PlaySession {
	return &PlaySession{
		currentState: StateStart,
		stateStack:   make([]GameState, 0, 10),
		round:        0,
	}
}
func (ps *PlaySession) Run() {
	for ps.currentState != StateEnd {
		stateCommands := ps.fetchUserStateCommands()
		ps.prompt(stateCommands)
	}
}

/**
Get list of commands an interactive Player can do.
*/
func (ps *PlaySession) fetchUserStateCommands() []GameCommand {
	result := make([]GameCommand, 0)
	// fetch global state commands
	result = append(result, globalCommands[ps.currentState]...)
	// fetch commands from cards
	// fetch other commands ????
	result = append(result, NewHelpCommand(result))
	return result
}

func NewHelpCommand(commands []GameCommand) GameCommand {
	return &SimpleCommand{
		"help",
		"you are looking at it now",
		func(ps *PlaySession, args []string) {
			println("help: you are looking at it now")
			for _, cmd := range commands {
				fmt.Printf("%s: %s\n", cmd.command(), cmd.help())
			}
		},
	}
}

type SimpleCommand struct {
	cmd string
	hlp string
	fnc func(*PlaySession, []string)
}

func NewCommand(cmd string, help string, fnc func(*PlaySession, []string)) GameCommand {
	return &SimpleCommand{
		cmd,
		help,
		fnc,
	}
}

func (cmd *SimpleCommand) command() string {
	return cmd.cmd
}

func (cmd *SimpleCommand) help() string {
	return cmd.hlp
}

func (cmd *SimpleCommand) call(ps *PlaySession, args []string) {
	cmd.fnc(ps, args)
}

func (ps *PlaySession) AddPlayer(name string, deck *Deck) {}

func (ps *PlaySession) Init(scenario *Scenario) {
	//Create a chaos Bag
	//Create AgendaDeck
	//Create ActDeck
	//GetLocation Cards
	//Create EncounterDeck
}
