package runtime

import (
	"fmt"
)

type GameState int

const (
	StateVoid GameState = iota
	StateStart
	StateMythos
	StateEnd
)

type GameCommand interface {
	Command() string

	Help() string

	Call(*PlaySession, []string)
}

type GameEntity interface {
	HandleEvent(state *PlaySession, event GameEvent)

	FetchCommands(state *PlaySession) []GameCommand
}

type GameEvent interface {
	Name() string
	Payload() map[string]interface{}
}

type PlaySession struct {
	CardDB         *CardDB
	CurrentState   GameState
	Round          int
	GlobalEntities []GameEntity
	Scenario       *Scenario
}

var globalCommands = map[GameState][]GameCommand{}

func NewGame(db *CardDB) *PlaySession {
	return &PlaySession{
		CardDB:         db,
		CurrentState:   StateVoid,
		Round:          0,
		GlobalEntities: []GameEntity{&GlobalGameController{}},
	}
}
func (ps *PlaySession) Run() {
	for ps.CurrentState != StateEnd {
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
	result = append(result, globalCommands[ps.CurrentState]...)
	// fetch commands from cards
	for _, entity := range ps.GlobalEntities {
		result = append(result, entity.FetchCommands(ps)...)
	}
	// fetch other commands ????
	result = append(result, NewHelpCommand(result))
	return result
}

func (ps *PlaySession) ChangeState(state GameState) {
	if state != ps.CurrentState {
		oldState := ps.CurrentState
		ps.CurrentState = state
		ps.EmitEvent(NewGameEvent("changeState",
			map[string]interface{}{
				"old": oldState,
				"new": state,
			}))
	}
}

func (ps *PlaySession) EmitEvent(event GameEvent) {
	for _, entity := range ps.GlobalEntities {
		entity.HandleEvent(ps, event)
	}

}

func NewHelpCommand(commands []GameCommand) GameCommand {
	return NewCommand(
		"help",
		"you are looking at it now",
		func(ps *PlaySession, args []string) {
			println("help: you are looking at it now")
			for _, cmd := range commands {
				fmt.Printf("%s: %s\n", cmd.Command(), cmd.Help())
			}
		})
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

func (cmd *SimpleCommand) Command() string {
	return cmd.cmd
}

func (cmd *SimpleCommand) Help() string {
	return cmd.hlp
}

func (cmd *SimpleCommand) Call(ps *PlaySession, args []string) {
	cmd.fnc(ps, args)
}

type SimpleEvent struct {
	name    string
	payload map[string]interface{}
}

func NewGameEvent(name string, payload map[string]interface{}) GameEvent {
	return &SimpleEvent{
		name,
		payload,
	}
}

func (se *SimpleEvent) Name() string {
	return se.name
}

func (se *SimpleEvent) Payload() map[string]interface{} {
	return se.payload
}
