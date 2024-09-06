package github

import "antibote/structs"
import "antibote/github/api"
import "encoding/json"
import "strconv"

func GetCommits(user string, repo string) []api.Commit {

	scraper := structs.NewScraper(1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Token": Token,
		"User-Agent": "git-identify (Cookie Engineer's Forensics Tools)",
	}

	commits := make([]api.Commit, 0)
	buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/commits?page=1")
	err := json.Unmarshal(buffer, &commits)

	if err == nil && len(commits) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_commits := make([]api.Commit, 0)

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

