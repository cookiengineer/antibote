package actions

import "antibote/github"
import "antibote/structs"
import "antibote/types"
import "fmt"

func ScrapeRepository(cache *structs.Cache, username string, reponame string) {

	if !cache.IsCompletedTask(username + "/" + reponame) {

		fmt.Println("Scrape repo \"" + username + "/" + reponame + "\"")

		user := cache.GetUser(username)

		if user == nil {
			tmp := types.NewUser(username)
			user = &tmp
		}

		if user.Name != "" {

			repository := user.GetRepository(reponame)

			if repository != nil && repository.IsFork == false {

				stargazers, err1 := github.GetStargazers(cache, user.Name, repository.Name)

				if err1 == nil {

					for s := 0; s < len(stargazers); s++ {
						cache.AddTask(stargazers[s].Name)
					}

				}

				commits, err2 := github.GetCommits(cache, user.Name, repository.Name)

				if err2 == nil {

					for c := 0; c < len(commits); c++ {
						repository.AddCommit(commits[c])
					}

				}

				user.TrackRepository(repository)
				cache.Write()

				if err1 == nil && err2 == nil {
					cache.CompleteTask(user.Name + "/" + repository.Name)
				}

			}

			cache.TrackUser(user)
			cache.Write()

		}

	}

}
