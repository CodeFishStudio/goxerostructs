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
	ContactID  string `json:"ContactID"`
	Name       string `json:"Name"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Email      string `json:"Email"`
	IsSupplier bool   `json:"IsSupplier"`
}
