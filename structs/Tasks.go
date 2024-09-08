package structs

type Tasks struct {
	users map[string]*Task
	repos map[string]*Task
}

func NewTasks() Tasks {

	var tasks Tasks

	tasks.users = make(map[string]*Task)
	tasks.repos = make(map[string]*Task)

	return tasks

}

func (tasks *Tasks) AddUser(user string, discover bool) {

	_, ok := tasks.users[user]

	if ok == false {

		task := Task{
			Type:     "user",
			User:     user,
			Repo:     "",
			Discover: discover,
			Completed: "",
		}

		tasks.users[user] = &task

	}

}

func (tasks *Tasks) AddRepo(user string, repo string, discover bool) {

	_, ok := tasks.repos[user + "/" + repo]

	if ok == false {

		task := Task{
			Type:      "repo",
			User:      user,
			Repo:      repo,
			Discover:  discover,
			Completed: "",
		}

		tasks.repos[user] = &task

	}

}

func (tasks *Tasks) IsDone() bool {

	var result bool = true

	for _, task := range tasks.repos {

		if task.IsComplete() == false {
			result = false
			break
		}

	}

	for _, task := range tasks.users {

		if task.IsComplete() == false {
			result = false
			break
		}

	}

	return result

}

func (tasks *Tasks) IsCompletedUser(user string) bool {

	var result bool = false

	task, ok := tasks.users[user]

	if ok == true {
		result = task.IsComplete()
	}

	return result

}

func (tasks *Tasks) IsCompletedRepo(user string, repo string) bool {

	var result bool = false

	task, ok := tasks.repos[user + "/" + repo]

	if ok == true {
		result = task.IsComplete()
	}

	return result

}

func (tasks *Tasks) Next() *Task {

	var result *Task = nil

	for name, task := range tasks.users {

		if task.IsComplete() == false {
			result = tasks.users[name]
			break
		}

	}

	for name, task := range tasks.repos {

		if task.IsComplete() == false {
			result = tasks.repos[name]
			break
		}

	}

	return result

}

func (tasks *Tasks) Remaining() int {

	var result int = 0

	for _, task := range tasks.repos {

		if task.IsComplete() == false {
			result++
		}

	}

	for _, task := range tasks.users {

		if task.IsComplete() == false {
			result++
		}

	}

	return result

}

func (tasks *Tasks) MarkComplete(task *Task) {

	if task.Type == "user" {
		task.Complete()
		tasks.users[task.User] = task
	} else if task.Type == "repo" {
		task.Complete()
		tasks.users[task.User + "/" + task.Repo] = task
	}

}

