package goxerostructs

//TrackingResult is an Tracking result from xero
type TrackingResult struct {
	ID           string     `json:"ID"`
	ProviderName string     `json:"ProviderName"`
	Status       string     `json:"Status"`
	Trackings    []Tracking `json:"TrackingCategories"`
}

//Tracking is an Tracking from xero
type Tracking struct {
	Name               string           `json:"Name,omitempty" xml:"Name,omitempty"`
	TrackingCategoryID string           `json:"TrackingCategoryID,omitempty" xml:"TrackingCategoryID,omitempty"`
	Options            []TrackingOption `json:"Options,omitempty" xml:"Options,omitempty"`
	Option             string           `json:"Option,omitempty" xml:"Option,omitempty"`
}

//TrackingOption is an Tracking from xero
type TrackingOption struct {
	TrackingOptionID string `json:"TrackingOptionID,omitempty" xml:"TrackingOptionID,omitempty"`
	Name             string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status           string `json:"Status,omitempty" xml:"Status,omitempty"`
	IsActive         bool   `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
}

/*
"Tracking": [
	{
		"Name": "Region",
		"Option": "Eastside",
		"TrackingCategoryID": "093af706-c2aa-4d97-a4ce-2d205a017eac",
		"Options": []
	}
],
*/
