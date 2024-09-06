package github

import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "strconv"

func GetCommits(cache *structs.Cache, user string, repo string) []types.Commit {

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	commits := make([]types.Commit, 0)
	buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/commits?page=1")
	err := json.Unmarshal(buffer, &commits)

	if err == nil && len(commits) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_commits := make([]types.Commit, 0)
			page_buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/commits?page=" + page)

			if len(page_buffer) > 0 {

				err := json.Unmarshal(page_buffer, &page_commits)

				if err == nil {

					for pc := 0; pc < len(page_commits); pc++ {
						commits = append(commits, page_commits[pc])
					}

					if len(page_commits) < 30 {
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

	return commits

}

