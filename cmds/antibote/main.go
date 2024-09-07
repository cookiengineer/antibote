package main

import "antibote/actions"
import "antibote/structs"
import "fmt"
import "os"
import os_user "os/user"
import "strconv"
import "strings"
import "time"

func showHelp() {

	fmt.Println("Usage:")
	fmt.Println("    antibote \"username\"")
	fmt.Println("    antibote \"username/project\"")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("    # scrapes botnet accounts to ~/Antibote")
	fmt.Println("    antibote name-of-doxxer;")

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

	if home != "" && user != "" {

		cache := structs.NewCache(home + "/Antibote")
		cache.Read()

		if user != "" && repo != "" {

			if !cache.IsCompletedTask(user) {
				cache.AddTask(user)
			}

			if !cache.IsCompletedTask(user + "/" + repo) {
				cache.AddTask(user + "/" + repo)
			}

		} else if user != "" {

			if !cache.IsCompletedTask(user) {
				cache.AddTask(user)
			}

		}

		user_tasks := cache.GetUserTasks()

		for len(user_tasks) > 0 {

			for u := 0; u < len(user_tasks); u++ {

				actions.ScrapeUser(&cache, user_tasks[u])
				fmt.Println("Remaining Users: " + strconv.Itoa(len(user_tasks)))
				time.Sleep(60 * time.Second)

			}

			user_tasks = cache.GetUserTasks()

		}

		repo_tasks := cache.GetRepoTasks()

		for len(repo_tasks) > 0 {

			for r := 0; r < len(repo_tasks); r++ {

				user := repo_tasks[r][0:strings.Index(repo_tasks[r], "/")]
				repo := repo_tasks[r][strings.Index(repo_tasks[r], "/")+1:]

				actions.ScrapeRepository(&cache, user, repo)

				fmt.Println("Remaining Repositories: " + strconv.Itoa(len(repo_tasks)))
				time.Sleep(60 * time.Second)

			}

		}

	} else {

		showHelp()
		os.Exit(1)

	}

}
