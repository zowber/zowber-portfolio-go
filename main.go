package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	db, err := newPortfolioDbClient()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/index.html"))
		data, err := db.getAllCaseStudies()
		if err != nil {
			log.Print("Error getting case studies from DB")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, data)
	})

	http.HandleFunc("/case-study/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/case-study.html"))

		parts := strings.Split(r.URL.Path, "/")
		caseStudyIdStr := parts[2]

		caseStudyId, err := strconv.Atoi(caseStudyIdStr)
		if err != nil {
			log.Print("Invalid query param from URL")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := db.getCaseStudy(caseStudyId)
		if err != nil {
			log.Print("Error getting case studies from DB")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, data)
	})

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/image.html"))

		path := r.URL.Query().Get("path")

		log.Print(path)

		temp.Execute(w, path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
