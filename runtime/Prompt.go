package runtime

import (
	"github.com/peterh/liner"
	"strings"
)

func (ps *PlaySession) prompt(cmd []GameCommand) {
	prompt := liner.NewLiner()
	defer prompt.Close()
	prompt.SetCtrlCAborts(true)
	prompt.SetWordCompleter(NewCommandCompleter(cmd))
	if line, err := prompt.Prompt("GAME> "); err != nil {
		panic(err)
	} else {
		if line == "" {
			line = "help"
		}
		results := strings.Split(line, " ")
		for _, c := range cmd {
			if c.Command() == results[0] {
				c.Call(ps, results[1:])
				return
			}
		}
		println("Command not found !")
	}
}

func NewCommandCompleter(cmd []GameCommand) liner.WordCompleter {
	return func(line string, pos int) (head string, completions []string, tail string) {
		head = line[:pos]
		tail = line[pos:]
		if strings.Contains(line, " ") {
			return head, []string{}, tail
		}
		completions = make([]string, 0)
		for _, c := range cmd {
			if strings.HasPrefix(c.Command(), head) {
				completions = append(completions, c.Command()[pos:])
			}
		}
		return head, completions, tail
	}
}
