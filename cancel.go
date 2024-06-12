package go_worldwide_api_wrapper

import (
	"encoding/json"

	"github.com/idoubi/goz"
)

const cancelUrl = "/open-api/v2/cancel"

type CancelRequest struct {
	CustomOrderSn string `json:"customOrderSn"`
}

func DoCancel(id, key, orderSn string) (*Response, error) {
	// req
	req := CancelRequest{CustomOrderSn: orderSn}

	client := goz.NewClient()
	response, err := client.Post(getRequestUrl(cancelUrl), goz.Options{
		Timeout: 3000,
		Headers: getHeader(id, key),
		JSON:    req,
	})
	if err != nil {
		return nil, err
	}
	body, err := response.GetBody()
	if err != nil {
		return nil, err
	}
	rep := &Response{}
	err = json.Unmarshal(body, rep)
	if err != nil {
		return nil, err
	}

	return rep, nil
}
