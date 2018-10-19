package web

import (
	"log"
	"net/http"

	"shortme/conf"
	"shortme/web/api"
	"shortme/web/www"

	"github.com/gorilla/mux"
)

// Start is
func Start() {
	log.Println("web starts")
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	r.StrictSlash(true)
	// r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/version", api.CheckVersion).Methods(http.MethodGet)
	r.HandleFunc("/health", api.CheckHealth).Methods(http.MethodGet)
	r.HandleFunc("/short", api.ShortURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/expand", api.ExpandURL).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/{shortenedURL:[a-zA-Z0-9]{1,11}}", api.Redirect).Methods(http.MethodGet)

	r.HandleFunc("/index.html", www.Index).Methods(http.MethodGet)

	r.Handle("/static/{type}/{file}", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Handle("/favicon.ico", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(conf.Conf.Http.Listen, r))
}
