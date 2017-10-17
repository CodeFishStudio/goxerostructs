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
	Name               string           `json:"Name"`
	TrackingCategoryID string           `json:"TrackingCategoryID"`
	Options            []TrackingOption `json:"Options"`
	Option             string           `json:"Option"`
}

//TrackingOption is an Tracking from xero
type TrackingOption struct {
	TrackingOptionID string `json:"TrackingOptionID"`
	Name             string `json:"Name"`
	Status           string `json:"Status"`
	IsActive         bool   `json:"IsActive"`
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
