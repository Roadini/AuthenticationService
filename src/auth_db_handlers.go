package main

import (
    "math/rand"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
    "reflect"
    "crypto/sha256"
    "errors"
    "bytes"
    "github.com/fatih/structs"
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

    insertUser, err := DataBase.Prepare("INSERT INTO user_details ( age , email, name, gender, salt, hash) VALUES (?, ?, ?, ?, ?, ?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec( user.Age , user.Email, user.Name, user.Gender, string(user.Salt), string(user.Hash[:]))
    return err
}

func GetUsers(getBy string, value interface {}) (user_list []UserToOutside, err error){

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
        err = errors.New("Invalid provided info")
    }

    if (getBy == "id") {
        query ="SELECT id, age, email, name, gender FROM user_details WHERE id = ?"
    } else if (getBy == "age") {
        query ="SELECT id, age, email, name, gender FROM user_details WHERE age = ?"
    } else if (getBy == "email") {
        query ="SELECT id, age, email, name, gender FROM user_details WHERE email = ?"
    } else if (getBy == "name") {
        query ="SELECT id, age, email, name, gender FROM user_details WHERE name = ?"
    } else if (getBy == "gender") {
        query ="SELECT id, age, email, name, gender FROM user_details WHERE gender = ?"
    } else {
        err = errors.New("Invalid provided info")
    }

    rows, err := DataBase.Query(query, value)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    //var user_list []UserToOutside
    // Get column names
    for rows.Next(){
        var u UserToOutside
        if err := rows.Scan(&u.Id, &u.Age, &u.Email, &u.Name, &u.Gender ); err != nil {
            log.Fatal(err)
        }
        user_list = append(user_list, u)
    }

    return user_list, err
}

func CheckUserPassDB(email string, pass string) (err error){

    err = nil

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        return
    }

    defer DataBase.Close()
    query := "SELECT id, email, salt, hash FROM user_details WHERE email = ?"
    row:= DataBase.QueryRow(query, email)


    var data struct{
        Id int
        Email string
        Salt string
        Hash string
    }

    err = row.Scan(&data.Id, &data.Email, &data.Salt, &data.Hash)
    if err != nil {
        err = errors.New("No user found")
    }

    log.Println(data.Id)
    log.Println(data.Email)
    log.Println(data.Salt)
    log.Println(data.Hash)


    var tmphash [32] byte = sha256.Sum256(append([]byte(pass), data.Salt...))
    
    // log.Println(tmphash[:])
    // log.Println([]byte(data.Hash))

    if !bytes.Equal([]byte(data.Hash), tmphash[:]) {
        err = errors.New("No user found (Pass error)")
    }
    log.Println([]byte(data.Hash))
    log.Println(sha256.Sum256(append([]byte(pass), data.Salt...)))


    return err
}

func DeleteUserDB(id int) (err error){

    err = nil

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        return
    }
    defer DataBase.Close()
    query := "DELETE FROM user_details WHERE id = ?;"
    _, err = DataBase.Exec(query, id)
    if err != nil{
        errors.New("Could not delet the user")
    }

    return err
}

func UpdateUserDB(id int, u User) (err error){

    err = nil

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        return
    }
    defer DataBase.Close()

    user := structs.Map(u)

    for key, val := range user {

        log.Println("Key: " + key + "; Value: " + val.(string) )

        /*
        switch concreteVal := val.(type) {
        case map[string]interface{}:
            fmt.Println(key)
            parseMap(val.(map[string]interface{}))
        case []interface{}:
            fmt.Println(key)
            parseArray(val.([]interface{}))
        default:
            fmt.Println(key, ":", concreteVal)
        }

        */
    }

    /*
    query := "UPDATE table_name SET  WHERE id = ?;"
    result, err := DataBase.Exec(query, id)

    if err != nil{
        errors.New("Could not delet the user")
    }
    */
    return err
}
