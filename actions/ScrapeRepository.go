package actions

import "antibote/console"
import "antibote/github"
import "antibote/structs"

func ScrapeRepository(cache *structs.Cache, tasks *structs.Tasks, task *structs.Task, discover_recursive bool) error {

	var err error = nil

	user := cache.GetUser(task.User)
	repository := user.GetRepository(task.Repo)

	if repository != nil && repository.IsFork == false {

		if task.Discover == true {

			stargazers, err1 := github.GetStargazers(cache, user.Name, repository.Name)

			if err1 == nil {

				for s := 0; s < len(stargazers); s++ {
					tasks.AddUser(stargazers[s].Name, discover_recursive)
				}

			} else {
				console.Error("GetStargazers() error: " + err1.Error())
				err = err1
			}

		}

		commits, err2 := github.GetCommits(cache, user.Name, repository.Name)

		if err2 == nil {

			for c := 0; c < len(commits); c++ {
				repository.AddCommit(commits[c])
			}

		} else {
			console.Error("GetCommits() error: " + err2.Error())
			err = err2
		}

		user.TrackRepository(repository)
		cache.Write()

	}

	return err

}
