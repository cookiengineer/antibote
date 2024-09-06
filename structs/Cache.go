package structs

import github "antibote/github/api"
import "encoding/json"
import "os"
import "strings"

type Cache struct {
	Folder string                  `json:"folder"`
	GitHub map[string]*github.User `json:"github"`
	KeyMap map[string][]string     `json:"keymap"`
}

func NewCache(folder string) Cache {

	var cache Cache

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	stat, err1 := os.Stat(folder + "/github")

	if err1 == nil && stat.IsDir() {

		cache.Folder = folder
		cache.GitHub = make(map[string]*github.User)
		cache.KeyMap = make(map[string][]string, 0)

	} else {

		err2 := os.MkdirAll(folder, 0750)

		if err2 == nil {

			cache.Folder = folder
			cache.GitHub = make(map[string]*github.User)
			cache.KeyMap = make(map[string][]string, 0)

		}

	}

	return cache

}

func (cache *Cache) Read() {

	stat, err1 := os.Stat(cache.Folder + "/github")

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(cache.Folder + "/github")

		if err2 == nil {

			for e := 0; e < len(entries); e++ {

				name := entries[e].Name()

				if strings.HasSuffix(name, ".json") {
					cache.ReadUser(name[0:len(name)-5])
				}

			}

		}

	}
	// TODO: Read from filesystem

}

func (cache *Cache) ReadUser(value string) bool {

	var result bool = false

	stat, err1 := os.Stat(cache.Folder + "/github/" + value + ".json")

	if err1 == nil && !stat.IsDir() {

		buffer, err2 := os.ReadFile(cache.Folder + "/github/" + value + ".json")

		if err2 == nil {

			var user github.User

			err3 := json.Unmarshal(buffer, &user)

			if err3 == nil {
				cache.AddUser(&user)
			}

		}

	}

	return result

}

func (cache *Cache) AddUser(value *github.User) {

	cache.GitHub[value.Name] = value

	keys := value.ToKeys()

	if len(keys) > 0 {

		for k := 0; k < len(keys); k++ {

			key := keys[k]

			users, ok := cache.KeyMap[key]

			if ok {

				found := false

				for u := 0; u < len(users); u++ {

					if users[u] == value.Name {
						found = true
						break
					}

				}

				if found == false {
					cache.KeyMap[key] = append(cache.KeyMap[key], value.Name)
				}

			} else {
				cache.KeyMap[key] = []string{value.Name}
			}

		}

	}

}

func (cache *Cache) GetUser(value string) *github.User {

	tmp, ok := cache.GitHub[value]

	if ok {
		return tmp
	}

	return nil

}

func (cache *Cache) Write() {

	for name, user := range cache.GitHub {

		file := cache.Folder + "/github/" + name + ".json"

		buffer, err := json.MarshalIndent(user, "", "\t")

		if err == nil {
			os.WriteFile(file, buffer, 0666)
		}

	}

	buffer, err := json.MarshalIndent(cache.KeyMap, "", "\t")

	if err == nil {
		os.WriteFile(cache.Folder + "/keymap.json", buffer, 0666)
	}

}
