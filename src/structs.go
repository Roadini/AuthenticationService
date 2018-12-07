package main

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "encoding/json"
    "time"
)

type MyStruct struct {
    SomeField string
}

type User struct {
    Id int		        `json:"id,omitempty"`
    Age int             `json:"age,omitempty"`
    Email string        `json:"email,omitempty"`
    Name string         `json:"name,omitempty"`
    Gender string       `json:"gender,omitempty"`
    Password string     `json:"password,omitempty"`
    Description string  `json:"description,omitempty"`
    Hash [32]byte       `json:"hash,omitempty"`
    Salt []byte         `json:"salt,omitempty"`
}

    description varchar(255)    DEFAULT NULL,

    accessToken varchar(300)    DEFAULT NULL,
    id          BIGINT         DEFAULT NULL,

type UserToOutside struct {
    Id              int     `json:"id,omitempty"`
    Age             NullInt64 `json:"age,omitempty"`
    Email           string  `json:"email,omitempty"`
    Name            string  `json:"name,omitempty"`
    Description     NullString  `json:"description,omitempty"`
    Gender          NullString  `json:"gender,omitempty"`
}


type UserToOutsideFB struct {
    Id              int     `json:"id,omitempty"`
    Age             NullInt64     `json:"age,omitempty"`
    Email           string  `json:"email,omitempty"`
    Name            string  `json:"name,omitempty"`
    Gender          NullString  `json:"gender,omitempty"`
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

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}


// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

