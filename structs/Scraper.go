package structs

import "antibote/console"
import "net/http"
import "io/ioutil"
import "strconv"
import "strings"
import "time"

var content_types []string = []string{
	"application/gzip",
	"application/json",
	"application/ld+json",
	"application/octet-stream",
	"application/rss+xml",
	"application/x-bzip2",
	"application/x-gzip",
	"application/xml",
	"application/zip",
	"text/html",
	"text/plain",
	"text/xml",
}

type Callback func([]byte)

type ScraperTask struct {
	Url      string
	Callback Callback
}

type Scraper struct {
	Cache     *Cache
	Busy      bool
	Limit     int
	Tasks     []ScraperTask
	Headers   map[string]string
	Throttled bool
}

func processRequests(scraper *Scraper) {

	var filtered []ScraperTask
	var limit int = scraper.Limit

	if scraper.Throttled == true {
		limit = 1
	}

	for t := 0; t < len(scraper.Tasks); t++ {

		if len(filtered) < limit {
			filtered = append(filtered, scraper.Tasks[t])
		} else {
			break
		}

	}

	if len(filtered) > 0 {

		for f := 0; f < len(filtered); f++ {

			task := filtered[f]
			buffer := scraper.Request(task.Url)
			task.Callback(buffer)

		}

		scraper.Tasks = scraper.Tasks[len(filtered):]

		if len(scraper.Tasks) > 0 {

			if scraper.Throttled == true {

				console.Log(strconv.Itoa(len(scraper.Tasks)) + " Request Tasks left...")

				time.AfterFunc(30*time.Second, func() {
					processRequests(scraper)
				})

			} else {

				time.AfterFunc(1*time.Second, func() {
					processRequests(scraper)
				})

			}

		} else {

			scraper.Busy = false

		}

	}

}

func NewScraper(cache *Cache, limit int) Scraper {

	if limit <= 0 {
		limit = 1
	}

	var scraper Scraper

	scraper.Cache = cache
	scraper.Busy = false
	scraper.Limit = limit
	scraper.Tasks = make([]ScraperTask, 0)
	scraper.Headers = make(map[string]string, 0)
	scraper.Throttled = false

	return scraper

}

func (scraper *Scraper) DeferRequest(url string, callback Callback) {

	scraper.Tasks = append(scraper.Tasks, ScraperTask{
		Url:      url,
		Callback: callback,
	})

	if scraper.Busy == false {

		scraper.Busy = true

		time.AfterFunc(1*time.Second, func() {
			processRequests(scraper)
		})

	}

}

func (scraper *Scraper) Request(url string) []byte {

	var buffer []byte
	var content_type string
	var status_code int

	if scraper.Cache.HasDownload(url) {

		buffer = scraper.Cache.ReadDownload(url)

		if len(buffer) > 0 {
			content_type = "application/json"
			status_code = 200
		}

	} else {

		client := &http.Client{}
		client.CloseIdleConnections()

		request, err1 := http.NewRequest("GET", url, nil)

		if err1 == nil {

			for key, val := range scraper.Headers {
				request.Header.Set(key, val)
			}

			response, err2 := client.Do(request)

			if err2 == nil {

				status_code = response.StatusCode

				if status_code == 200 || status_code == 304 {

					if len(response.Header["Content-Type"]) > 0 {
						content_type = response.Header["Content-Type"][0]
					} else {
						content_type = "application/octet-stream"
					}

					var valid bool = false

					for c := 0; c < len(content_types); c++ {

						if strings.Contains(content_type, content_types[c]) {
							valid = true
							break
						}

					}

					if valid == true {

						data, err2 := ioutil.ReadAll(response.Body)

						if err2 == nil {
							buffer = data
						}

					}

				}

			}

			if len(buffer) > 0 {
				scraper.Cache.WriteDownload(url, buffer)
			}

		}

		if len(buffer) > 0 {

			console.Log("Request \"" + url + "\" succeeded")

		} else {

			console.Error("Request \"" + url + "\" failed")

			if content_type != "" {
				console.Error("Unsupported Content-Type \"" + content_type + "\"")
			}

			if status_code != 0 {
				console.Error("Unsupported Status Code \"" + strconv.Itoa(status_code) + "\"")
			}

		}

	}

	return buffer

}
