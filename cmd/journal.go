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
		err := journal.Connect().Create(args[0], beginDFlag)
		print(err)
	},
}

var fangCMD = &cobra.Command{
	Use:   "fang",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a journal",
	Long: `Delete a journal. Default journal can't be deleted.
Example: jou fang Didle`,
	Run: func(cmd *cobra.Command, args []string) {
		err := journal.Connect().Remove(args[0])
		print(err)
	},
}

var useCMD = &cobra.Command{
	Use:   "use",
	Args:  cobra.ExactArgs(1),
	Short: "Use journal as default",
	Long: `Use journal as default. Useful for ommiting journal name on certain commands
Example: jou use bubu`,
	Run: func(cmd *cobra.Command, args []string) {
		err := journal.Connect().SetDefault(args[0])
		print(err)
	},
}

var renameCMD = &cobra.Command{
	Use:   "rename",
	Args:  cobra.ExactArgs(2),
	Short: "Rename a journal",
	Long:  `Rename a journal. Example: jou rename oldname newname`,
	Run: func(cmd *cobra.Command, args []string) {
		err := journal.Connect().Rename(args[0], args[1])
		print(err)
	},
}

var openCMD = &cobra.Command{
	Use:   "open",
	Args:  cobra.MaximumNArgs(1),
	Short: "Show journal entries",
	Long: `Show journal entries. Ommit name to open default journal.
Example: jou open, jou open bubu`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			j, err := journal.Connect().GetDefault()
			print(err, j)
		} else {
			j, err := journal.Connect().Get(args[0])
			print(err, j)
		}
	},
}

var listCMD = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "List all journals",
	Long:  `List all journals. Example: jou list`,
	Run: func(cmd *cobra.Command, args []string) {
		j, err := journal.Connect().GetAll()
		print(err, j)
	},
}

func init() {
	rootCMD.AddCommand(beginCMD)
	beginCMD.Flags().BoolVarP(&beginDFlag, "default", "d", false, "use to begin journal as default")
	rootCMD.AddCommand(fangCMD)
	rootCMD.AddCommand(listCMD)
	rootCMD.AddCommand(useCMD)
	rootCMD.AddCommand(openCMD)
	rootCMD.AddCommand(renameCMD)
}
