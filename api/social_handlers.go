package main

import (
    "net/http"
)

// Handler para seguir alguém
var FollowsFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})

// Obter Follows
var FollowsFollowRequestsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{users : [{Name: João Teste; Age:18; Gender: Male;}]}"

    w.Write([]byte(response))
})

var FollowsAllowFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})

var FollowsRemoveFollowerHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})


var FollowsRemoveFollowingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})


var LikesGetLikesHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{users : [{Name: João Teste; Age:18; Gender: Male;}]}"
    w.Write([]byte(response))
})


var LikesLikeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})


var LikesUnlikeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})

var RoutesCreateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})


var RoutesRemoveHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{code: success}"
    w.Write([]byte(response))
})


var RoutesGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := "{routes : [3, 5]; putas: teste}"
    w.Write([]byte(response))
})