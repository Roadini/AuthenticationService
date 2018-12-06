package main

import (
    "net/http"
    "encoding/json"
    "log"
)

var LikesGetLikesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    _, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }
    log.Printf("GetLikes\n")


    decoder := json.NewDecoder(r.Body)
    var id Like
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }
    var user_list []UserToOutside
    user_list , err = GetLikesDBHandler(id.Id_route)


    js, err := json.Marshal(user_list)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
    } else if len(user_list) == 0 {
        w.Write([]byte("[]"))
    } else {
        w.Write([]byte(js))
    }

})

var LikesLikeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    id_user, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    decoder := json.NewDecoder(r.Body)
    var id Like
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    log.Printf("Like (User_ID: %d; Route_ID: %d)\n", id_user, id.Id_route)

    err = LikeDBHandler(id.Id_route, id_user)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

	response := "{code: success}"
    w.Write([]byte(response))
})


var LikesUnlikeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    id_user, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    decoder := json.NewDecoder(r.Body)
    var id Like
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    log.Printf("Like (User_ID: %d; Route_ID: %d)\n", id_user, id.Id_route)

    err = UnlikeDBHandler(id.Id_route, id_user)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    response := "{code: success}"
    w.Write([]byte(response))
})