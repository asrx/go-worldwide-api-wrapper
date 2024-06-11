package go_worldwide_api_wrapper

import (
	"fmt"

	"github.com/idoubi/goz"
)

const cancelUrl = "/open-api/v2/cancel"

type CancelRequest struct {
	CustomOrderSn string `json:"customOrderSn"`
}

func DoCancel(id, key, orderSn string) (rep interface{}, err error) {
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

	fmt.Println(body.String())

	return nil, nil
}
