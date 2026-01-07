# dinar 💰

A lightweight Go CLI to track expenses from the terminal: **list**, **add**, **remove**, **update** and **clear** items stored in a local `data.json`, rendered as a clean ASCII table with a total cost footer .

## Features 

- Stores expenses in a simple JSON file (`data.json`) 
- Commands:
  - `list` — display all expenses as an ASCII table + total footer 
  - `add <title> <cost> <quantity>` — add a new expense item 
  - `remove <id>` — remove an expense by ID 
  - `update <id> <title> <cost> <quantity>` — update an expense by ID ✏️
  - `clear` — cLear the list 
- Pretty terminal tables using `go-pretty` 
- Clean structure: Cobra commands in `cmd/`, models + JSON helpers in `pkg/` 

## Project structure 

```bash
go-dinar/
├─ cmd/ # Cobra commands (list/add/remove/update)
├─ pkg/ # Shared code
│ ├─ models.go # Expense structs/models
│ └─ utils.go # JSON read/write helpers (load/save)
├─ data.json # Persistent storage (JSON array of expenses)
├─ main.go # App entrypoint (root command execution)
├─ go.mod
└─ go.sum
```

## Data format 

`data.json` is a JSON array of expense objects:

```json
[
{ "ID": "1", "Title": "buy coffee", "Cost": 100, "Quantity": 3 },
{ "ID": "2", "Title": "netflix subscription", "Cost": 1800, "Quantity": 1 },
{ "ID": "3", "Title": "rent", "Cost": 15000, "Quantity": 1 }
]
```

## Installation 

### Option A: build locally 

```
git clone https://github.com/zineeddinehazi/go-dinar.git
cd go-dinar
go build -o dinar .
```

Run it:

```
./dinar list
```

### Option B: install to GOPATH/bin then run (depending on your PATH):

```
go install ./...
dinar list
```

## Usage 

### List expenses 

```
dinar list
```
```
+-----------------+---------------------+-----------+-------------+
| ID              | ITEM                | COST      |    QUANTITY |
+-----------------+---------------------+-----------+-------------+
| 8ac6d27b        | Buy coffee          | DA 100.0  |           3 |
| ff82fd4a        | Dinner with friends | DA 2000.0 |           1 |
+-----------------+---------------------+-----------+-------------+
| NUM OF ITEMS: 2 | TOTAL               | DA 2300.0 | -----X----- |
+-----------------+---------------------+-----------+-------------+
```

### Add an expense 

```
dinar add "dinner date" 4000 1
```

### Remove an expense 

```
dinar remove ff82fd4a
```

### Update an expense 

```
dinar update ff82fd4a "netflix (family)" 1800 1
```

### Clear the entire list 

```
dinar clear
```

## Roadmap for future improvements 

- [ ] Add `--file` flag to choose a custom JSON path 
- [ ] Add sorting/filtering (by cost, title, etc.) 
- [ ] Export to CSV 
- [ ] Add tests for JSON utils + commands
