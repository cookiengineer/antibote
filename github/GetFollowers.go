package github

import "antibote/constants"
import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "strconv"

func GetFollowers(cache *structs.Cache, user string) []types.User {

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": constants.Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	followers := make([]types.User, 0)
	buffer := scraper.Request("https://api.github.com/users/" + user + "/followers?page=1")
	err := json.Unmarshal(buffer, &followers)

	if err == nil && len(followers) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_followers := make([]types.User, 0)
			page_buffer := scraper.Request("https://api.github.com/users/" + user + "/followers?page=" + page)

			if len(page_buffer) > 0 {

				err := json.Unmarshal(page_buffer, &page_followers)

				if err == nil {

					for pf := 0; pf < len(page_followers); pf++ {
						followers = append(followers, page_followers[pf])
					}

					if len(page_followers) < 30 {
						break
					}

				} else {
					break
				}

			}

		}

	}

	return followers

}
