package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	var mux = http.NewServeMux() // Это специальный хендлер, который перенаправляет запросы в дригие хэндлеры.
	// Таким образом происходит цепное подключение хэндлеров.
	// Это позволяет нам более удобнную маршрутизацию.

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// log.Println(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/snippet/delete", app.snippetDelete)

	return mux
}
