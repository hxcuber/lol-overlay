package activeplayer

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetActivePlayerName() (string, error) {
	resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayer")
	if err != nil {
		return "", err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var name string
	err = json.Unmarshal(resBody, &name)
	if err != nil {
		return "", err
	}
	return name, nil
}
