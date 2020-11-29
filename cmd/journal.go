package cmd

import (
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
		journal.Connect().Create(args[0], beginDFlag)
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
