package structs

import "time"

func (cache *Cache) AddTask(value string) {
	cache.TaskMap[value] = ""
}

func (cache *Cache) RemainingTasks() []string {

	var result []string

	for name, time := range cache.TaskMap {

		if time == "" {
			result = append(result, name)
		}

	}

	return result

}

func (cache *Cache) IsCompletedTask(value string) bool {

	var result bool = false

	time, ok := cache.TaskMap[value]

	if ok == true && time != "" {
		result = true
	}

	return result

}

func (cache *Cache) CompleteTask(task *Task) {

	if task.Type == "user" {
		cache.TaskMap[task.User] = time.Now().Format(time.RFC3339)
	} else if task.Type == "repo" {
		cache.TaskMap[task.User + "/" + task.Repo] = time.Now().Format(time.RFC3339)
	}

}

