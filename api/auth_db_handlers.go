package main

import (
    "encoding/json"
    "math/rand"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
    "reflect"
    "crypto/sha256"
)

var DataBase *sql.DB
var connect string = "root:pass@tcp(172.17.0.3:3306)/db?charset=utf8mb4,utf8&parseTime=True"

func PingDB(){
    // Open doesn't open a connection. Validate DSN data:

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer DataBase.Close()

    err = DataBase.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    log.Printf("Pinged the database succefully\n")
}


func InsertUser(user *User) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer DataBase.Close()

    salt := make([]byte, 8)
    rand.Read(salt)
    user.Salt = salt

    user.Hash = sha256.Sum256(append([]byte(user.Pass), salt...))

    pr , _ := json.Marshal(user)
    log.Println(string(pr))


    insertUser, err := DataBase.Prepare("INSERT INTO user_details ( age , email, name, gender, salt, hash) VALUES (?, ?, ?, ?, ?, ?)") // ? = placeholder
    if err != nil {
    	log.Println(1)
        panic(err.Error()) // proper error handling instead of panic in your app
    	log.Println(2)
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec( user.Age , user.Email, user.Name, user.Gender, string(user.Salt), string(user.Hash[:]))
    return err
}

func GetUsers(getBy string, value interface {}) ([]UserToOutside){

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer DataBase.Close()

    var query string

    if valueString, ok := value.(string); ok {
        value = valueString
    } else if valueInt, ok := value.(int); ok {
        value = valueInt
    } else if valueInt, ok := value.(float64); ok {
        value = valueInt
    } else{
        fmt.Println(reflect.TypeOf(value), " ", value)
        panic("Error")
    }

    log.Println("teste")
    log.Println(getBy)

    if (getBy == "user_id") {
        query ="SELECT user_id, age, email, name, gender FROM user_details WHERE user_id = ?"
    } else if (getBy == "age") {
        query ="SELECT user_id, age, email, name, gender FROM user_details WHERE age = ?"
    } else if (getBy == "email") {
        query ="SELECT user_id, age, email, name, gender FROM user_details WHERE email = ?"
    } else if (getBy == "name") {
        query ="SELECT user_id, age, email, name, gender FROM user_details WHERE name = ?"
    } else if (getBy == "gender") {
        query ="SELECT user_id, age, email, name, gender FROM user_details WHERE gender = ?"
    } else {
        panic( fmt.Sprintf("Bad Info. getBy = (%s)", getBy) )
    }

    rows, err := DataBase.Query(query, value)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    var allu []UserToOutside
    // Get column names
    for rows.Next(){

        var u UserToOutside
        if err := rows.Scan(&u.Id, &u.Age, &u.Email, &u.Name, &u.Gender ); err != nil {
            log.Fatal(err)
        }
        allu = append(allu, u)
    }
    
    return allu
}
