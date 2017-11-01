package goxerostructs

//ContactResult is an Contact result from xero
type ContactResult struct {
	ID           string    `json:"ID"`
	ProviderName string    `json:"ProviderName"`
	Status       string    `json:"Status"`
	Contacts     []Contact `json:"Contacts"`
}

//Contact is an Contact from xero
type Contact struct {
	ContactID  string `json:"ContactID,omitempty" xml:"ContactID,omitempty"`
	Name       string `json:"Name,omitempty" xml:"Name,omitempty"`
	FirstName  string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName   string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	Email      string `json:"Email,omitempty" xml:"Email,omitempty"`
	IsSupplier bool   `json:"IsSupplier,omitempty" xml:"IsSupplier,omitempty"`
}
