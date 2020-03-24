package routes

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() *mux.Router {
	var entry string
	var static string

	flag.StringVar(&static, "static", "./web/dist", "the directory to serve static files from.")
	flag.StringVar(&entry, "entry", "./web/dist/index.html", "the entry point to serve.")
	flag.Parse()

	r := mux.NewRouter()


	// Routes consist of a path and a handler function.
	staticAssetsHandler := http.FileServer(http.Dir(static))
	r.PathPrefix("/dist").Handler(http.StripPrefix("/dist", staticAssetsHandler))


	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler(entry))

	return r
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}