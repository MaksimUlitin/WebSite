package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError записывает сообщение об ошибке в errorLog и
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// омощник clientError отправляет "BadRequest", когда есть проблема с пользовательским запросом
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// помощник для clientError, которая отправляет пользователю ответ "404 Страница не найдена"
func (app *application) NotFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
