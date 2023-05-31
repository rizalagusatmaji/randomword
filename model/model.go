package model

type RandomWordsResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type DataRespone struct {
	Words []string `json:"data"`
}
