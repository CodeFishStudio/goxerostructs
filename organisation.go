package goxerostructs

//OrganisationResult is an Organisation result from xero
type OrganisationResult struct {
	ID            string         `json:"ID"`
	ProviderName  string         `json:"ProviderName"`
	Status        string         `json:"Status"`
	Organisations []Organisation `json:"Contacts"`
}

//Organisation is an Organisation from xero
type Organisation struct {
	Name string `json:"Name,omitempty" xml:"Name,omitempty"`
}
