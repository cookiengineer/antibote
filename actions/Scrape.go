package actions

import "antibote/github"
import "antibote/structs"
import "antibote/types"
import "fmt"

func Scrape(cache *structs.Cache, name string) {

	if !cache.IsCompletedTask(name) {

		fmt.Println("Scrape user \"" + name + "\"")

		user := cache.GetUser(name)

		if user == nil {
			tmp := types.NewUser(name)
			user = &tmp
		}

		if user.Name != "" {

			cache.TrackUser(user)
			cache.Write()

			if !cache.IsCompletedTask(user.Name + ":repositories") {

				repositories := github.GetRepositories(cache, user.Name)

				if len(repositories) > 0 {

					for r := 0; r < len(repositories); r++ {

						repository := repositories[r]

						if repository.IsFork == false && user.HasRepository(repository.Name) == false {

							repository.Commits = github.GetCommits(cache, user.Name, repository.Name)
							user.TrackRepository(&repository)

						}

					}

					cache.CompleteTask(user.Name + ":repositories")
					cache.Write()

				}

			}

			if !cache.IsCompletedTask(user.Name + ":followers") {

				followers := github.GetFollowers(cache, user.Name)

				if len(followers) > 0 {

					for f := 0; f < len(followers); f++ {
						cache.AddTask(followers[f].Name)
					}

					cache.CompleteTask(user.Name + ":followers")
					cache.Write()

				}

			}

			keys := user.ToKeys()

			for k := 0; k < len(keys); k++ {
				cache.AddKey(keys[k].ID, keys[k].Email)
			}

			cache.Write()

			if len(user.Repositories) > 0 {
				cache.CompleteTask(user.Name)
			}

		}

	}

}
