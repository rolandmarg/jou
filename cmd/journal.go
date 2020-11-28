package cmd

import (
	"github.com/spf13/cobra"
)

var journalCMD = &cobra.Command{
	Use:   "journal",
	Short: "Create a journal",
	Long:  `jou journal - Create a journal. Examples: jou journal Secret, jou journal 'My thought palace'`,
	Run:   func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCMD.AddCommand(journalCMD)
}
