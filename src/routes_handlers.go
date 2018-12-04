package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

var RoutesCreateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    /* Follower data */
    follower, err := ValidateSession(r)
    if err != nil  {
        http.Error(w, `"code": "Invalid Session. Please Relog"` , 400)
        return
    }

    /* Followed data */
    decoder := json.NewDecoder(r.Body)
    
    var id struct {Description string `json:"description"`}
    err = decoder.Decode(&id)
    if err != nil {
        http.Error(w, `"code": "Invalid json format"` , 400)
        return
    }

    res, err := NewRouteDBHandler(follower, id.Description)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

	response := fmt.Sprintf("{\"code\": \"%d\"}", res)
    w.Write([]byte(response))
})


var RoutesRemoveHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
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

    err = DelRouteDBHandler(follower, id.Id)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
        return
    }

	response := "{code: success}"
    w.Write([]byte(response))
})


var RoutesGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    /* Follower data */
    _, err := ValidateSession(r)
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

	/*foll, err := TestFollow(follower, id.Id)
    if err != nil {
        http.Error(w, `"code": "Invalid Ids"` , 400)
        return
    }*/

    var route_list []Route
    route_list , err = GetRoutesRouteDBHandler(id.Id)


    /*var route_list []Route
    if foll || follower == id.Id{
    route_list , err = GetRoutesRouteDBHandler(id.Id)
    } else{
	    http.Error(w, `"code": "Invalid Id or you don't follow the guy!"` , 400)
	    return
    }*/

    js, err := json.Marshal(route_list)
    if err != nil {
        http.Error(w, `{"code": "`+ err.Error()+ `"}`, 400)
    }

    w.Write([]byte(js))

})
