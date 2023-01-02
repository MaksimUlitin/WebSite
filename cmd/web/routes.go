package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// запись в обработчик
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/pageBox", app.displayNotes)
	mux.HandleFunc("/pageBox/formNotes", app.formNotes)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	//fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	//mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
