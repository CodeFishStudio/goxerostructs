package goxerostructs

import "time"

//InvoiceResult is an Invoice result from xero
type InvoiceResult struct {
	ID           string    `json:"ID,omitempty"`
	ProviderName string    `json:"ProviderName,omitempty"`
	Status       string    `json:"Status,omitempty"`
	DateCreated  time.Time `json:"DateCreated,omitempty"`
	Params       string    `json:"Params,omitempty"`
	Invoices     []Invoice `json:"Invoices,omitempty"`
}

//Invoice is an Invoice from xero
type Invoice struct {
	Type            string     `json:"Type,omitempty" xml:"Type,omitempty"`
	Contact         Contact    `json:"Contact,omitempty" xml:"Contact,omitempty"`
	Date            string     `json:"Date,omitempty" xml:"Date,omitempty"`
	DateString      string     `json:"DateString,omitempty" xml:"DateString,omitempty"`
	DueDate         string     `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Status          string     `json:"Status,omitempty" xml:"Status,omitempty"`
	InvoiceNumber   string     `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`
	LineAmountTypes string     `json:"LineAmountTypes,omitempty" xml:"LineAmountTypes,omitempty"`
	InvoiceID       string     `json:"InvoiceID,omitempty" xml:"InvoiceID,omitempty"`
	TotalTax        float64    `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`
	Total           float64    `json:"Total,omitempty" xml:"Total,omitempty"`
	Tracking        []Tracking `json:"Tracking,omitempty" xml:"Tracking,omitempty"`
	LineItems       []LineItem `json:"LineItems,omitempty" xml:"LineItems,omitempty"`
}
