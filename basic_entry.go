package go_worldwide_api_wrapper

type ShipFrom struct {
	Code string `json:"code"`
}

type ShipTo struct {
	ShipInfo *ShipInfo `json:"shipInfo"`
}

type ShipInfo struct {
	Address   *Address `json:"address"`
	Name      string   `json:"name"`
	Attention string   `json:"attention"`
	Phone     *Phone   `json:"phone"`
}

type Address struct {
	StateCode      string `json:"stateCode"`
	City           string `json:"city"`
	DetailAddress  string `json:"detailAddress"`
	DetailAddress2 string `json:"detailAddress2"`
	PostCode       string `json:"postCode"`
	Residential    bool   `json:"residential"`
	CountryCode    string `json:"countryCode"`
}

type Phone struct {
	Number    string `json:"number"`
	Extension string `json:"extension"`
}

// -------------------
type Package struct {
	Length     string        `json:"length"`
	Width      string        `json:"width"`
	Height     string        `json:"height"`
	Weight     string        `json:"weight"`
	Insurance  string        `json:"insurance"`
	Quantity   int           `json:"quantity"`
	References []*References `json:"references"`
}

/*
reference #1: 备注编码1 UPS的 reference #1；Fedex面单上的 REF
reference #2: 备注编码2 UPS的 reference #2；Fedex面单上的 PO
INVOICE_NUMBER: UPS没有；FedEx面单上的INV；
DEPARTMENT_NUMBER: UPS没有；FedEx面单上的DEPT；
*/
type References struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

// ------- 响应结构 -----------
type Response struct {
	Code          int         `json:"code"`
	Success       bool        `json:"success"`
	Data          interface{} `json:"data"`
	Msg           string      `json:"msg"`
	TransactionId string      `json:"transactionId"`
}

// ------- 估价结构体 ----------
type ResponseRate struct {
	*Response
	Data *RateStruct `json:"data"`
}

type RateStruct struct {
	Rates []*RateInfo `json:"rates,omitempty"`
}

type RateInfo struct {
	RateId              string `json:"rateId"`
	Price               string `json:"price"`
	TransportationPrice string `json:"transportationPrice"`
	InsurancePrice      string `json:"insurancePrice"`
	Carrier             string `json:"carrier"`
	ServiceCode         string `json:"serviceCode"`
	ServiceName         string `json:"serviceName"`
	Zone                string `json:"zone"`
	OneRate             int    `json:"oneRate"`
	ChannelName         string `json:"channelName"`
	ChannelType         int    `json:"channelType"`
	PriceItems          []struct {
		Name  string `json:"name"`
		Money string `json:"money"`
	} `json:"priceItems"`
	EstimatedTime   string `json:"estimatedTime"`
	EstimatedDays   string `json:"estimatedDays"`
	FromCode        string `json:"fromCode"`
	FromAddressCode string `json:"fromAddressCode"`
	FromAddressInfo string `json:"fromAddressInfo"`
}

// ---------- 估价下单结构体 -----------
type ResponseShip struct {
	*Response
	Data []*ShipStruct `json:"data"`
}

type ShipStruct struct {
	OrderSn       string `json:"orderSn"`
	CustomOrderSn string `json:"customOrderSn"`
	Status        string `json:"status"`
	Price         string `json:"price"`
	Carrier       string `json:"carrier"`
	ServiceCode   string `json:"serviceCode"`
	ServiceName   string `json:"serviceName"`
	Labels        []struct {
		PackageSn      string `json:"packageSn"`
		TrackingNumber string `json:"trackingNumber"`
		LabelUrl       string `json:"labelUrl"`
		LabelUrl2      string `json:"labelUrl2"`
		Status         string `json:"status"`
	} `json:"labels"`
	Msg string `json:"msg"`
}
