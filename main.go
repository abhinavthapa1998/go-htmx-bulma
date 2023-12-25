package main

import (
  "html/template"
  "log"
  "net/http"
)

type Film struct {
  Title string
  Director string
}

func main(){
  
  getHome := func (w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("index.html"))
    films := map[string][]Film{
      "Films":{
        {Title: "The Godfather", Director: "Francis Ford Coppola"},
        {Title: "Bladerunner", Director: "Ridley Scott"},
        {Title: "The Thing", Director: "John Carpenter"},
      },
    }
    tmpl.Execute(w, films)
  }

  createFilm := func (w http.ResponseWriter, r *http.Request){
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")  
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
  }

  
  http.HandleFunc("/", getHome)
  http.HandleFunc("/add-film/", createFilm)
  log.Fatal(http.ListenAndServe(":8080",nil))
}
