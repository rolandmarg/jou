## jou

jou is a CLI app for writing down notes.  
Built with golang, sqlite, spf13/cobra for CLI interface

## Commands

- jou add

Add a note. Ommit journal name for default journal.  
Examples: `jou add "goes to default journal"`, `jou add secretj "deep note"`

```
  -h, --help   help for add
```

- jou begin

Begin a new journal. Specify -d(--default) to make it default.  
Examples: `jou begin books`, `jou begin -d mix`

```
  -d, --default   use to begin journal as default
  -h, --help      help for begin
```

- jou fang

Delete a journal. Default journal can't be deleted.  
Example: `jou fang Didle`

```
  -h, --help   help for fang
```

- jou list

List all journals.  
Example: `jou list`

```
  -h, --help   help for list
```

- jou use

Use journal as default. Useful for ommiting journal name on certain commands  
Example: `jou use bubu`

```
  -h, --help   help for use
```

- jou open

Show journal entries. Ommit name to open default journal.  
Examples: `jou open`, `jou open bubu`

```
  -h, --help   help for open
```
