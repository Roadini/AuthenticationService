package main

import (
    "database/sql"
	"log"
)

func NewRouteDBHandler(id_user int, description string) (insertedId int64, err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    insertUser, err := DataBase.Prepare("INSERT INTO routes (id_user, description) VALUES (?, ?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    res, err := insertUser.Exec( id_user , description)
    insertedId, err = res.LastInsertId()

    return
}

func DelRouteDBHandler(id_user int, id int) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "DELETE FROM routes WHERE id = ? AND id_user = ?"

    insertUser, err := DataBase.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer insertUser.Close()

	_, err = insertUser.Exec(id, id_user)
   
    return err
}



func GetRoutesRouteDBHandler(id int) (route_list []Route, err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "SELECT id, description FROM routes WHERE id_user = ?"

    rows, err := DataBase.Query(query, id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()


    for rows.Next(){
        var u Route
        if err := rows.Scan(&u.Id, &u.Description); err != nil {
            log.Fatal(err)
        }

        route_list = append(route_list, u)
    }

    return 
}