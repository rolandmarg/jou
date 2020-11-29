package cmd

import (
	"fmt"
	"os"

	"github.com/rolandmarg/jou/internal/app/journal"
	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

func journalGet(args []string) {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	j, err := r.Get(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if j == nil {
		s := fmt.Sprintf(`Journal "%v" not found`, args[0])
		fmt.Println(s)
		os.Exit(0)
	}

	fmt.Println(j)
}

func journalCreate(args []string, isDefault bool) {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	_, err = r.Create(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s := fmt.Sprintf(`Journal "%v" created`, args[0])
	fmt.Println(s)

	if isDefault {
		err = r.SetDefault(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func journalRemove(args []string) {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	err = r.Remove(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s := fmt.Sprintf(`Journal "%v" created`, args[0])
	fmt.Println(s)
}

func journalGetDefault() {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	j, err := r.GetDefault()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if j == nil {
		s := fmt.Sprintf(`Default journal not found`)
		fmt.Println(s)
		os.Exit(0)
	}

	fmt.Println(j)
}

func journalGetAll() {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	j, err := r.GetAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if j == nil || len(j) == 0 {
		s := fmt.Sprintf(`No journals found`)
		fmt.Println(s)
		os.Exit(0)
	}

	for _, jo := range j {
		fmt.Println(jo)
	}
}

func journalSetDefault(args []string) {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	err = r.SetDefault(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s := fmt.Sprintf(`Journal "%v" set as default`, args[0])
	fmt.Println(s)
}

func journalUpdate(args []string) {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.Close()

	r := journal.MakeRepository(db)

	err = r.Update(args[0], args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	s := fmt.Sprintf(`Journal "%v" renamed to "%v"`, args[0], args[1])
	fmt.Println(s)
}
