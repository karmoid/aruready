package main

import (
	"fmt"
	"strconv"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

type sheeting struct {
	shelf  *Shelf
	basket *Basket
}

func (sh *sheeting) AddToBasket(activityName, weight string) error {
	sh.basket.AddItem(activityName, weight)
	return nil
}

func (sh *sheeting) CheckBasketTimesheetCount(timesheetsCount int) error {
	if sh.basket.GetBasketSize() != timesheetsCount {
		return fmt.Errorf(
			"expected %d timesheets but there are %d",
			sh.basket.GetBasketSize(),
			timesheetsCount,
		)
	}
	return nil
}

func (sh *sheeting) theOverallBasketWeightShouldBe(basketWeight float64) error {
	if sh.basket.GetBasketTotal() != basketWeight {
		return fmt.Errorf(
			"expected basket total to be %.2f but it is %.2f",
			basketWeight,
			sh.basket.GetBasketTotal(),
		)
	}
	return nil
}

func (sh *sheeting) addProject(projectCode, projectName string, projectId int) (err error) {
	sh.shelf.AddProject(projectCode, projectName, projectId)
	return
}

func (sh *sheeting) addActivity(activityName string) (err error) {
	sh.shelf.AddActivity(activityName)
	return
}

func (sh *sheeting) addActivities(activities *gherkin.DataTable) error {
	for idx, activity := range activities.Rows {
		if idx > 0 {
			sh.shelf.AddActivity(activity.Cells[0].Value)
		}
	}
	return nil
}

func (sh *sheeting) addProjects(projects *gherkin.DataTable) error {
	for idx, project := range projects.Rows {
		if idx > 0 {
			id, err := strconv.Atoi(project.Cells[2].Value)
			if err != nil {
				return err
			}
			sh.shelf.AddProject(project.Cells[0].Value, project.Cells[1].Value, id)
		}
	}
	return nil
}

func (sh *sheeting) CheckProjetsListSize(projectCount int) error {
	if sh.shelf.ProjectCount() != projectCount {
		return fmt.Errorf(
			"expected projects count to be %.2f but it is %.2f",
			projectCount,
			sh.shelf.ProjectCount(),
		)
	}
	return nil
}

func (sh *sheeting) CheckActivitiesListSize(activityCount int) error {
	if sh.shelf.ActivityCount() != activityCount {
		return fmt.Errorf(
			"expected activities count to be %.2f but it is %.2f",
			activityCount,
			sh.shelf.ActivityCount(),
		)
	}
	return nil
}

func (sh *sheeting) ProjectExists(projectName string) error {
	if sh.shelf.FindProject(projectName) < 0 {
		return fmt.Errorf("Project %s must exists but it doesn't", projectName)
	}
	return nil
}

func (sh *sheeting) ProjectNotExists(projectName string) error {
	if sh.shelf.FindProject(projectName) >= 0 {
		return fmt.Errorf("Project %s mustn't exists but it does", projectName)
	}
	return nil
}

func (sh *sheeting) ActivityExists(activityName string) error {
	if sh.shelf.FindActivity(activityName) < 0 {
		return fmt.Errorf("Activity %s must exists but it doesn't", activityName)
	}
	return nil
}

func (sh *sheeting) ActivityNotExists(activityName string) error {
	if sh.shelf.FindActivity(activityName) >= 0 {
		return fmt.Errorf("Activity %s mustn't exists but it does", activityName)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {

	sh := &sheeting{}

	s.BeforeScenario(func(interface{}) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()
	})

	s.Step(`^there is a "([^"]*)" project with "([^"]*)" as name and (\d+) as ID$`, sh.addProject)
	s.Step(`^there is a "([^"]*)" activity$`, sh.addActivity)
	s.Step(`^I add the "([^"]*)" with "([^"]*)" activity into the basket$`, sh.AddToBasket)
	s.Step(`^I should have (\d+) timesheet in the basket$`, sh.CheckBasketTimesheetCount)
	// s.Step(`^the overall basket weight should be (\d+)\.(\d+)$`, sh.theOverallBasketWeightShouldBe)
	s.Step(`^the overall basket weight should be ([-+]?[0-9]*\.?[0-9]+)$`, sh.theOverallBasketWeightShouldBe)
	s.Step(`^the following activities exist:$`, sh.addActivities)
	s.Step(`^I ask for projets list size I get (\d+)$`, sh.CheckProjetsListSize)
	s.Step(`^activities list size is (\d+)$`, sh.CheckActivitiesListSize)
	s.Step(`^I\'m looking for "([^"]*)" project it exists$`, sh.ProjectExists)
	s.Step(`^"([^"]*)" project doesn\'t$`, sh.ProjectNotExists)
	s.Step(`^"([^"]*)" activity exists$`, sh.ActivityExists)
	s.Step(`^"([^"]*)" activity doesn\'t$`, sh.ActivityNotExists)
	s.Step(`^the following projects exist:$`, sh.addProjects)
}
