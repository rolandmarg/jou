package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
		journalCreate(args, beginDFlag)
	},
}

var removeJFlag bool
var removeCMD = &cobra.Command{
	Use:        "remove",
	SuggestFor: []string{"delete", "erase"},
	Args:       cobra.ExactArgs(1),
	Short:      "Remove a note or journal",
	Long: `Remove a note or journal. Flags [-j(--journal), -n(--note)]
Examples: jou remove "my note", jou remove -n "my note", jou remove -j myJournal`,
	Run: func(cmd *cobra.Command, args []string) {
		if removeJFlag {
			journalRemove(args)
		} else {
			// TODO remove note
		}
	},
}

var useCMD = &cobra.Command{
	Use:   "use",
	Args:  cobra.ExactArgs(1),
	Short: "Use journal as default",
	Long:  "Use journal as default to ommit journal name during jou add, remove. Example: jou use myJournal",
	Run: func(cmd *cobra.Command, args []string) {
		journalSetDefault(args)
	},
}

var listCMD = &cobra.Command{
	Use:   "list",
	Args:  cobra.NoArgs,
	Short: "list journals",
	Long:  "List all journals. Example: jou list",
	Run: func(cmd *cobra.Command, args []string) {
		journalGetAll()
	},
}

// alias jou use to jou journal default set [name]
// alias jou list to jou journal get
// TODO add jou open, which lists all notes and adds
// button for note creation. it also accepts optional journal name

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCMD.AddCommand(beginCMD)
	beginCMD.Flags().BoolVarP(&beginDFlag, "default", "d", false, "use to begin journal as default")
	rootCMD.AddCommand(useCMD)
	rootCMD.AddCommand(listCMD)
	rootCMD.AddCommand(removeCMD)
	removeCMD.Flags().BoolVarP(&removeJFlag, "journal", "j", false, "use to remove journal")

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

jou add [?note]
*/
