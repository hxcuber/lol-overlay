package item

import (
	"encoding/json"
	"fmt"
	"github.com/hxcuber/lol-overlay/model/datadragon/fullitems"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
)

func (i impl) GetPrice(id int) (int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/item.json", i.address))
	if err != nil {
		return 0, errors.WithStack(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	var data struct {
		Data map[string]fullitems.FullItem `json:"data"`
	}
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return data.Data[strconv.FormatInt(int64(id), 10)].Gold.Total, nil
}
