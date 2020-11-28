package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "jou",
	Args:  cobra.MaximumNArgs(4),
	Short: "Your private thought palace",
	Long:  `jou is a CLI app that archives your thoughts locally, no network involved.`,
}

// journal related commands

// alias jou use to jou journal default set [name]
var journalSetDefaultCMD = &cobra.Command{
	Use:   "use",
	Args:  cobra.ExactArgs(1),
	Short: "set default journal",
	Long:  `jou use [name] - set journal for default operations: TODO list. Example: jou use secret`,
	Run: func(cmd *cobra.Command, args []string) {
		journalSetDefault(args)
	},
}

// alias jou list to jou journal get
var journalGetCMD = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "list journals",
	Long:  `jou list - list journals. Default journal will always be first in the list. Example: jou list`,
	Run: func(cmd *cobra.Command, args []string) {
		journalGetAll()
	},
}

// TODO add jou open, which lists all entries and adds
// button for entry creation. it also accepts optional journal name

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCMD.AddCommand(journalCMD)
	// journalCMD.AddCommand(journalGetCMD)

	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/*
jou add journal ourJournal
jou rename journal myJournal yourJournal
jou delete journal myJournal
jou use journal myJournal
jou open journal [?journalName]
jou list journal

jou add [?entry]
*/
