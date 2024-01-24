package main

import (
	"fmt"
	"net/http"

	"acm_web_page/views/user"

	"github.com/a-h/templ"
)

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

var ACM_boards = []ACM_board{
	{
		Chair: ACM_board_member{
			Person: Person{
				Name:           "Jean Paul",
				Last_names:     "Ari Maldonado",
				Personal_email: "_",
			},
			UCSP_email: "_",
		},
		Vice_chair: ACM_board_member{
			Person: Person{
				Name:           "Andrea",
				Last_names:     "Chávez Aguilar",
				Personal_email: "_",
			},
			UCSP_email: "_",
		},
		Treasurer: ACM_board_member{
			Person: Person{
				Name:           "Gabriel",
				Last_names:     "Hugo Manchego",
				Personal_email: "_",
			},
			UCSP_email: "_",
		},
	},
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the GET request to "/search" here
	query := r.URL.Query().Get("search")
	if query != "" {
		// templ.Handler(component)
		responseHTML := fmt.Sprintf("<div>%s</div>", query)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(responseHTML))
	} else {
		fmt.Fprint(w, "Please provide a search query")
	}
}

func main() {
	component := user.Hello("ma naim ist john")
	// fmt.Sprintf(component)

	// http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/search", searchHandler)

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
