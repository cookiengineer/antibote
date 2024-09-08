package main

import "antibote/actions"
import "antibote/console"
import "antibote/structs"
import "os"
import os_user "os/user"
import "strconv"
import "strings"
import "time"

func showHelp() {

	console.Info("Antibote")
	console.Info("GitHub Botnet Discovery Tool")

	console.Group("Usage:")
	console.Log("    antibote \"username\"")
	console.Log("    antibote \"username/project\"")
	console.Log("")
	console.GroupEnd("------")

	console.Group("Examples:")
	console.Log("    # scrapes followers to ~/Antibote")
	console.Log("    antibote name-of-doxxer;")
	console.Log("")
	console.Log("    # scrapes followers/stagazers to ~/Antibote")
	console.Log("    antibote name-of-doxxer/name-of-repo;")
	console.Log("")
	console.Log("    # scrapes remaining users and repos to ~/Antibote")
	console.Log("    antibote;")
	console.GroupEnd("---------")

}

func main() {

	home := os.Getenv("HOME")
	user := ""
	repo := ""

	if len(os.Args) == 2 {

		tmp := strings.TrimSpace(os.Args[1])

		if strings.HasPrefix(tmp, "\"") && strings.HasSuffix(tmp, "\"") {
			tmp = tmp[0:len(tmp)-1]
		} else if strings.HasPrefix(tmp, "'") && strings.HasSuffix(tmp, "'") {
			tmp = tmp[0:len(tmp)-1]
		}

		if strings.Contains(tmp, "/") {
			user = strings.TrimSpace(tmp[0:strings.Index(tmp, "/")])
			repo = strings.TrimSpace(tmp[strings.Index(tmp, "/")+1:])
		} else {
			user = strings.TrimSpace(tmp)
		}

	}

	if home == "" {

		user, err := os_user.Current()

		if err == nil {

			if user.Username == "root" {
				home = "/root"
			} else {
				home = "/home/" + user.Username
			}

		}

	}

	if home != "" {

		cache := structs.NewCache(home + "/Antibote")
		tasks := structs.NewTasks()
		cache.Read()

		discover := true

		if user != "" && repo != "" {

			if !cache.IsCompletedTask(user) {
				tasks.AddUser(user, true)
			}

			if !cache.IsCompletedTask(user + "/" + repo) {
				tasks.AddRepo(user, repo, true)
			}

			discover = false

		} else if user != "" {

			if !cache.IsCompletedTask(user) {
				tasks.AddUser(user, true)
			}

			discover = false

		} else {

			remaining := cache.RemainingTasks()

			for r := 0; r < len(remaining); r++ {

				tmp := remaining[r]

				if strings.Contains(tmp, "/") {

					user := tmp[0:strings.Index(tmp, "/")]
					repo := tmp[strings.Index(tmp, "/")+1:]

					tasks.AddRepo(user, repo, true)

				} else {
					tasks.AddUser(tmp, true)
				}

			}

		}

		if tasks.IsDone() {

			console.Info("No Remaining Tasks.")

		} else {

			for !tasks.IsDone() {

				task := tasks.Next()

				if task != nil {

					console.Group("Task " + task.String())

					var err error = nil

					if task.Type == "user" {
						err = actions.ScrapeUser(&cache, &tasks, task, discover)
					} else if task.Type == "repo" {
						err = actions.ScrapeRepository(&cache, &tasks, task, discover)
					}

					console.GroupEnd("")

					remaining := tasks.Remaining()

					if remaining > 0 {
						console.Log("Remaining Tasks: " + strconv.Itoa(remaining))
					}

					if err != nil {
						console.Warn("Rate limited by GitHub, waiting for 5 Minutes ...")
						time.Sleep(5 * time.Minute)
					} else {
						cache.CompleteTask(task)
						tasks.MarkComplete(task)
						cache.Write()
					}

				} else {
					break
				}

			}

		}

	} else {

		showHelp()
		os.Exit(1)

	}

}
