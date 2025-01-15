// You can edit this code!
// Click here and start typing.
package main

import (
	"io"
	"net/http"
)

func GetWithHeader(client *http.Client, url string, header *http.Header) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = *header

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
	GetWithHeader(client, "https://www.sweetscape.com/010editor/repository/templates/", nil)
}
