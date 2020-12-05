package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/rolandmarg/jou/internal/app/journal"
)

func print(err error, args ...interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(args) != 0 {
		fmt.Println(args...)
	}
}

var rootCMD = &cobra.Command{
	Use:   "jou",
	Short: "Your private thought palace",
	Long:  `jou is a CLI app that archives your thoughts locally, no network involved.`,
}
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

var addCMD = &cobra.Command{
	Use:   "add",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Add a note",
	Long: `Add a note. Ommit journal name for default journal.
Examples: jou add "goes to default journal", jou add secretj "deep note"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			err := journal.Connect().CreateNote(args[0], args[1])
			print(err)
		} else {
			err := journal.Connect().CreateNote("", args[0])
			print(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCMD.AddCommand(beginCMD)
	beginCMD.Flags().BoolVarP(&beginDFlag, "default", "d", false, "use to begin journal as default")
	rootCMD.AddCommand(fangCMD)
	rootCMD.AddCommand(listCMD)
	rootCMD.AddCommand(useCMD)
	rootCMD.AddCommand(openCMD)
	rootCMD.AddCommand(addCMD)

	err := rootCMD.Execute()
	print(err)
}
