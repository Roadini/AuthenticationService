


package main

import (
	"io/ioutil"
	"time"
    "os"
    "path/filepath"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"log"
)

func getDir() string  {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    return filepath.Dir(ex)
}

func GenerateJWT(id int, nome string, token string) []byte {

	dir := getDir()

	bytes, _ := ioutil.ReadFile(dir + "/keys/sample_key.priv")
	claims := jws.Claims{}
	
	claims.SetExpiration(time.Now().Add(time.Duration(100) * time.Second))
	claims.SetIssuedAt(time.Now())
	
	claims.SetIssuer("JB's Micro Service")
	claims.SetSubject(nome)
	claims.Set("Id", id)
	claims.Set("Token",token)
	
	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	
	accessToken, _ := jws.NewJWT(claims, crypto.SigningMethodRS256).Serialize(rsaPrivate)
	return accessToken
}
	
func ValidateJWT(accessToken []byte) {
	
	dir := getDir()

	bytes, _ := ioutil.ReadFile(dir + "/keys/sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes)

	jwt, err := jws.ParseJWT([]byte(accessToken))
	if err != nil {
		log.Fatal(err)
	}

	// Validate token
	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		log.Fatal(err)
	}
	
	
}

