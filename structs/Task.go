package structs

import "time"

type Task struct {
	Type      string `json:"type"`
	User      string `json:"user"`
	Repo      string `json:"repo"`
	Discover  bool   `json:"discover"`
	Completed string `json:"completed"`
}

func (task *Task) Complete() {
	task.Completed = time.Now().Format(time.RFC3339)
}

func (task *Task) IsComplete() bool {
	return task.Completed != ""
}

func (task *Task) String() string {

	var result string

	if task.Type == "user" {
		result = "@" + task.User
	} else if task.Type == "repo" {
		result = "@" + task.User + "/" + task.Repo
	}

	return result

}
