package engine

import (
	"github.com/peterh/liner"
	"strings"
)

func (engine *Engine) prompt(cmd []GameCommand) {
	prompt := liner.NewLiner()
	defer prompt.Close()
	prompt.SetWordCompleter(NewCOmmandCompleter(cmd))
	if line, err := prompt.Prompt("GAME> "); err != nil {
		panic(err)
	} else {
		if line == "" {
			line = "help"
		}
		results := strings.Split(line, " ")
		for _, c := range cmd {
			if c.command() == results[0] {
				c.call(engine, results[1:])
				return
			}
		}
		println("Command not found !")
	}
}

func NewCOmmandCompleter(cmd []GameCommand) liner.WordCompleter {
	return func(line string, pos int) (head string, completions []string, tail string) {
		head = line[:pos]
		tail = line[pos:]
		if strings.Contains(line, " ") {
			return head, []string{}, tail
		}
		completions = make([]string, 0)
		for _, c := range cmd {
			if strings.HasPrefix(c.command(), head) {
				completions = append(completions, c.command()[pos:])
			}
		}
		return head, completions, tail
	}
}
