package main

import (

)

type MyStruct struct {
    SomeField string 
}

type User struct {
    Id int      	`json:"id,omitempty"`
    Age int         `json:"age,omitempty"`
    Email string    `json:"email,omitempty"`
    Name string     `json:"name,omitempty"`
    Gender string   `json:"gender,omitempty"`
    Pass string     `json:"pass,omitempty"`
    Hash [32]byte   `json:"hash,omitempty"`
    Salt []byte     `json:"salt,omitempty"`
}

type UserToOutside struct {
    Id , Age int
    Email, Name, Gender string
}
