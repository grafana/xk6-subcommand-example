// Package example contains the xk6-subcommand-example extension.
package example

import (
	"github.com/spf13/cobra"
	"go.k6.io/k6/cmd/state"
)

// newSubcommand creates the "example" subcommand for the xk6 extension.
// This subcommand serves as a parent for additional commands, such as "greet".
// It demonstrates how to structure an xk6 extension with custom subcommands.
// The "example" subcommand itself does not perform any actions directly,
// but provides a namespace for related commands.
// You can add more subcommands to this parent command as needed.
// The greet command is included as an example of a functional subcommand.
func newSubcommand(gs *state.GlobalState) *cobra.Command {
	subCmd := &cobra.Command{
		Use:   "example",
		Short: "An example xk6 subcommand",
		Long: `An example xk6 subcommand demonstrating custom subcommand creation.

This subcommand includes a 'greet' command that prints greeting messages to the console,
showcasing how to add custom functionality. You can extend it with additional commands
as needed.

This parent command has no standalone functionality - use its subcommands to perform
actions, though you can add direct functionality if desired.
`,
	}

	subCmd.AddCommand(newGreetCmd(gs))

	return subCmd
}
