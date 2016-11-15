package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
)

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func ValidateToken(c martini.Context, w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/login" && r.Method == "POST" {
		c.Next()
		return
	}

	tokenString := r.Header.Get("Authorization")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid
	at(time.Now(), func() {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
		if token == nil {
			w.Write([]byte("Bad Token"))
			return
		}
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)

			stat, accountItem := GetAccountItemByTokenMongo(tokenString)
			if stat == true {
				r.Header.Set("user_name", accountItem.UserName)
			}

			r.Header.Set("database_prefix", claims.Foo)
			c.Next()
		} else {
			fmt.Println(err)
			w.Write([]byte("Bad Token"))
		}
	})

}

func CreateJWTToken(iss string) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		iss,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 1000000,
			Issuer:    iss,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}
