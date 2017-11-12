package main

import (
	"fmt"
	"strings"
)

type ProjectManager struct {
	projects map[int]*Item
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{projects: make(map[int]*Item)}
}

// Add put a new Item in projects array
func (pm ProjectManager) Find(id int) *Item {
	return pm.projects[id]
}

// Count returns Map size
func (pm ProjectManager) Count() int {
	return len(pm.projects)
}

// Add put a new Item in projects array
func (pm ProjectManager) FindByName(name string) *Item {
	for _, project := range pm.projects {
		if strings.Compare(project.Name, name) == 0 {
			return project
		}
	}
	return nil
}

// Add put a new Item in projects array
func (pm ProjectManager) Add(code, name string, id int) (*Item, error) {
	project := pm.Find(id)
	if project != nil {
		return project, fmt.Errorf("DUPLICATE: Le projet %s portant l'id %d existe deja", project.Code, id)
	}
	item := &Item{Code: code, Name: name, ID: id}
	pm.projects[id] = item
	return item, nil
}
