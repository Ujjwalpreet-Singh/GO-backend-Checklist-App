package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Println(`
Usage:
  checklist create <list>
  checklist add <list> <text>
  checklist done <list> <id>
  checklist delete <list> <id>
  checklist list <list>
`)
}

func main() {

  startServer()
	if len(os.Args) < 2 {
		usage()
		return
	}

	cmd := os.Args[1]
	db := loadDatabase()

	var err error

	switch cmd {

	case "create":
		if len(os.Args) < 3 {
			fmt.Println("missing list name")
			return
		}
		err = CreateChecklist(&db, os.Args[2])

	case "add":
		if len(os.Args) < 4 {
			fmt.Println("missing item text")
			return
		}
		list := os.Args[2]
		text := strings.Join(os.Args[3:], " ")
		err = addItem(&db, list, text)

	case "done":
		if len(os.Args) < 4 {
			fmt.Println("missing id")
			return
		}
		id, _ := strconv.Atoi(os.Args[3])
		err = MarkDone(&db, os.Args[2], id)

	case "delete":
		if len(os.Args) < 4 {
			fmt.Println("missing id")
			return
		}
		id, _ := strconv.Atoi(os.Args[3])
		err = deleteItem(&db, os.Args[2], id)

	case "list":
		if len(os.Args) < 3 {
			fmt.Println("missing list name")
			return
		}
		items, e := listItems(&db, os.Args[2])
		if e != nil {
			fmt.Println("error:", e)
			return
		}

		for _, it := range items {
			box := "[ ]"
			if it.Done {
				box = "[x]"
			}
			fmt.Printf("%d %s %s\n", it.ID, box, it.Text)
		}

	default:
		usage()
		return
	}

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	SaveDatabase(db)
}
