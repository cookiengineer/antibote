package structs

import "antibote/types"
import "encoding/json"
import "os"
import "strings"

type Cache struct {
	Folder  string                 `json:"folder"`
	GitHub  map[string]*types.User `json:"github"`
	KeyMap  map[string][]string    `json:"keymap"`
	TaskMap map[string]string      `json:"taskmap"`
}

func NewCache(folder string) Cache {

	var cache Cache

	if strings.HasSuffix(folder, "/") {
		folder = folder[0 : len(folder)-1]
	}

	stat, err1 := os.Stat(folder)

	if err1 == nil && stat.IsDir() {

		cache.Folder = folder
		cache.GitHub = make(map[string]*types.User)
		cache.KeyMap = make(map[string][]string, 0)
		cache.TaskMap = make(map[string]string)

	} else {

		err2 := os.MkdirAll(folder, 0750)
		err3 := os.MkdirAll(folder + "/github", 0750)
		err4 := os.MkdirAll(folder + "/downloads", 0750)

		if err2 == nil && err3 == nil && err4 == nil {

			cache.Folder = folder
			cache.GitHub = make(map[string]*types.User)
			cache.KeyMap = make(map[string][]string, 0)
			cache.TaskMap = make(map[string]string)

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

	buffer1, err1 := os.ReadFile(cache.Folder + "/keymap.json")

	if err1 == nil {
		json.Unmarshal(buffer1, &cache.KeyMap)
	}

	buffer2, err2 := os.ReadFile(cache.Folder + "/taskmap.json")

	if err2 == nil {
		json.Unmarshal(buffer2, &cache.TaskMap)
	}

}

func (cache *Cache) Write() {

	for name, _ := range cache.GitHub {
		cache.WriteUser(name)
	}

	buffer1, err1 := json.MarshalIndent(cache.KeyMap, "", "\t")

	if err1 == nil {
		os.WriteFile(cache.Folder + "/keymap.json", buffer1, 0666)
	}

	buffer2, err2 := json.MarshalIndent(cache.TaskMap, "", "\t")

	if err2 == nil {
		os.WriteFile(cache.Folder + "/taskmap.json", buffer2, 0666)
	}

}
