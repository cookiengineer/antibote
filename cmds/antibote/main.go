package main

import "antibote/actions"
import "antibote/structs"
import "fmt"
import "os"
import os_user "os/user"
import "strings"

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
			home = "/home/" + user.Username
		}

	}

	if home != "" {

		cache := structs.NewCache(home + "/Antibote")

		actions.Scrape(&cache, user)

	}

}
