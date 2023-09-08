package activeplayer

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (i impl) GetName() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/activeplayername", i.address))
	if err != nil {
		return "", errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}
	var name string
	err = json.Unmarshal(resBody, &name)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return name, nil
}
