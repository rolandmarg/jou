package main

import (
	"fmt"
)

func main() {
	journal := NewJournal("journal")

	journal.AddEntry(&JournalEntryInput{
		title: "my first entry",
		body:  "I'm feeling good that I write Go",
	})

	fmt.Print(journal)
}

/*
TODO add sqlite
TODO add tests
TODO add pincode on journal
TODO add geolocation
*/