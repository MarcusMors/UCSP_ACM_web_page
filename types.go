package main

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
	Name             string `json:"name"`
	Planning_period  string `json:"planning_period"`
	Execution_period string `json:"execution_period"`
	Results_period   string `json:"results_period"`
	Results_data_id  string `json:"results_data_id"`
}
