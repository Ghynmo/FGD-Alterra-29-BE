package request

import reportcases "fgd-alterra-29/business/report_cases"

type AddCase struct {
	Case string `form:"case" json:"case"`
}

func (ac *AddCase) ToDomain() reportcases.Domain {
	return reportcases.Domain{
		Case: ac.Case,
	}
}
