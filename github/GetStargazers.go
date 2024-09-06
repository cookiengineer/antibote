package github

import "antibote/constants"
import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "strconv"

func GetStargazers(cache *structs.Cache, user string, repo string) []types.User {

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": constants.Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	users := make([]types.User, 0)
	buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/stargazers?page=1")
	err := json.Unmarshal(buffer, &users)

	if err == nil && len(users) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_users := make([]types.User, 0)
			page_buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/stargazers?page=" + page)

			if len(page_buffer) > 0 {

				err := json.Unmarshal(page_buffer, &page_users)

				if err == nil {

					for pu := 0; pu < len(page_users); pu++ {
						users = append(users, page_users[pu])
					}

					if len(page_users) < 30 {
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

	return users

}

