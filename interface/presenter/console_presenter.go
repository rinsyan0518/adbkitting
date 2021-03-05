package presenter

import "github.com/spf13/cobra"

// ConsolePresenter is writable console output
type ConsolePresenter interface {
	Printfln(format string, i ...interface{})
	Errorfln(format string, i ...interface{})
}

type consolePresenter struct {
	cmd *cobra.Command
}

// NewConsolePresenter returns new presenter
func NewConsolePresenter(cmd *cobra.Command) ConsolePresenter {
	return &consolePresenter{
		cmd: cmd,
	}
}

func (p consolePresenter) Printfln(format string, i ...interface{}) {
	p.cmd.Printf(format, i...)
	p.cmd.Println()
}

func (p consolePresenter) Errorfln(format string, i ...interface{}) {
	p.cmd.PrintErrf(format, i...)
	p.cmd.PrintErrln()
}
