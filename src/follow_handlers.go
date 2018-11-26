package main

import (
    "net/http"
    // "log"
    "encoding/json"
)

// Handler para seguir alguém
var FollowsFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    /* Followed data */
    decoder := json.NewDecoder(r.Body)
    
    var id struct {Id int `json:"id"`}
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    followed , err := GetUsers("id", id.Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }
    if len(followed) == 0 {
        http.Error(w, `"code": "Invalid user to follow"` , 400)
        return
    }

    err = FollowDBHandler(follower, followed[0].Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    w.Write([]byte("{code: success}"))
})

// Obter Follows
var FollowsFollowRequestsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    id, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }
    
    // Obter 

    users, err := FollowRequestsDBHandler(id);
    if  err!= nil{
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    js, err := json.Marshal(users)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
    }

    w.Write([]byte(js))
})

var FollowsAllowFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    /* Followed data */
    decoder := json.NewDecoder(r.Body)
    
    var id struct {Id int `json:"id"`}
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    followed , err := GetUsers("id", id.Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }
    if len(followed) == 0 {
        http.Error(w, `"code": "Invalid user to follow"` , 400)
        return
    }

    err = AcceptFollowDBHandler(followed[0].Id, follower)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

	response := "{code: success}"
    w.Write([]byte(response))
})

var FollowsGetFollowers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    users, err := DBGetFollowering(follower, true)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    js, err := json.Marshal(users)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
    }

    w.Write([]byte(js))
})

var FollowsGetFollowing = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    /* Follower data */
    id, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    users, err := DBGetFollowering(id, false)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    js, err := json.Marshal(users)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
    }

    w.Write([]byte(js))
})


var FollowsRemoveFollowerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    /* Followed data */
    decoder := json.NewDecoder(r.Body)
    
    var id struct {Id int `json:"id"`}
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    followed , err := GetUsers("id", id.Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }
    if len(followed) == 0 {
        http.Error(w, `"code": "Invalid user to follow"` , 400)
        return
    }

    err = DBRemoveFollowering(follower, followed[0].Id, false)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    response := "{code: success}"
    w.Write([]byte(response))
})


var FollowsRemoveFollowingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    /* Followed data */
    decoder := json.NewDecoder(r.Body)
    
    var id struct {Id int `json:"id"`}
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    followed , err := GetUsers("id", id.Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }
    if len(followed) == 0 {
        http.Error(w, `"code": "Invalid user to follow"` , 400)
        return
    }

    err = DBRemoveFollowering(follower, followed[0].Id, true)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

    response := "{code: success}"
    w.Write([]byte(response))
})
























