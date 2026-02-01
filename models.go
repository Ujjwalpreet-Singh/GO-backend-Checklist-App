package main

type Item struct {
  ID int
  Text string
  Done bool
}

type Checklist struct {
  Name string
  Items []Item
}

type Database struct {
  Lists map[string]*Checklist
}
