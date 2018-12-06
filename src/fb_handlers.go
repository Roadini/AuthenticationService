package main

import (
    "net/http"
    "encoding/json"
    // "log"
    fb "github.com/huandu/facebook"
    "fmt"
    "database/sql"
)

var LoginFb = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    decoder := json.NewDecoder(r.Body)
    var data FB_Login
    err := decoder.Decode(&data)

    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    res, _ := fb.Get("/me", fb.Params{
        "fields": "id,name,age_range,email,gender",
        "access_token": data.AccessToken,
    })

    res["email"]="luis@ua.pt"

    fmt.Println("Here is my Facebook first name:", res["name"])
    fmt.Println("Here is my Facebook first name:", res["id"])
    fmt.Println("Here is my Facebook first name:", res["age_range"])
    fmt.Println("Here is my Facebook first name:", res["email"])
    fmt.Println("Here is my Facebook first name:", res["gender"])

    var i int
    fmt.Sscan(res["id"].(string), &i)

    err = InsertUserFB(i, res["name"].(string), res["email"].(string), data.AccessToken)
    if err != nil {
        http.Error(w, `"code": "Error inserting user"` , 400)
        return
    }

    LoginSession(w, r, i, res["name"].(string))

    w.Write([]byte(`{"code": "success"}`))
})

func InsertUserFB(id int, name string, email string, accessToken string) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    insertUser, err := DataBase.Prepare("INSERT INTO user_details ( id, name, email, accessToken) VALUES (?, ?, ?, ?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec( id, name, email, accessToken)
    if err != nil {
        panic(err.Error())
    }

    return 
}