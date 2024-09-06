package github

import "antibote/constants"
import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "strconv"

func GetRepositories(cache *structs.Cache, user string) []types.Repository {

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": constants.Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	repositories := make([]types.Repository, 0)
	buffer := scraper.Request("https://api.github.com/users/" + user + "/repos?page=1")
	err := json.Unmarshal(buffer, &repositories)

	if err == nil && len(repositories) == 30 {

		for p := 2; p <= 10; p++ {

			page := strconv.Itoa(p)
			page_repositories := make([]types.Repository, 0)
			page_buffer := scraper.Request("https://api.github.com/users/" + user + "/repos?page=" + page)

			if len(page_buffer) > 0 {

				err := json.Unmarshal(page_buffer, &page_repositories)

				if err == nil {

					for pr := 0; pr < len(page_repositories); pr++ {
						repositories = append(repositories, page_repositories[pr])
					}

					if len(page_repositories) < 30 {
						break
					}

				} else {
					break
				}

			} else {
				break
			}

		}

	}

	return repositories

}
