package main

import "antibote/actions"
import "antibote/structs"
import "fmt"
import "os"
import os_user "os/user"
import "strings"
import "time"

func showHelp() {

	fmt.Println("antibote <github-username>")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("    # creates recursive cache in ~/Antibote/github;")
	fmt.Println("    antibote name-of-doxxer;")

}

func main() {

	home := os.Getenv("HOME")
	user := ""

	if len(os.Args) == 2 {
		user = strings.TrimSpace(os.Args[1])
	}

	if home == "" && user != "" {

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
		cache.Read()

		if !cache.IsCompletedTask(user) {
			cache.AddTask(user)
		}

		users := cache.GetTasks()

		for u := 0; u < len(users); u++ {

			actions.Scrape(&cache, users[u])
			time.Sleep(60 * time.Second)

		}


	}

}
