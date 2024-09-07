package actions

import "antibote/github"
import "antibote/structs"
import "antibote/types"
import "fmt"

func ScrapeUser(cache *structs.Cache, name string) {

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

			if !cache.IsCompletedTask(user.Name) {

				followers, err1 := github.GetFollowers(cache, user.Name)

				if err1 == nil {

					for f := 0; f < len(followers); f++ {
						cache.AddTask(followers[f].Name)
					}

					cache.Write()

				}

				repositories, err2 := github.GetRepositories(cache, user.Name)

				if err2 == nil {

					for r := 0; r < len(repositories); r++ {
						cache.AddTask(user.Name + "/" + repositories[r].Name)
					}

					cache.Write()

				}

				if err1 == nil && err2 == nil {
					cache.CompleteTask(user.Name)
				}

			}

			keys := user.ToKeys()

			for k := 0; k < len(keys); k++ {
				cache.AddKey(keys[k].ID, keys[k].Email)
			}

			cache.Write()

		}

	}

}
