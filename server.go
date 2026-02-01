package main

import (
	"fmt"
	"net/http"
  "strconv"
  "strings"
)

func startServer() {
  http.HandleFunc("/", handleHome)
  http.HandleFunc("/list/", handleList)
  http.HandleFunc("/create", handleCreate)
  http.HandleFunc("/add", handleAdd)
  http.HandleFunc("/done", handleDone)
  http.HandleFunc("/delete", handleDelete)

  http.Handle("/static/",
  	http.StripPrefix("/static/",
  		http.FileServer(http.Dir("static"))))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	db := loadDatabase()

	html := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Checklists</title>
  <link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
  <div class="card">
	<h1>All Checklists</h1>
	<ul>
	`

	for name := range db.Lists {
		html += fmt.Sprintf(
			"<li><a href='/list/%s'>%s</a></li>",
			name, name,
		)
	}

	html += `
	</ul>

	<form action="/create" method="POST">
	<input name="name" placeholder="New list name">
	<button>Create</button>
	</form>
  </div>
	</body>
	</html>
	`

	w.Write([]byte(html))
}

func handleList(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/list/")

	db := loadDatabase()
	list, ok := db.Lists[name]
	if !ok {
		w.Write([]byte("<h1>List not found</h1>"))
		return
	}

	html := `
	<!DOCTYPE html>
	<html>
	<head>
  <link rel="stylesheet" href="/static/styles.css">
	</head>
	<body>
  <div class="card">
	<h1>` + name + `</h1>
	<ul>
	`

	for _, it := range list.Items {
		class := ""
		if it.Done {
			class = "class='done'"
		}

		html += fmt.Sprintf(
			"<li %s><a href='/done?list=%s&id=%d'>%s</a> <a class='delete' href='/delete?list=%s&id=%d'>[delete]</a></li>",
			class,
			name, it.ID, it.Text,
			name, it.ID,
		)
	}

	html += fmt.Sprintf(`
	</ul>

	<form action="/add" method="POST">
	<input type="hidden" name="list" value="%s">
	<input name="text" placeholder="New task">
	<button>Add</button>
	</form>

	<p><a class="back" href="/">← back</a></p>
  </div>
	</body>
	</html>
	`, name)

	w.Write([]byte(html))
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	text := r.FormValue("text")

	db := loadDatabase()
  list := r.FormValue("list")

  addItem(&db, list, text)
  SaveDatabase(db)

  http.Redirect(w, r, "/list/"+list, http.StatusSeeOther)

}

func handleDone(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db := loadDatabase()
  list := r.URL.Query().Get("list")

  MarkDone(&db, list, id)
  SaveDatabase(db)

  http.Redirect(w, r, "/list/"+list, http.StatusSeeOther)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db := loadDatabase()
  list := r.URL.Query().Get("list")

  deleteItem(&db, list, id)
  SaveDatabase(db)

  http.Redirect(w, r, "/list/"+list, http.StatusSeeOther)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	db := loadDatabase()
	CreateChecklist(&db, name)
	SaveDatabase(db)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
