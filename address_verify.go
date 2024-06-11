package go_worldwide_api_wrapper

import (
	"github.com/idoubi/goz"
)

const addressVerify = "/open-api/v2/addressVerify"

type AddrVerifyRequest struct {
	StateCode      string `json:"stateCode"`
	City           string `json:"city"`
	DetailAddress  string `json:"detailAddress"`
	DetailAddress2 string `json:"detailAddress2"`
	PostCode       string `json:"postCode"`
	Residential    string `json:"residential"`
}

func getAddrVerify(id, key string, addr AddrVerifyRequest) (rep interface{}, err error) {
	client := goz.NewClient()
	response, err := client.Post(getRequestUrl(addressVerify), goz.Options{
		Timeout: 3000,
		Headers: getHeader(id, key),
		JSON:    addr,
	})
	if err != nil {
		return nil, err
	}
	body, err := response.GetBody()
	if err != nil {
		return nil, err
	}

	return body.String(), nil
}
