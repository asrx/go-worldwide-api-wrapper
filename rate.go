package go_worldwide_api_wrapper

import (
	"encoding/json"
	"errors"

	"github.com/idoubi/goz"
)

const rateUrl = "/open-api/v2/getRates"

type RateReuqest struct {
	ServiceCode   string    `json:"serviceCode"`
	ShipFrom      ShipFrom  `json:"shipFrom"`
	ShipTo        ShipTo    `json:"shipTo"`
	Unit          string    `json:"unit"`
	Confirmation  string    `json:"confirmation"`
	Packages      []Package `json:"packages"`
	ReturnService bool      `json:"returnService"`
}

func NewRateRequest(shipFromCode string, shipTo ShipTo, packages []Package, signature bool) *RateReuqest {
	confirmation := "None"
	if signature {
		confirmation = "Adult"
	}
	return &RateReuqest{
		ServiceCode:   "",
		ShipFrom:      ShipFrom{shipFromCode},
		ShipTo:        shipTo,
		Unit:          "I",
		Confirmation:  confirmation,
		Packages:      packages,
		ReturnService: false,
	}
}

func GetRate(id, key, shipFromCode string, shipTo ShipTo, packages []Package, signature bool) (*ResponseRate, error) {
	request := NewRateRequest(shipFromCode, shipTo, packages, signature)
	client := goz.NewClient()
	response, err := client.Post(getRequestUrl(rateUrl), goz.Options{
		Timeout: 3000,
		Headers: getHeader(id, key),
		JSON:    request,
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
	if rep.Code != 200 {
		return nil, errors.New(rep.Msg)
	}

	rep2 := &ResponseRate{}
	err = json.Unmarshal(body, rep2)
	if err != nil {
		return nil, err
	}

	return rep2, nil
}
