package go_worldwide_api_wrapper

type ShipFrom struct {
	Code string `json:"code"`
}

type ShipTo struct {
	ShipInfo ShipInfo `json:"shipInfo"`
}

type ShipInfo struct {
	Address   Address `json:"address"`
	Name      string  `json:"name"`
	Attention string  `json:"attention"`
	Phone     Phone   `json:"phone"`
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
	Length     string       `json:"length"`
	Width      string       `json:"width"`
	Height     string       `json:"height"`
	Weight     string       `json:"weight"`
	Insurance  string       `json:"insurance"`
	Quantity   int          `json:"quantity"`
	References []References `json:"references"`
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
