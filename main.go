// You can edit this code!
// Click here and start typing.
package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
)

func GetWithHeader(client *http.Client, url string, header *http.Header) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = *header
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	var get = func(url string) ([]byte, error) {
		return GetWithHeader(client, url, nil)
	}

	bytes, err := get("https://www.sweetscape.com/010editor/repository/templates/")
	if err != nil {
		log.Fatal(err)
	}

	stringSubmatched := regexp.MustCompile(`href="(.*bt)">Download</a>`).FindAllStringSubmatch(string(bytes), -1)
	for _, t := range stringSubmatched {
		u, _ := url.Parse(t[1])
		base, err := url.Parse("https://www.sweetscape.com/010editor/repository/templates/")
		if err != nil {
			log.Fatal(err)
		}

		directLink := base.ResolveReference(u)
		bytes, err := get(directLink.String())
		if err != nil {
			log.Fatal(err)
		}

		os.WriteFile(path.Base(u.Path), bytes, 0644)
	}
}
