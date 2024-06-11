package go_worldwide_api_wrapper

import (
	"fmt"
	"testing"
)

func TestAddrVerify(t *testing.T) {
	apiId := ""
	apiKey := ""

	addr := AddrVerifyRequest{
		StateCode:      "VA",
		City:           "Nokesville",
		DetailAddress:  "12180 Hazelwood Dr",
		DetailAddress2: "",
		PostCode:       "20181",
		Residential:    "true",
	}

	rep, err := DoAddrVerify(apiId, apiKey, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----------------")
	fmt.Println(rep)
	fmt.Println("----------------")
}
