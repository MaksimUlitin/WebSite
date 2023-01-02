package main

import (
	"errors"
	"fmt"
	"github.com/maksimUlitin/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "история про улитку "
	content := "улитка выползла из раковинв,\n вытянула рожки \n и спрятала из обратно "
	expires := "7"

	id, err := app.pageBox.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/pageBox?id=%d", id), http.StatusSeeOther)

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.portal.tmpl",
	}

	tmp, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err) //http.Error(w, "внутренняя ошибка на сервере", 500)
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err) //http.Error(w, "внутренняя ошибка на сервере", 500)
	}

	w.Write([]byte(""))
}

// отображение заметок
func (app *application) displayNotes(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w) //http.NotFound(w, r)
		return
	}
	s, err := app.pageBox.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorOnRecord) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%v", s)
	//fmt.Fprintf(w, "отображение заметки под id %d", id)

}

// форма создания заметок
func (app *application) formNotes(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) //http.Error(w, "get-пост запрещен!", 405)
		return
	}
	w.Write([]byte("форма для создания заметок"))
}
