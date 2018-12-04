package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"errors"
)

var (
	store = sessions.NewCookieStore([]byte("super-secret-key"))
)

func ValidateSession(r *http.Request) (id int, err error){
	session, err := store.Get(r, "jwt")
	if err != nil {
		panic(err)
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		err = errors.New("Not authenticated")
		return
	}

	id, err = ValidateJWT([]byte(session.Values["jwt"].(string)))

	if err != nil {
        return
	}
	return id, err
}

func LoginSession(w http.ResponseWriter, r *http.Request, id int,  name string) {
	session, _ := store.Get(r, "jwt")
	session.Values["authenticated"] = true
	session.Values["jwt"] = string(GenerateJWT(id, name))
	err := session.Save(r, w)
	if err != nil {
		panic(err)
	}
}

func LogoutSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "jwt")
	if err != nil {
		panic(err)
	}
	session.Values["authenticated"] = false
	session.Values["jwt"] = ""
	session.Options.MaxAge = -1
	err = session.Save(r, w)
}
