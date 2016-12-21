package goxerostructs

//TimesheetResult is an Timesheet result from xero
type TimesheetResult struct {
	ID           string  `json:"ID"`
	ProviderName string  `json:"ProviderName"`
	Status       string  `json:"Status"`
	Timesheets   PayItem `json:"Timesheets"`
}

//Timesheet is an timesheet within xero for payroll
type Timesheet struct {
	EmployeeID string `json:"EmployeeID"`
	StartDate  string `json:"StartDate"`
	EndDate    string `json:"EndDate"`
	Status     string `json:"Status"`
}
