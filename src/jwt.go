


package main

import (
	"io/ioutil"
	"time"
    "os"
    "path/filepath"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

func getDir() string  {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    return filepath.Dir(ex)
}

func GenerateJWT(id int, nome string) []byte {

	dir := getDir()

	bytes, _ := ioutil.ReadFile(dir + "/keys/sample_key.priv")
	claims := jws.Claims{}
	
	claims.SetExpiration(time.Now().Add(time.Duration(1000) * time.Second))
	claims.SetIssuedAt(time.Now())
	
	claims.SetIssuer("JB's Micro Service")
	claims.SetSubject(nome)
	claims.Set("Id", id)
	
	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	
	accessToken, _ := jws.NewJWT(claims, crypto.SigningMethodRS256).Serialize(rsaPrivate)
	return accessToken
}
	
func ValidateJWT(accessToken []byte) (id int, err error){
	
	dir := getDir()

	bytes, _ := ioutil.ReadFile(dir + "/keys/sample_key.pub")
	rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM(bytes)

	j, err := jws.ParseJWT([]byte(accessToken))

	err = j.Validate(rsaPublic, crypto.SigningMethodRS256)
	id = int(jwt.Claims(j.Claims()).Get("Id").(float64))


	return id, err
	
	
}

