package main

type Item struct {
	Code string
	Name string
	ID   int
}

type ItemManager interface {
	Add(code, name string, id int) (*Item, error)
	Find(id int) *Item
	Find_by_name(name string) *Item
	Count() int
	// Remove(id int) error
}
