package models

const (
	StatusInProgress = "In Progress"
	StatusIncomplete = "Incomplete"
	StatusCancelled  = "Cancelled"
	StatusComplete   = "Complete"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
