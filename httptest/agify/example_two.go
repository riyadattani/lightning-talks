package agify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Agify struct {
	url string
}

func NewAgify(url string) *Agify {
	return &Agify{url: url}
}

func (a Agify) GetAge(name string) (int, error) {
	res, err := http.Get(a.url + "/?name=" + name)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("uh oh, err: %v", err)
	}

	var apiRes APIResponse
	err = json.NewDecoder(res.Body).Decode(&apiRes)
	if err != nil {
		return 0, err
	}

	return apiRes.Age, nil
}

type APIResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}
