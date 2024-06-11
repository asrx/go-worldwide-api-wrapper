package go_worldwide_api_wrapper

import (
	"github.com/idoubi/goz"
)

const shipmentUrl = "/open-api/v2/rateAndShipment"

type ShipmentRequest struct {
	CustomOrderSn string   `json:"customOrderSn"`
	AsyncShipment bool     `json:"asyncShipment"`
	Shipment      Shipment `json:"shipment"`
}

type Shipment struct {
	ServiceCode     string          `json:"serviceCode"`
	ShipFrom        ShipFrom        `json:"shipFrom"`
	ShipTo          ShipTo          `json:"shipTo"`
	Unit            string          `json:"unit"`
	Confirmation    string          `json:"confirmation"`
	Packages        []Package       `json:"packages"`
	SpecialServices SpecialServices `json:"specialServices"`
	ReturnService   bool            `json:"returnService"`
}

// 无需对接
type SpecialServices struct {
	PushManifest int `json:"pushManifest"`
	ConfirmLabel int `json:"confirmLabel"`
}

func NewShipmentRequest(orderNo string, shipFromCode string, shipTo ShipTo, packages []Package, signature bool) *ShipmentRequest {
	confirmation := "None"
	if signature {
		confirmation = "Adult"
	}
	return &ShipmentRequest{
		CustomOrderSn: orderNo,
		AsyncShipment: false,
		Shipment: Shipment{
			ServiceCode:  "",
			ShipFrom:     ShipFrom{shipFromCode},
			ShipTo:       shipTo,
			Unit:         "I",
			Confirmation: confirmation,
			Packages:     packages,
			// SpecialServices: SpecialServices{},
			ReturnService: false,
		},
	}
}

func GetShipment(orderNo, id, key, shipFromCode string, shipTo ShipTo, packages []Package, signature bool) (rep interface{}, err error) {
	request := NewShipmentRequest(orderNo, shipFromCode, shipTo, packages, signature)
	client := goz.NewClient()
	response, err := client.Post(getRequestUrl(shipmentUrl), goz.Options{
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

	return body.String(), nil
}
