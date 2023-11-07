package main

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

var t *Template

type Template struct {
	templates *template.Template
}

func (t *Template) render(w io.Writer, name string, data any) error{
	return t.templates.ExecuteTemplate(w, name, data)
}

type User struct{
	Username string
	IsUser 	 bool
	Age      int
}

func main(){
	t = &Template{
		templates: template.Must(template.ParseGlob("www/*.html")),
	}
	http.HandleFunc("/", makeAPIFunc(handleHome))
	http.HandleFunc("/api/user", makeAPIFunc(handleUser))
	http.ListenAndServe(": 3000", nil)

}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request){
		if err := fn(w, r); err != nil{
			writeJson(w, http.StatusInternalServerError, map[string] string{"error": err.Error()})
		}
	}
}

func handleUser( w http.ResponseWriter, r *http.Request) error{
	return writeJson(w, http.StatusOK, map[string] string{"message":"hello there, Banki"})
}

func handleHome(w http.ResponseWriter, r *http.Request) error{
	user := User{
		Username: "Dio",
		IsUser: true,
		Age: 137,
	}
	return t.render(w, "index.html", user)
}

func writeJson(w http.ResponseWriter, code int, v any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}