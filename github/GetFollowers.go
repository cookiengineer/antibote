package github

import "antibote/constants"
import "antibote/structs"
import "antibote/types"
import "encoding/json"
import "errors"
import "strconv"

func GetFollowers(cache *structs.Cache, user string) ([]types.User, error) {

	var err error = nil

	scraper := structs.NewScraper(cache, 1)
	scraper.Headers = map[string]string{
		"Accept": "application/json",
		"Authorization": constants.Token,
		"User-Agent": "antibote (Bot Detector)",
	}

	followers := make([]types.User, 0)
	buffer := scraper.Request("https://api.github.com/users/" + user + "/followers?page=1")
	err1 := json.Unmarshal(buffer, &followers)

	if err1 == nil && len(followers) == 30 {

		for p := 2; p <= 50; p++ {

			page := strconv.Itoa(p)
			page_followers := make([]types.User, 0)
			page_buffer := scraper.Request("https://api.github.com/users/" + user + "/followers?page=" + page)

			if len(page_buffer) > 0 {

				err2 := json.Unmarshal(page_buffer, &page_followers)

				if err2 == nil {

					for pf := 0; pf < len(page_followers); pf++ {
						followers = append(followers, page_followers[pf])
					}

					if len(page_followers) < 30 {
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

	} else if len(buffer) == 0 {
		err = errors.New("403 Unauthorized")
	}

	return followers, err

}
