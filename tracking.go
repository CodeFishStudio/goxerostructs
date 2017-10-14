package goxerostructs

//TrackingResult is an Tracking result from xero
type TrackingResult struct {
	ID           string     `json:"ID"`
	ProviderName string     `json:"ProviderName"`
	Status       string     `json:"Status"`
	Trackings    []Tracking `json:"Trackings"`
}

//Tracking is an Tracking from xero
type Tracking struct {
	Name               string `json:"Name"`
	Option             string `json:"Option"`
	TrackingCategoryID string `json:"TrackingCategoryID"`
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
