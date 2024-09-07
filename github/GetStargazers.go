package github

import "antibote/constants"
import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "errors"
import "strconv"

func GetStargazers(cache *structs.Cache, user string, repo string) ([]types.User, error) {

	var err error = nil

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": constants.Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	users := make([]types.User, 0)
	buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/stargazers?page=1")
	err1 := json.Unmarshal(buffer, &users)

	if err1 == nil && len(users) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_users := make([]types.User, 0)
			page_buffer := scraper.Request("https://api.github.com/repos/" + user + "/" + repo + "/stargazers?page=" + page)

			if len(page_buffer) > 0 {

				err2 := json.Unmarshal(page_buffer, &page_users)

				if err2 == nil {

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
				err = errors.New("403 Unauthorized")
				break
			}

		}

	} else {
		err = errors.New("403 Unauthorized")
	}

	return users, err

}

