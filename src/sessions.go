package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"errors"
	"log"
)

var (
	store = sessions.NewCookieStore([]byte("super-secret-key"))
)

func ValidateSession(w http.ResponseWriter, r *http.Request) (id int, err error){
	session, _ := store.Get(r, "jwt")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		err = errors.New("Not authenticated")
		return
	}

	log.Println("coisas")
	log.Println(string(session.Values["jwt"].([]byte)))

	id, err = ValidateJWT(session.Values["jwt"].([]byte))
	if err != nil {
        return
	}

	return id, err
}

func LoginSession(w http.ResponseWriter, r *http.Request, id int,  name string) {
	session, _ := store.Get(r, "jwt")

	session.Values["authenticated"] = true
	session.Values["jwt"] = GenerateJWT(id, name)
	session.Save(r, w)
}

func LogoutSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "jwt")

	session.Values["authenticated"] = false
	session.Values["jwt"] = ""
	session.Save(r, w)
}
