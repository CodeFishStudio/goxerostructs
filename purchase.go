package goxerostructs

import "time"

//PurchaseResult is an Purchase result from xero
type PurchaseResult struct {
	ID             string     `json:"ID"`
	ProviderName   string     `json:"ProviderName"`
	Status         string     `json:"Status"`
	DateCreated    time.Time  `json:"DateCreated"`
	PurchaseOrders []Purchase `json:"PurchaseOrders"`
}

//Purchase is an Purchase from xero
type Purchase struct {
	Contact         Contact    `json:"Contact"`
	Date            string     `json:"DateString"`
	DeliveryDate    string     `json:"DeliveryDate"`
	Status          string     `json:"Status"`
	Reference       string     `json:"Reference"`
	LineAmountTypes string     `json:"LineAmountTypes"`
	PurchaseOrderID string     `json:"PurchaseOrderID"`
	TotalTax        float64    `json:"TotalTax"`
	Total           float64    `json:"Total"`
	Tracking        []Tracking `json:"Tracking"`
	LineItems       []LineItem `json:"LineItems"`
}
