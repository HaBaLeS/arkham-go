package engine

import "fmt"

type GameState int

const (
	StateStart GameState = iota
	StateMythos
	StateEnd
)

var globalCommands = map[GameState][]GameCommand{
	StateStart: {NewCommand("start",
		"start <scenario> # starts the scenario of the given name",
		startScenario),
		NewCommand("exit",
			"exit the game",
			exitScenario),
	},
}

type GameCommand interface {
	command() string

	help() string

	call(*Engine, []string)
}

type Engine struct {
	currentState GameState
	stateStack   []GameState
	round        int
}

func NewGame() *Engine {
	return &Engine{
		currentState: StateStart,
		stateStack:   make([]GameState, 0, 10),
		round:        0,
	}
}

func (engine *Engine) Run() {
	for engine.currentState != StateEnd {
		stateCommands := engine.fetchUserStateCommands()
		engine.prompt(stateCommands)
	}
}

/**
Get list of commands an interactive Player can do.
*/
func (engine *Engine) fetchUserStateCommands() []GameCommand {
	result := make([]GameCommand, 0)
	// fetch global state commands
	result = append(result, globalCommands[engine.currentState]...)
	// fetch commands from cards
	// fetch other commands ????
	result = append(result, NewHelpCommand(result))
	return result
}

func NewHelpCommand(commands []GameCommand) GameCommand {
	return &SimpleCommand{
		"help",
		"you are looking at it now",
		func(engine *Engine, args []string) {
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
	fnc func(*Engine, []string)
}

func NewCommand(cmd string, help string, fnc func(*Engine, []string)) GameCommand {
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

func (cmd *SimpleCommand) call(engine *Engine, args []string) {
	cmd.fnc(engine, args)
}
