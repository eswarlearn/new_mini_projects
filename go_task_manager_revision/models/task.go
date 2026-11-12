
package models

type Task struct {
    ID          uint32
    Title       string
    Description string
    Status      string
    Priority    string
}

const (
    StatusPending    = "pending"
    StatusInProgress = "in_progress"
    StatusCompleted  = "completed"
)
