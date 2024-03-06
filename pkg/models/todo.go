package models

const (
	StatusInProgress = "In Progress"
	StatusIncomplete = "Incomplete"
	StatusCancelled  = "Cancelled"
	StatusComplete   = "Complete"
)

type Todo struct {
	ID     int
	Task   string
	Status string
}
