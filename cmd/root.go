package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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

// alias jou use to jou journal default set [name]
// alias jou list to jou journal get
// TODO add jou open, which lists all notes and adds
// button for note creation. it also accepts optional journal name

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCMD.Execute()
	print(err)
}

/*
jou add journal ourJournal
jou rename journal myJournal yourJournal
jou delete journal myJournal
jou use journal myJournal
jou open journal [?journalName]
jou list journal

jou add [?note]
*/
