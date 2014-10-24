package main

import (
	"appengine"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	g "github.com/jeisenberg/goa"
	"log"
	"net/http"
)

type User struct {
	Name string
	Id   string
}

func (u *User) BeforeSave() {
	log.Println("test before save")
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/test", createUser)
	r.HandleFunc("/users/{id}", getUser)

	n := negroni.Classic()
	n.UseHandler(r)
	http.Handle("/", n)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("called")
	c := appengine.NewContext(r)
	user := &User{
		Name: "test",
	}
	err := g.Save(c, user)
	if err != nil {
		log.Printf("Errror: %s", err)
		w.Write([]byte("not ok"))
		return
	}
	user.Name = "testyyy"
	err = g.Update(c, user)
	log.Println(user)
	w.Write([]byte(user.Id))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	g.Get(c, id, &user)
	w.Write([]byte(user.Name))
}
