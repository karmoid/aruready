Feature: Timesheet basket
  In order to fill my timesheets
  As a user
  I need to be able to put timesheet item into a basket

  Rules:
  - One day timesheet is full (~8hours)
  - No hour micromanagement
  - Using Small, Medium, Half, Full work period

  Background:
  Given there is a "Known" project with "known project" as name and 100 as ID
  And there is a "Meeting" activity
  And the following projects exist:
    | code  | name     | id |
    | Prj2  | projet 1 | 1  |
    | Prj3  | projet 2 | 2  |
    | Prj4  | projet 3 | 3  |
  And the following activities exist:
    | code  | name    | id  |
    | meet  | meeting | 123 |
    | study | study   | 124 |
    | dev   | develop | 125 |

  Scenario: Getting list size give good result
    When I ask for projets list size I get 4
    And activities list size is 4

  Scenario: Check for present/missing values
    When I'm looking for "Known" project it exists
    And "Unknown" project doesn't
    And "meet" activity exists
    But "code" activity doesn't

  Scenario: Adding a work day with 1 large activity
    When I add the "Known/Meeting" with "Large" activity into the basket
    Then I should have 1 timesheet in the basket
    And the overall basket weight should be .25

  Scenario: Adding a full work day with 4 activities
    When I add the "Known/Meeting" with "Large" activity into the basket
    And I add the "Known/Meeting" with "Medium" activity into the basket
    And I add the "Known/Meeting" with "Medium" activity into the basket
    And I add the "Known/Meeting" with "Half" activity into the basket
    Then I should have 4 timesheet in the basket
    And the overall basket weight should be 1.0

  Scenario: Adding a extra full work day with 2 activities
    When I add the "Known/Meeting" with "Full" activity into the basket
    And I add the "Known/Meeting" with "Medium" activity into the basket
    Then I should have 2 timesheet in the basket
    And the overall basket weight should be 1.0
