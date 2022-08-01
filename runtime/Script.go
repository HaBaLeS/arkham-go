package runtime

import (
	"errors"
	"fmt"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"reflect"
	"regexp"
)

type Script struct {
	interpreter  *interp.Interpreter
	EventMethods map[string]EventHandlerMethod
}

var Symbols = map[string]map[string]reflect.Value{
	"runtime/runtime": {
		"GameState":   reflect.ValueOf((*GameState)(nil)),
		"GameCommand": reflect.ValueOf((*GameCommand)(nil)),
		"GameEntity":  reflect.ValueOf((*GameEntity)(nil)),
		"GameEvent":   reflect.ValueOf((*GameEvent)(nil)),
		"PlaySession": reflect.ValueOf((*PlaySession)(nil)),
	},
}

type EventHandlerMethod = func(*PlaySession, GameEvent)

var methodHandlerRegex = regexp.MustCompile("^Handle([A-Za-z]+)Event$")

func NewScript(script string) (*Script, error) {
	inter := interp.New(interp.Options{})
	err := inter.Use(stdlib.Symbols)
	if err != nil {
		return nil, err
	}
	err = inter.Use(Symbols)
	if err != nil {
		return nil, err
	}
	_, err = safeEval(inter, script)
	if err != nil {
		return nil, err
	}
	eventMethods := make(map[string]EventHandlerMethod, 0)
	symbols := inter.Symbols("game")
	for name, val := range symbols["game"] {
		iface := val.Interface()
		switch meth := iface.(type) {
		case EventHandlerMethod:
			if submatch := methodHandlerRegex.FindStringSubmatch(name); submatch != nil {
				eventMethods[submatch[1]] = meth
			}
			break
		default:
			// Nil
		}
	}
	return &Script{
		inter,
		eventMethods,
	}, nil
}

func (script *Script) ReadConstant(nane string) (reflect.Value, error) {
	return safeEval(script.interpreter, "game."+nane)

}

func (script *Script) CallEventIfExists(session *PlaySession, event GameEvent) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("Event Handler Error", r))
		}
	}()
	if handler := script.EventMethods[event.Name()]; handler != nil {
		handler(session, event)
	}
	return
}

func safeEval(inter *interp.Interpreter, script string) (result reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint("Script parse Error", r))
		}
	}()
	result, err = inter.Eval(script)
	return
}
