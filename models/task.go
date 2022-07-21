package models

var PathTasks = "/tasks/"

type TaskBody struct {
	Id   int    `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}
