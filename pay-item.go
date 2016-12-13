package goxerostructs

//PayItemResult is an pay item result from xero
type PayItemResult struct {
	ID           string  `json:"ID"`
	ProviderName string  `json:"ProviderName"`
	Status       string  `json:"Status"`
	PayItems     PayItem `json:"PayItems"`
}

//PayItem is an employee from xero
type PayItem struct {
	EarningsRates []EarningsRate `json:"EarningsRates"`
	LeaveTypes    []LeaveType    `json:"LeaveTypes"`
}

//EarningsRate is an earning rate from xero
type EarningsRate struct {
	Name              string  `json:"Name"`
	IsExemptFromTax   bool    `json:"IsExemptFromTax"`
	IsExemptFromSuper bool    `json:"IsExemptFromSuper"`
	EarningsType      string  `json:"EarningsType"`
	EarningsRateID    string  `json:"EarningsRateID"`
	RateType          string  `json:"RateType"`
	RatePerUnit       float64 `json:"RatePerUnit"`
	Multiplier        float64 `json:"Multiplier"`
	AccrueLeave       bool    `json:"AccrueLeave"`
	Amount            float64 `json:"Amount"`
}

//LeaveType is an earning rate from xero
type LeaveType struct {
	Name string `json:"Name"`
}
