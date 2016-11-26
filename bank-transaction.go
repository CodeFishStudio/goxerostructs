package goxero

//BankTransaction  a bank transaction from xero
type BankTransaction struct {
	BankTransactionID string  `json:"BankTransactionID"`
	Type              string  `json:"Type"`
	Reference         string  `json:"Reference"`
	IsReconciled      bool    `json:"IsReconciled"`
	HasAttachments    bool    `json:"HasAttachments"`
	DateString        string  `json:"DateString"`
	Status            string  `json:"Status"`
	LineAmountTypes   string  `json:"LineAmountTypes"`
	SubTotal          float64 `json:"SubTotal"`
	Total             float64 `json:"Total"`
	TotalTax          float64 `json:"TotalTax"`
	CurrencyCode      float64 `json:"CurrencyCode"`
}
