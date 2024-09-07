package structs

import "strings"
import "time"

func (cache *Cache) AddTask(value string) {
	cache.TaskMap[value] = ""
}

func (cache *Cache) GetUserTasks() []string {

	users := make(map[string]bool)

	for task, time := range cache.TaskMap {

		if time == "" {

			if !strings.Contains(task, "/") {
				users[task] = true
			}

		}

	}

	var result []string

	for user, _ := range users {
		result = append(result, user)
	}

	return result

}

func (cache *Cache) GetRepoTasks() []string {

	users := make(map[string]bool)

	for task, time := range cache.TaskMap {

		if time == "" {

			if strings.Contains(task, "/") {
				users[task] = true
			}

		}

	}

	var result []string

	for user, _ := range users {
		result = append(result, user)
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

func (cache *Cache) CompleteTask(value string) {
	cache.TaskMap[value] = time.Now().Format(time.RFC3339)
}

