package gotafseer

import (
	"encoding/json"
	"fmt"
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
	var ch []Chapter
	r := c.BaseURL.ResolveReference(&QuranPath)
	resp, err := http.Get(r.String())
	if err != nil {
		return ch, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return ch, err
	}
	return ch, nil
}

func (c *TafseerApiClient) ListTafseers() ([]Tafseer, error) {
	var t []Tafseer
	r := c.BaseURL.ResolveReference(&TafseerPath)
	resp, err := http.Get(r.String())
	if err != nil {
		return t, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (c *TafseerApiClient) GetVerse(s string, v string) (Verse, error) {
	var vs Verse
	r := c.BaseURL.ResolveReference(&QuranPath)
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", r.String(), s, v))
	if err != nil {
		return vs, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&vs)
	if err != nil {
		return vs, err
	}
	return vs, nil
}

func (c *TafseerApiClient) GetVerseTafseer(t string, s string, v string) (VerseTafseer, error) {
	var vt VerseTafseer
	r := c.BaseURL.ResolveReference(&TafseerPath)
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/%s", r.String(), t, s, v))
	if err != nil {
		return vt, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&vt)
	if err != nil {
		return vt, err
	}
	return vt, nil
}

func (c *TafseerApiClient) GetVerseTafseerRange(t string, s string, vs string, ve string) ([]VerseTafseer, error) {
	var vt []VerseTafseer
	r := c.BaseURL.ResolveReference(&TafseerPath)
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/%s/%s", r.String(), t, s, vs, ve))
	if err != nil {
		return vt, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&vt)
	if err != nil {
		return vt, err
	}
	return vt, nil
}
