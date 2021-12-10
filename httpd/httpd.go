package httpd

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mister-turtle/golang-webapp-structure/evidence"
)

//go:embed embedded/*.gohtml
var templateFS embed.FS

//go:embed embedded/*.css
var staticFS embed.FS

type Server struct {
	address    string
	router     chi.Router
	templates  *template.Template
	iocService iocService
}

func (s Server) GetIOCs(w http.ResponseWriter, r *http.Request) {

	iocs, err := s.iocService.FindAll(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("error finding iocs: %s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", iocs)
	err = s.templates.ExecuteTemplate(w, "index.gohtml", iocs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("error executing index template: %s\n", err.Error())
		return
	}
}

func (s Server) NewIOC(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("error executing parsing IOC form: %s\n", err.Error())
		return
	}

	iocDate := r.Form.Get("iocDate")
	iocType := r.Form.Get("iocType")
	iocValue := r.Form.Get("iocValue")
	iocSource := r.Form.Get("iocSource")

	if iocDate == "" || iocType == "" || iocValue == "" || iocSource == "" {
		http.Error(w, "missing ioc data", http.StatusBadRequest)
		return
	}

	jsDate, err := time.Parse("2006-01-02T15:04:05.999Z07:00", iocDate)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		log.Printf("error converting date format: %s\n", err.Error())
		return
	}

	newIOC := evidence.NewIOC(iocType, iocValue, jsDate, iocSource)
	err = s.iocService.Create(r.Context(), newIOC)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("error creating ioc: %s\n", err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s Server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}

func NewServer(addr string, iocService iocService) (Server, error) {

	newServer := Server{
		address:    addr,
		iocService: iocService,
	}

	templates, err := template.ParseFS(templateFS, "embedded/*.gohtml")
	if err != nil {
		return newServer, err
	}

	newServer.templates = templates

	staticSubFS, err := fs.Sub(staticFS, "embedded")
	if err != nil {
		return newServer, err
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", newServer.GetIOCs)
	router.Post("/", newServer.NewIOC)
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticSubFS))))

	newServer.router = router
	return newServer, nil
}
