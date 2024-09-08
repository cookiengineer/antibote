package actions

import "antibote/console"
import "antibote/github"
import "antibote/structs"

func ScrapeUser(cache *structs.Cache, tasks *structs.Tasks, task *structs.Task, discover_recursive bool) error {

	var err error = nil

	user := cache.GetUser(task.User)

	if user.Name != "" {

		if task.Discover == true {

			followers, err1 := github.GetFollowers(cache, user.Name)

			if err1 == nil {

				for f := 0; f < len(followers); f++ {
					tasks.AddUser(followers[f].Name, discover_recursive)
				}

				cache.Write()

			} else {
				console.Error("GetFollowers() error: " + err1.Error())
				err = err1
			}

		}

		repositories, err2 := github.GetRepositories(cache, user.Name)

		if err2 == nil {

			for r := 0; r < len(repositories); r++ {

				repository := repositories[r]

				user.TrackRepository(&repository)

				tasks.AddRepo(user.Name, repositories[r].Name, discover_recursive)

			}

			cache.Write()

		} else {
			console.Error("GetRepositories() error: " + err2.Error())
			err = err2
		}

		keys := user.ToKeys()

		for k := 0; k < len(keys); k++ {
			cache.AddKey(keys[k].ID, keys[k].Email)
		}

		cache.Write()

	}

	return err

}
