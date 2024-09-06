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
	fmt.Println("    # scrapes botnet accounts to ~/Antibote")
	fmt.Println("    antibote name-of-doxxer;")

}

func main() {

	home := os.Getenv("HOME")
	user := ""

	if len(os.Args) == 2 {
		user = strings.TrimSpace(os.Args[1])
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

		if !cache.IsCompletedTask(user) {
			cache.AddTask(user)
		}

		users := cache.GetTasks()

		for u := 0; u < len(users); u++ {

			actions.Scrape(&cache, users[u])
			time.Sleep(60 * time.Second)

		}

	} else {

		showHelp()
		os.Exit(1)

	}

}
