package main

import (
    "net/http"
    "encoding/json"
    mysql "github.com/go-sql-driver/mysql"
    "log"
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
        http.Error(w, `"code": "Invalid Format"` , 400)
        return
    }
    log.Println(string(x))

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if( user.Pass == "" ||
    user.Email == "" ||
    user.Name == ""){
        http.Error(w, `"code": "Invalid info"` , 400)
        return
    }
    
    err = InsertUser(&user)
    if err != nil && err.(*mysql.MySQLError).Number == 1062 {
        http.Error(w, `{"code": "Duplicate email"}`, 400)
        return
    }
    
    w.Write([]byte(`{"code": "success"}`))
})

var GetUsersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
    _, err := ValidateSession(w, r)
    if err!= nil{
        http.Error(w, `{"code": "Not Logged In or Invalid session. Please Relog"}`, 400)
        return
    }
    
    decoder := json.NewDecoder(r.Body)

    var j struct {
        GetBy string
        Value interface{}
    }

    err = decoder.Decode(&j)
    if err != nil {
        panic(err)
    }
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    users, err := GetUsers(j.GetBy, j.Value);
    if  err!= nil{
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    js, err := json.Marshal(users)
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }

    w.Write([]byte(js))
})

var LoginUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    decoder := json.NewDecoder(r.Body)

    var login_dets struct {
        Email, Pass string
    }

    err := decoder.Decode(&login_dets)
    if err != nil {
        http.Error(w, `{"code": "Invalid Json structure"}`, 400)
        return
    }

    if err := CheckUserPassDB(login_dets.Email, login_dets.Pass); err!= nil{
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    users, _:= GetUsers("email", login_dets.Email)
    LoginSession(w, r, users[0].Id, users[0].Name)

    w.Write([]byte(`{"code": "success"}`))

})

var LogoutUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

    _, err := ValidateSession(w, r)
    if err!= nil{
        http.Error(w, `{"code": "Invalid Cookie. No need to logout. Relog"}`, 400)
        return
    }
    LogoutSession(w, r)
    
    w.Write([]byte(`{"code": "success"}`))

})


var DeleteUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

    id, err := ValidateSession(w, r)
    if err!= nil{
        http.Error(w, `{"code": "Please relog"}`, 400)
        return
    }

    LogoutSession(w, r)
    DeleteUserDB(id)
    w.Write([]byte(`{"code": "success"}`))
})

var UpdateUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    
    /*
    
    id, err := ValidateSession(w, r)
    if err!= nil{
        http.Error(w, `{"code": "Please relog"}`, 400)
        return
    }
	*/
    decoder := json.NewDecoder(r.Body)
    var u User
    err := decoder.Decode(&u)
    if err!= nil{
        http.Error(w, `{"code": "Invalid Json Format"}`, 400)
        return
    }
    
    var objmap map[string]*json.RawMessage
    foo_marshalled, _ := json.Marshal(u)
    
    log.Println(string(foo_marshalled))
    
    
    err = json.Unmarshal(foo_marshalled, &objmap)
    for key, val := range objmap {

	    log.Println("Key: " + key + "; Value: ")
	    log.Println(val)
    }
    log.Println(objmap)
    /*
    
    
    
    b, err = json.Marshal(u)
	

    UpdateUserDB(id, u)
    */
})