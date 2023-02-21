package models

var PathTasks = "/tasks/"

type TaskBody struct {
	Id     int    `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	IsDone bool   `json:"is_done"`
}
