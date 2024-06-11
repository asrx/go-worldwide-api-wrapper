package go_worldwide_api_wrapper

import (
	"fmt"
	"testing"
)

func TestShipment(t *testing.T) {
	apiId := ""
	apiKey := ""
	shipFromCode := ""
	customerOrderNo := ""

	shipTo := ShipTo{
		ShipInfo{
			Address: Address{
				StateCode:      "VA",
				City:           "Nokesville",
				DetailAddress:  "12180 Hazelwood Dr",
				DetailAddress2: "",
				PostCode:       "20181",
				Residential:    true,
				CountryCode:    "US",
			},
			Name: "Donovan.xu",
			Phone: Phone{
				Number:    "8000000000",
				Extension: "",
			},
		}}
	var pkgs = make([]Package, 0)
	p := Package{
		Length:    "9.15",
		Width:     "6.15",
		Height:    "5.15",
		Weight:    "5.15",
		Insurance: "300",
		Quantity:  1,
	}
	pkgs = append(pkgs, p)

	rep, err := GetShipment(customerOrderNo, apiId, apiKey, shipFromCode, shipTo, pkgs, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----------------")
	fmt.Println(rep)
	fmt.Println("----------------")
}
