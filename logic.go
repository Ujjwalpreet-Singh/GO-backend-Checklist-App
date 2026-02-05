package main

import "errors"

func CreateChecklist(db *Database, name string) error {
  if _,exists := db.Lists[name]; exists{
    return errors.New("checklist already exists")
  }
  db.Lists[name] = &Checklist{
    Name : name,
    Items : []Item{},
  }

  return nil
}

func DeleteChecklist(db *Database,name string) error {
  if _,exists := db.Lists[name]; !exists{
    return errors.New("No checklist like this exists")
  }
  delete(db.Lists,name)
  return nil
}

func addItem(db *Database, listName, text string) error {
  list,ok := db.Lists[listName]
  if !ok{
    return errors.New("checklist not found")
  }

  id := len(list.Items) + 1

  item := Item{
    ID : id,
    Text: text,
    Done: false,
  }

  list.Items = append(list.Items,item)
  return nil
}

func MarkDone(db *Database,listName string,id int) error {
  list,ok := db.Lists[listName]
  if !ok{
    return errors.New("checklist not found")
  }

  for i := range list.Items {
    if list.Items[i].ID == id {
      list.Items[i].Done = !list.Items[i].Done
      return nil
    }
  }

  return errors.New("Item not found")
}

func deleteItem(db *Database, listName string,id int) error {
  list,ok := db.Lists[listName]
  if !ok{
    return errors.New("checklist not found")
  }

  newItems := []Item{}

  for _,item := range list.Items {
    if item.ID != id {
      newItems = append(newItems,item)
    }
  }
  list.Items = newItems
  return nil
}

func listItems(db *Database, listName string) ([]Item, error) {
	list, ok := db.Lists[listName]
	if !ok {
		return nil, errors.New("checklist not found")
	}

	return list.Items, nil
}
