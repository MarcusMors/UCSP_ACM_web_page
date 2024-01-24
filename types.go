package types

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Person struct {
	Name           string `json:"name"`
	Last_names     string `json:"last_names"`
	Personal_email string `json:"personal_email"`
}

type ACM_board_member struct {
	Person
	UCSP_email string `json:"ucsp_email"`
}

type ACM_board struct {
	Chair      ACM_board_member
	Vice_chair ACM_board_member
	Treasurer  ACM_board_member
	Period
}

type Event struct {
	Name                  string `json:"name"`
	planning_period_date  string `json:"planning_period_date"`
	execution_period_date string `json:"execution_period_date"`
	results_period_date   string `json:"results_period_date"`
	results_data_id       string `json:"results_data_id"`
}
