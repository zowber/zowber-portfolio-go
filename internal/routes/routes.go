package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/zowber/zowber-portfolio-go/internal/data"
)

var db, err = data.NewPortfolioDbClient()

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/case-study/", caseStudyHandler)

	return mux
}

var errorHandler = func(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(http.StatusNotFound)
	temp := template.Must(template.ParseFiles("./templates/error.html"))
	temp.Execute(w, r.URL.Path)
}

var rootHandler = func(w http.ResponseWriter, r *http.Request) {

	log.Println("Req:", r.URL.Path)
	// 404 on anything but root
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	temp := template.Must(template.ParseFiles("./templates/index.html"))
	data, err := db.GetAllCaseStudies()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

var caseStudyHandler = func(w http.ResponseWriter, r *http.Request) {

	log.Println("Req:", r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) > 3 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	caseStudyIdStr := parts[2]

	caseStudyId, err := strconv.Atoi(caseStudyIdStr)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := db.GetCaseStudy(caseStudyId)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	temp := template.Must(template.ParseFiles("./templates/case-study.html"))
	temp.Execute(w, data)
}
