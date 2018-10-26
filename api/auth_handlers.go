package main

import (
    "net/http"
    "encoding/json"
    mysql "github.com/go-sql-driver/mysql"
    "log"
    "github.com/gorilla/sessions"
)

var CreateUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

    decoder := json.NewDecoder(r.Body)

    var user User
    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }

    x , _ := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    log.Println(string(x))

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if( user.Pass == "" ||
    user.Email == "" ||
    user.Name == ""){
    	// js, _ := json.RawMessage(`"error": "Invalid info"`)
        http.Error(w, `"code": "Invalid info"` , 400)
        return
    }
    
    err = InsertUser(&user)
    if err != nil && err.(*mysql.MySQLError).Number == 1062 {
    	//js, _ := json.RawMessage(`"error": "Duplicate email"`)
    	//js, _ := json.Marshal([]byte(`{"error": "Duplicate email"}`))
        http.Error(w, `{"code": "Duplicate email"}`, 400)
        return
    }
    
    //js, _ := json.Marshal(`teste`)
    w.Write([]byte(`{"code": "success"}`))
})

var GetUsersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
    decoder := json.NewDecoder(r.Body)

    var j struct {
        GetBy string
        Value interface{}
    }

    err := decoder.Decode(&j)
    if err != nil {
        panic(err)
    }

	/*
    x , _ := json.Marshal(j)
    if err != nil {
        panic(err)
    }
    log.Println(string(x))
	*/
	
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    js, err := json.Marshal(GetUsers(j.GetBy, j.Value))
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }
    log.Printf("%s", js)

    w.Write([]byte(js))
})



var DeleteUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
    decoder := json.NewDecoder(r.Body)

    var j struct {
        GetBy string
        Value interface{}
    }

    err := decoder.Decode(&j)
    if err != nil {
        panic(err)
    }

	/*
    x , _ := json.Marshal(j)
    if err != nil {
        panic(err)
    }
    log.Println(string(x))
	*/
	
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    js, err := json.Marshal(GetUsers(j.GetBy, j.Value))
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }
    log.Printf("%s", js)

    w.Write([]byte(js))
})