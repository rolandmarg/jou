package cmd

import (
	"github.com/rolandmarg/jou/internal/app/note"
	"github.com/spf13/cobra"
)

var addCMD = &cobra.Command{
	Use:   "add",
	Args:  cobra.RangeArgs(1, 2),
	Short: "Add a note",
	Long: `Add a note. Ommit journal name for default journal.
Examples: jou add "goes to default journal", jou add secretj "deep note"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			err := note.Connect().Create(args[0], args[1])
			print(err)
		} else {
			err := note.Connect().CreateDefault(args[0])
			print(err)
		}
	},
}

func init() {
	rootCMD.AddCommand(addCMD)
}
