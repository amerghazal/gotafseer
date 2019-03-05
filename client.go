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
	r := c.BaseURL.ResolveReference(&QuranPath)
	resp, err := http.Get(r.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ch []Chapter
	err = json.NewDecoder(resp.Body).Decode(&ch)
	return ch, err
}

func (c *TafseerApiClient) ListTafseers() ([]Tafseer, error) {
	r := c.BaseURL.ResolveReference(&TafseerPath)
	resp, err := http.Get(r.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var t []Tafseer
	err = json.NewDecoder(resp.Body).Decode(&t)
	return t, err
}
