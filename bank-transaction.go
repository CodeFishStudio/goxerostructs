package goxerostructs

//BankTransactionResult  a bank transaction from xero
type BankTransactionResult struct {
	ID               string            `json:"ID"`
	ProviderName     string            `json:"ProviderName"`
	Status           string            `json:"Status"`
	BankTransactions []BankTransaction `json:"BankTransactions"`
}

//BankTransaction  a bank transaction from xero
type BankTransaction struct {
	BankTransactionID string     `json:"BankTransactionID"`
	Type              string     `json:"Type"`
	Reference         string     `json:"Reference"`
	IsReconciled      bool       `json:"IsReconciled"`
	HasAttachments    bool       `json:"HasAttachments"`
	DateString        string     `json:"DateString"`
	Status            string     `json:"Status"`
	LineAmountTypes   string     `json:"LineAmountTypes"`
	SubTotal          float64    `json:"SubTotal"`
	Total             float64    `json:"Total"`
	TotalTax          float64    `json:"TotalTax"`
	CurrencyCode      string     `json:"CurrencyCode"`
	Lineitems         []LineItem `json:"Lineitems"`
	Contact           Contact    `json:"Contact"`
}

//LineItem is a line item of a bank transaction
type LineItem struct {
	Description string     `json:"Description,omitempty" xml:"Description,omitempty"`
	AccountCode string     `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`
	TaxType     string     `json:"TaxType,omitempty" xml:"TaxType,omitempty"`
	TaxAmount   float64    `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	UnitAmount  float64    `json:"UnitAmount,omitempty" xml:"UnitAmount,omitempty"`
	Quantity    float64    `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	LineAmount  float64    `json:"LineAmount,omitempty" xml:"LineAmount,omitempty"`
	Tracking    []Tracking `json:"Tracking,omitempty" xml:"Tracking,omitempty"`
}
