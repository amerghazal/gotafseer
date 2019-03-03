package gotafseer

import (
	"encoding/json"
	//	"fmt"
	"net/http"
	"net/url"
)

type TafseerApiClient struct {
	BaseURL *url.URL
}

var (
	QuranPath   url.URL = url.URL{Path: "/quran"}
	TafseerPath url.URL = url.URL{Path: "/tafseer"}
)

func (c *TafseerApiClient) ListChapters() ([]Chapter, error) {
	ResourcePath := c.BaseURL.ResolveReference(&QuranPath)
	resp, err := http.Get(ResourcePath.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var chapters []Chapter
	err = json.NewDecoder(resp.Body).Decode(&chapters)
	return chapters, err
}

func (c *TafseerApiClient) ListTafseers() ([]Tafseer, error) {
	ResourcePath := c.BaseURL.ResolveReference(&TafseerPath)
	resp, err := http.Get(ResourcePath.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tafseers []Tafseer
	err = json.NewDecoder(resp.Body).Decode(&tafseers)
	return tafseers, err
}
