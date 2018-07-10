package goxerostructs

import "time"

//OrganisationResult is an Organisation result from xero
type OrganisationResult struct {
	ID            string         `json:"ID"`
	ProviderName  string         `json:"ProviderName"`
	Status        string         `json:"Status"`
	Organisations []Organisation `json:"Contacts"`
	DateCreated   time.Time      `json:"DateCreated"`
}

//Organisation is an Organisation from xero
type Organisation struct {
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`
}
