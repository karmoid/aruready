package main

// Shelf stores a list of products which are available for purchase
type Shelf struct {
	projects   *ProjectManager
	activities map[string]int
}

// NewShelf instantiates a new Shelf object
func NewShelf() *Shelf {
	return &Shelf{
		projects:   NewProjectManager(),
		activities: make(map[string]int),
	}
}

func (s *Shelf) AddProject(projectCode, projectName string, projectId int) {
	s.projects.Add(projectCode, projectName, projectId)
}

func (s *Shelf) AddActivity(activityName string) {
	s.activities[activityName] = 0
}

func (s *Shelf) FindProject(projectName string) int {
	if item := s.projects.FindByName(projectName); item != nil {
		return item.ID
	}
	return -1
}

func (s *Shelf) FindActivity(activityName string) int {
	return s.activities[activityName]
}

func (s *Shelf) ProjectCount() int {
	return s.projects.Count()
}

func (s *Shelf) ActivityCount() int {
	return len(s.activities)
}
