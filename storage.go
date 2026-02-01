package main

import (
  "encoding/json"
  "os"
  "path/filepath"
)

func dbPath() string {
	base, err := os.UserConfigDir()
	if err != nil {
		panic("cannot find config directory")
	}

	dir := filepath.Join(base, "checklist")
	os.MkdirAll(dir, 0755)

	return filepath.Join(dir, "data.json")
}

func loadDatabase() Database {
  path := dbPath()
  db := Database{
    Lists: map[string]*Checklist{},
  }

  data,err := os.ReadFile(path)
  if err != nil {

    if os.IsNotExist(err) {
      return db
    }

    panic("Error reading file")
  }

  json.Unmarshal(data,&db)
  return db
}

func SaveDatabase(db Database) error{
  data,err := json.MarshalIndent(db,""," ")
  if err != nil {
    panic("Error saving file")
  }
  return os.WriteFile(dbPath(),data,0644)
}
