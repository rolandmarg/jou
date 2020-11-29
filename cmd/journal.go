package cmd

import (
	"fmt"
	"os"

	"github.com/rolandmarg/jou/internal/app/journal"
	"github.com/spf13/cobra"
)

var beginDFlag bool
var beginCMD = &cobra.Command{
	Use:   "begin",
	Args:  cobra.ExactArgs(1),
	Short: "Begin a new journal",
	Long: `Begin a new journal, specify -d(--default) to make it default.
Examples: jou begin books, jou begin -d mix`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := journal.Connect().Create(args[0], beginDFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)       
			os.Exit(1)
		}
	},
}

var listCMD = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "List all journals",
	Long:  `List all journals. Example: jou list`,
	Run: func(cmd *cobra.Command, args []string) {
		journal.Connect().GetAll()
	},
}

func init() {
	rootCMD.AddCommand(beginCMD)
	beginCMD.Flags().BoolVarP(&beginDFlag, "default", "d", false, "use to begin journal as default")
	rootCMD.AddCommand(listCMD)
}
