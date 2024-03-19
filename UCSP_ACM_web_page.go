package main

import (
	"UCSP_ACM_web_page/views/user"
	"fmt"
	"net/http"
	"os"
)

var Events = []Event{
	{
		Name:             "Concurso Escolar de Programación",
		Planning_period:  "a",
		Execution_period: "a",
		Results_period:   "a",
		Results_data_id:  "a",
	},
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

// func searchHandler(w http.ResponseWriter, r *http.Request) {
//   Hello("Hello world").Render(r.Context(),w)
// }

func main() {

	staticDir := "/static/img/"

	fileServer := http.FileServer(http.Dir(staticDir))

	http.Handle("/static/img/", http.StripPrefix("/static/img/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Index_main().Render(r.Context(), w)
		// Index(search_result).Render(r.Context(), w)
		if err := user.Layout("UCSP ACM", user.Index_main()).Render(r.Context(), w); err != nil {
			panic("Error creating Layout" + ": " + err.Error())
		}
	})

	// index := Layout(Index_main())
	// http.HandleFunc("/", templ.Handler(index))

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n", ACM_boards[0].Chair.Person.Name)
		searchResult := user.Search(ACM_boards[0].Chair.Person.Name)
		if err := searchResult.Render(r.Context(), w); err != nil {
			panic("Error searching result: " + err.Error())
		}
	})

	http.HandleFunc("/static/img/ucsp_2021_cropped.png", func(w http.ResponseWriter, r *http.Request) {
		imagePath := "static/img/ucsp_2021_cropped.png"

		if _, err := os.Stat(imagePath); err == nil {
			http.ServeFile(w, r, imagePath)
		} else {
			http.Error(w, "Image not found", http.StatusNotFound)
		}
	})

	http.HandleFunc("/static/styles/normalize.css", func(w http.ResponseWriter, r *http.Request) {
		stylePath := "static/styles/normalize.css"

		if _, err := os.Stat(stylePath); err == nil {
			http.ServeFile(w, r, stylePath)
		} else {
			fmt.Println("NOT FOUND NORMALIZE")
			http.Error(w, "Image not found", http.StatusNotFound)
		}
	})

	http.HandleFunc("/static/styles/layout.css", func(w http.ResponseWriter, r *http.Request) {
		stylePath := "static/styles/layout.css"

		if _, err := os.Stat(stylePath); err == nil {
			http.ServeFile(w, r, stylePath)
		} else {
			fmt.Println("NOT FOUND LAYOUT")
			http.Error(w, "Image not found", http.StatusNotFound)
		}
	})

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	Hello("Hello world").Render(r.Context(), w)
	// })

	port := "3000"
	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic("Error starting the server: " + err.Error())
	}

}
