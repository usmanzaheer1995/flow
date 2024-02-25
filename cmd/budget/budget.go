package budget

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/spf13/cobra"
)

// budgetCmd represents the budget command
var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "Manage your budget",
	Long: `Budget command allows users to manage their budgetary allocations 
for different spending categories. With this command, you can create, view, 
and adjust their budgets to effectively track and control their expenses.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nothing specified, nothing added.")
	},
}

func init() {
	cmd.RootCmd.AddCommand(budgetCmd)
	// Added subcommands
	budgetCmd.AddCommand(createCmd)
	budgetCmd.AddCommand(viewCmd)
	budgetCmd.AddCommand(adjustCmd)
	budgetCmd.AddCommand(removeCmd)
}
