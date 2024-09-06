package actions

import "antibote/github"
import "antibote/structs"
import "fmt"

var alreadyScraped = map[string]bool

func init() {
	alreadyScraped = make(map[string]bool)
}

func Scrape(cache *structs.Cache, name string) {

	_, ok := alreadyScraped[name]

	if ok == false {

		fmt.Println("Scrape user \"" + name + "\"")

		user = cache.GetUser(name)

		if user == nil {
			user = github.GetUser(name)
		}

		if user.Name != "" {

			repositories := github.GetRepositories(user.Name)

			for r := 0; r < len(repositories); r++ {

				repository := repositories[r]

				if repository.IsFork == false && user.HasRepository(repository.Name) == false {

					repository.Commits = github.GetCommits(user.Name, repository.Name)
					user.AddRepository(repository)

				}

			}

			alreadyScraped[user.Name] = true
			cache.AddUser(user)
			cache.Write()

			followers := github.GetFollowers(user.Name)

			for f := 0; f < len(followers); f++ {
				Scrape(followers[f].Name)
			}

		}

	}

}
