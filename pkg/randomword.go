package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quote/model"
)

func RandomWordsGet() ([]string, error) {
	var client = &http.Client{}

	req, err := http.NewRequest("GET", "https://random-indonesian-word.p.rapidapi.com/words/random?limit=5", nil)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	req.Header.Set("X-RapidAPI-Host", "random-indonesian-word.p.rapidapi.com")
	req.Header.Set("X-RapidAPI-Key", "e9b079b792msh3e1aa8608df520dp15d272jsn0831149d10d6")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	var data model.RandomWordsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return data.Data, nil
}
