package main

import (
    "net/http"
    // "encoding/json"
)

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