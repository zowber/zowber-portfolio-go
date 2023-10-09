package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(http.StatusNotFound)
	temp := template.Must(template.ParseFiles("./templates/404.html"))
	temp.Execute(w, r.URL.Path)
}

func main() {

	db, err := newPortfolioDbClient()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFoundHandler(w, r, http.StatusNotFound)
			return
		}

		temp := template.Must(template.ParseFiles("./templates/index.html"))
		data, err := db.getAllCaseStudies()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, data)
	})

	http.HandleFunc("/case-study/", func(w http.ResponseWriter, r *http.Request) {

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) > 3 {
			notFoundHandler(w, r, http.StatusNotFound)
			return
		}

		caseStudyIdStr := parts[2]

		caseStudyId, err := strconv.Atoi(caseStudyIdStr)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := db.getCaseStudy(caseStudyId)
		if err != nil {
			log.Println(err)
			notFoundHandler(w, r, http.StatusNotFound)
			return
		}

		temp := template.Must(template.ParseFiles("./templates/case-study.html"))
		temp.Execute(w, data)
	})

	log.Fatal(http.ListenAndServeTLS(":8080", "server.pem", "server.key", nil))

}
