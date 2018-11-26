package main
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DataBase *sql.DB
var connect string = "root:pass@tcp(db:3306)/db?charset=utf8mb4,utf8&parseTime=True"

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


