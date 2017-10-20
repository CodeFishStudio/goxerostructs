package goxerostructs

import "time"

//InvoiceResult is an Invoice result from xero
type InvoiceResult struct {
	ID           string    `json:"ID"`
	ProviderName string    `json:"ProviderName"`
	Status       string    `json:"Status"`
	DateCreated  time.Time `json:"DateCreated"`
	Params       string    `json:"Params"`
	Invoices     []Invoice `json:"Invoices"`
}

//Invoice is an Invoice from xero
type Invoice struct {
	Type            string     `json:"Type"`
	Contact         Contact    `json:"Contact"`
	Date            string     `json:"DateString"`
	DueDate         string     `json:"DueDate"`
	Status          string     `json:"Status"`
	InvoiceNumber   string     `json:"InvoiceNumber"`
	LineAmountTypes string     `json:"LineAmountTypes"`
	InvoiceID       string     `json:"InvoiceID"`
	TotalTax        float64    `json:"TotalTax"`
	Total           float64    `json:"Total"`
	Tracking        []Tracking `json:"Tracking"`
	LineItems       []LineItem `json:"LineItems"`
}
