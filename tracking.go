package goxerostructs

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
