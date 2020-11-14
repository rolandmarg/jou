package main

import (
	"time"
)

// JournalEntry contains entry user input and other metadata
type JournalEntry struct {
	id        int
	deletedAt time.Time
	createdAt time.Time
	*JournalEntryInput
}

// JournalEntryInput contains all fields user can provide
type JournalEntryInput struct {
	title string
	body  string
	tags  []string
	mood  string
}

// NewJournalEntry creates a new journal entry
func NewJournalEntry(input *JournalEntryInput) *JournalEntry {
	entry := &JournalEntry{id: 1, createdAt: time.Now(), JournalEntryInput: input}

	return entry
}

// SoftDeleteJournalEntry marks a journal entry as deleted
func (journalEntry *JournalEntry) SoftDeleteJournalEntry() *JournalEntry {
	journalEntry.deletedAt = time.Now()

	return journalEntry
}

// IsJournalEntryDeleted returns boolean whether entry is softdeleted
func (journalEntry *JournalEntry) IsJournalEntryDeleted() bool {
	return !journalEntry.deletedAt.IsZero()
}
