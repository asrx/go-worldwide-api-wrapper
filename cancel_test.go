package go_worldwide_api_wrapper

import (
	"fmt"
	"testing"
)

func Test_Cancel(t *testing.T) {
	apiId := ""
	apiKey := ""
	orderSn := ""

	rep, err := DoCancel(apiId, apiKey, orderSn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------")
	fmt.Println(rep.Code)
	fmt.Printf("%+v\n", rep)
	fmt.Println("----------------")
}
