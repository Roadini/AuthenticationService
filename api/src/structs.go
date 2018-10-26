package main

/*
import (
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "io/ioutil"
    "net/http/httputil"
    "log"
)
*/
type User struct {
    UserID int
    Age int
    Email string
    Name string
    Gender string
    Pass string
    Hash [32]byte
    Salt []byte
}

type UserToOutside struct {
    Id , Age int
    Email, Name, Gender string
}
