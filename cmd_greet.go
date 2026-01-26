package example

import (
	"github.com/spf13/cobra"
	"go.k6.io/k6/cmd/state"
)

// newGreetCmd creates the "greet" subcommand for the "example" xk6 extension.
// This command prints a customizable greeting message to the console.
// It demonstrates how to add functional subcommands to an xk6 extension.
// Users can specify a name to greet using the --name flag.
func newGreetCmd(_ *state.GlobalState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greet",
		Short: "Prints a greeting message",
		Long: `Prints a customizable greeting message to the console.

Customize the greeting by specifying a name with the --name flag.
`,
		Example: "  k6 x example greet --name Alice",
	}

	name := cmd.Flags().StringP("name", "n", "User", "Name of the person to greet")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		println("Hello, " + *name + "! Happy subcommanding with k6!")
		return nil
	}

	return cmd
}
