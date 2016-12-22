package goxerostructs

//TimesheetResult is an Timesheet result from xero
type TimesheetResult struct {
	ID           string      `json:"ID"`
	ProviderName string      `json:"ProviderName"`
	Status       string      `json:"Status"`
	Timesheets   []Timesheet `json:"Timesheets"`
}

type Timesheets []Timesheet

//Timesheet is an timesheet within xero for payroll
type Timesheet struct {
	EmployeeID     string         `json:"EmployeeID"`
	StartDate      string         `json:"StartDate"`
	EndDate        string         `json:"EndDate"`
	Status         string         `json:"Status"`
	TimesheetLines TimesheetLines `json:"TimesheetLines"`
}

type TimesheetLines struct {
	TimesheetLine []TimesheetLine `json:"TimesheetLine"`
}

//TimesheetLine is an timesheet within xero for payroll
type TimesheetLine struct {
	EarningsRateID string        `json:"EarningsRateID"`
	NumberOfUnits  NumberOfUnits `xml:"NumberOfUnits"`
}

//NumberOfUnits is an timesheet within xero for payroll
type NumberOfUnits struct {
	NumberOfUnits []NumberOfUnit `xml:"NumberOfUnit"`
}

type NumberOfUnit struct {
	Content string `xml:",chardata"`
}

/*
//Timesheet is an timesheet within xero for payroll
type Timesheet struct {
	EmployeeID     string          `json:"EmployeeID"`
	StartDate      string          `json:"StartDate"`
	EndDate        string          `json:"EndDate"`
	Status         string          `json:"Status"`
	TimesheetLines []TimesheetLine `json:"TimesheetLines"`
}

//TimesheetLine is an timesheet within xero for payroll
type TimesheetLine struct {
	TimesheetValue TimesheetValue `xml:"TimesheetLine"`
}

//TimesheetLine is an timesheet within xero for payroll
type TimesheetValue struct {
	EarningsRateID string        `json:"EarningsRateID"`
	NumberOfUnits  NumberOfUnits `xml:"NumberOfUnits"`
}

//NumberOfUnits is an timesheet within xero for payroll
type NumberOfUnits struct {
	NumberOfUnits []NumberOfUnit `xml:"NumberOfUnit"`
}

type NumberOfUnit struct {
	Content string `xml:",chardata"`
}
*/
