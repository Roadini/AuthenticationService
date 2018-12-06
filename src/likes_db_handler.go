package main

import (
    "database/sql"
    mysql "github.com/go-sql-driver/mysql"
	"log"
    "errors"
)

func LikeDBHandler(id_route int, id_user int) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    insertUser, err := DataBase.Prepare("INSERT INTO likes (id_route, id_user) VALUES (?, ?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec( id_route , id_user)
    if err != nil && err.(*mysql.MySQLError).Number == 1062{
        err = errors.New("Duplicate like")
    }

    return
}

func UnlikeDBHandler(id_route int, id_user int) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "DELETE FROM likes WHERE id_route = ? AND id_user = ?"

    insertUser, err := DataBase.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer insertUser.Close()

	_, err = insertUser.Exec(id_route, id_user)
   
    return err
}




func GetLikesDBHandler(route int) (user_list []UserToOutside, err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "SELECT id_user FROM likes WHERE id_route = ?"

    rows, err := DataBase.Query(query, route)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    var tmp_list []int

    for rows.Next(){
        var u int
        if err := rows.Scan(&u); err != nil {
            log.Fatal(err)
        }

        tmp_list = append(tmp_list, u)
    }

    for _, user_id := range tmp_list {
        u, err := GetUsers("id", user_id)
        if err != nil {
            log.Fatal(err)
        }
        user_list = append(user_list, u[0])
    }
    return 
}