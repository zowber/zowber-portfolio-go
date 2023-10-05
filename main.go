package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/index.html"))
		data, err := getAllCaseStudies()
		if err != nil {
			fmt.Printf("Error getting case studies")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		temp.Execute(w, data)
	})

	http.HandleFunc("/case-study/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/case-study.html"))

		parts := strings.Split(r.URL.Path, "/")
		caseStudyIdStr := parts[len(parts)-1]

		caseStudyId, err := strconv.Atoi(caseStudyIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := getCaseStudy(caseStudyId)
		if err != nil {
			fmt.Printf("Error getting case study")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		temp.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
