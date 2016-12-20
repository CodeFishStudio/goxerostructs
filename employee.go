package goxerostructs

//EmployeeResult is an employee result from xero
type EmployeeResult struct {
	ID           string     `json:"ID"`
	ProviderName string     `json:"ProviderName"`
	Status       string     `json:"Status"`
	Employees    []Employee `json:"Employees"`
}

//Employee is an employee from xero
type Employee struct {
	EmployeeID             string      `json:"EmployeeID"`
	FirstName              string      `json:"FirstName"`
	LastName               string      `json:"LastName"`
	Status                 string      `json:"Status"`
	Email                  string      `json:"Email"`
	DateOfBirth            string      `json:"DateOfBirth"`
	Gender                 string      `json:"Gender"`
	Phone                  string      `json:"Phone"`
	Mobile                 string      `json:"Mobile"`
	OrdinaryEarningsRateID string      `json:"OrdinaryEarningsRateID"`
	PayTemplate            PayTemplate `json:"PayTemplate"`
}

//PayTemplate is an pay template from xero
type PayTemplate struct {
	EarningsLines []EarningsLine `json:"EarningsLines"`
	LeaveLines    []LeaveLine    `json:"LeaveLines"`
}

//EarningsLine is an earning line from xero
type EarningsLine struct {
	EarningsRateID       string  `json:"EarningsRateID"`
	CalculationType      string  `json:"CalculationType"`
	AnnualSalary         float64 `json:"AnnualSalary"`
	RatePerUnit          float64 `json:"RatePerUnit"`
	NumberOfUnits        float64 `json:"NormalNumberOfUnits"`
	NumberOfUnitsPerWeek float64 `json:"NumberOfUnitsPerWeek"`
}

//LeaveLine is a leave line from xero
type LeaveLine struct {
	LeaveTypeID                    string  `json:"LeaveTypeID"`
	CalculationType                string  `json:"CalculationType"`
	AnnualNumberOfUnits            float64 `json:"AnnualNumberOfUnits"`
	FullTimeNumberOfUnitsPerPeriod float64 `json:"FullTimeNumberOfUnitsPerPeriod"`
	EntitlementFinalPayPayoutType  string  `json:"EntitlementFinalPayPayoutType"`
}
