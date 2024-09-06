package structs

import "os"
import "path"
import "strings"

func toFilePath(url string) string {

	var result string

	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {

		tmp1 := strings.Split(url[8:], "/")
		tmp2 := []string{}

		for t := 0; t < len(tmp1); t++ {

			chunk := strings.TrimSpace(tmp1[t])

			if chunk != "" {
				tmp2 = append(tmp2, chunk)
			}

		}

		result = strings.Join(tmp2, "/")

	}

	return result

}

func (cache *Cache) HasDownload(url string) bool {

	var result bool = false

	file := toFilePath(url)

	if file != "" {

		stat, err := os.Stat(cache.Folder + "/downloads/" + file)

		if err == nil && !stat.IsDir() {
			result = true
		}

	}

	return result

}

func (cache *Cache) ReadDownload(url string) []byte {

	var result []byte

	file := toFilePath(url)

	if file != "" {

		stat, err1 := os.Stat(cache.Folder + "/downloads/" + file)

		if err1 == nil && !stat.IsDir() {

			buffer, err2 := os.ReadFile(cache.Folder + "/" + file)

			if err2 == nil {
				result = buffer
			}

		}

	}

	return result

}

func (cache *Cache) WriteDownload(url string, buffer []byte) bool {

	var result bool = false

	file := toFilePath(url)

	if file != "" {

		folder := path.Dir(file)
		stat, err1 := os.Stat(cache.Folder + "/downloads/" + folder)

		if err1 == nil && stat.IsDir() {

			err2 := os.WriteFile(cache.Folder + "/downloads/" + file, buffer, 0666)

			if err2 == nil {
				result = true
			}

		} else {

			err2 := os.MkdirAll(cache.Folder + "/downloads/" + folder, 0755)

			if err2 == nil {

				err3 := os.WriteFile(cache.Folder + "/downloads/" + file, buffer, 0666)

				if err3 == nil {
					result = true
				}

			}

		}

	}

	return result

}
