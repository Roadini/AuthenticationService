package main

import (

)

type MyStruct struct {
    SomeField string 
}

type User struct {
    Id int      	    `json:"id,omitempty"`
    Age int             `json:"age,omitempty"`
    Email string        `json:"email,omitempty"`
    Name string         `json:"name,omitempty"`
    Gender string       `json:"gender,omitempty"`
    Password string     `json:"password,omitempty"`
    Description string  `json:"description,omitempty"`
    Hash [32]byte       `json:"hash,omitempty"`
    Salt []byte         `json:"salt,omitempty"`
}

type UserToOutside struct {
    Id              int     `json:"id,omitempty"`
    Age             int     `json:"age,omitempty"`
    Description     string  `json:"description,omitempty"`
    Email           string  `json:"email,omitempty"`
    Name            string  `json:"name,omitempty"`
    Gender          string  `json:"gender,omitempty"`
}

type UserToOutsideFB struct {
    Id              int     `json:"id,omitempty"`
    Age             int     `json:"age,omitempty"`
    Email           string  `json:"email,omitempty"`
    Name            string  `json:"name,omitempty"`
    Gender          string  `json:"gender,omitempty"`
}

type Follows struct {
    Id_follower int `json:"id_follower,omitempty"`
    Id_followed int `json:"id_followed,omitempty"`
    Accepted string `json:"accepted,omitempty"`
}

type Route struct{
    Id int              `json:"id,omitempty"`
    Description string  `json:"description,omitempty"`
}

type Like struct{
    Id_route    int  `json:"id_route,omitempty"`
    Id_user     int  `json:"id_user,omitempty"`
}

type FB_Login struct{
    Id              string  `json:"userID,omitempty"`
    AccessToken     string  `json:"accessToken,omitempty"`
    SignedRequest   string  `json:"signedRequest,omitempty"`
}