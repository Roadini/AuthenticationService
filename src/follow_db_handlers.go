package main

import (
    "database/sql"
	"log"
)

func FollowDBHandler(id_follower int, id_followed int) (err error){
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()


    insertUser, err := DataBase.Prepare("INSERT INTO follows ( id_follower , id_followed, accepted ) VALUES (?, ?, ?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec( id_follower , id_followed, 0)
    return err
}


func FollowRequestsDBHandler(id_followed int) (user_list []UserToOutside, err error){

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    rows, err := DataBase.Query("SELECT id_follower FROM follows WHERE id_followed = ? AND accepted = 0", id_followed)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for rows.Next(){
        var u Follows
        if err := rows.Scan(&u.Id_followed); err != nil {
            log.Fatal(err)
        }
        var user[] UserToOutside
       	user, err = GetUsers("id", u.Id_followed)
        if err != nil {
            log.Fatal(err)
        }
        user_list = append(user_list, user[0])
    }
    return user_list, err
}

func AcceptFollowDBHandler(id_follower int, id_followed int) (err error){

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    insertUser, err := DataBase.Prepare("UPDATE follows SET accepted = 1 WHERE id_follower = ? AND id_followed = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer insertUser.Close() // Close the statement when we leave main() / the program terminates

    _, err = insertUser.Exec(id_follower , id_followed)

    return err
}

func DBGetFollowering(id int, me bool) (user_list []UserToOutside, err error) {

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := ""
    if me {
    	query = "SELECT id_follower FROM follows WHERE id_followed = ? AND accepted = 1"
    }else{
    	query = "SELECT id_followed FROM follows WHERE id_follower = ? AND accepted = 1"
    }


	rows, err := DataBase.Query(query, id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for rows.Next(){
        var u int
        if err := rows.Scan(&u); err != nil {
            log.Fatal(err)
        }
        var user[] UserToOutside
       	user, err = GetUsers("id", u)
        if err != nil {
            log.Fatal(err)
        }
        user_list = append(user_list, user[0])
    }
    return user_list, err

}

func DBRemoveFollowering(id int, other_id int, following bool) (err error) {

    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "DELETE FROM follows WHERE id_follower = ? AND id_followed = ?"

    insertUser, err := DataBase.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer insertUser.Close()

    if following {
    	_, err = insertUser.Exec(id, other_id)
    } else {
    	_, err = insertUser.Exec(other_id, id)
    }
   
    return err

}



func TestFollow(user1 int, user2 int) (foll bool, err error) {
    DataBase,  err := sql.Open("mysql", connect)
    if err != nil {
        panic(err.Error()) 
    }
    defer DataBase.Close()

    query := "SELECT * FROM follows WHERE id_follower = ? AND id_followed = ?"
    rows, err := DataBase.Query(query, user1, user2)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer rows.Close()

    foll = false
    for rows.Next(){
        foll = true
    }

    return

}