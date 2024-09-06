package structs

import "antibote/types"
import "encoding/json"
import "os"

func (cache *Cache) GetUser(value string) *types.User {

	tmp, ok := cache.GitHub[value]

	if ok {
		return tmp
	}

	return nil

}

func (cache *Cache) ReadUser(value string) bool {

	var result bool = false

	stat, err1 := os.Stat(cache.Folder + "/github/" + value + ".json")

	if err1 == nil && !stat.IsDir() {

		buffer, err2 := os.ReadFile(cache.Folder + "/github/" + value + ".json")

		if err2 == nil {

			var user types.User

			err3 := json.Unmarshal(buffer, &user)

			if err3 == nil {
				cache.TrackUser(&user)
			}

		}

	}

	return result

}

func (cache *Cache) TrackUser(value *types.User) {
	cache.GitHub[value.Name] = value
}

func (cache *Cache) WriteUser(value string) bool {

	var result bool = false

	user, ok := cache.GitHub[value]

	if ok == true {

		buffer, err1 := json.MarshalIndent(user, "", "\t")

		if err1 == nil {

			err2 := os.WriteFile(cache.Folder + "/github/" + user.Name + ".json", buffer, 0666)

			if err2 == nil {
				result = true
			}

		}

	}

	return result

}
