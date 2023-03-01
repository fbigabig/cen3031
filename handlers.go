package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"pablo": "bueno",
	"aaron": "gill",
}

// temporary should be from db in actual implementation

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// the credentials are sent as json

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	var games db
	games.init()

	http.HandleFunc("/create", Create)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handles the three funcs and creates basic local server

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func Create(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var credentials Credentials
	// user and pass object

	err := json.NewDecoder(r.Body).Decode(&credentials)
	// decode the pass if no err continue

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//hash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 10)

	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	//encrypts pw
	hashPW, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	credentials.Password = string(hashPW)
	if err != nil {
		log.Fatal(err)
	}
	users[credentials.Username] = credentials.Password

	// adds the user and pass to the list and allows it to log in

	// *needs to be fixed to allow encoding* and *databases*
}

func Login(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var credentials Credentials
	// user and pass object

	err := json.NewDecoder(r.Body).Decode(&credentials)
	// decode the pass if no err continue

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//decrypts pw
	expectedPassword, ok := users[credentials.Username]
	err = bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(credentials.Password))
	matchPW := err == nil
	if !ok || !matchPW {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// if wrong pass show its unauthorized

	expirationTime := time.Now().Add(time.Minute * 5)

	// if user and pass is correct then create a 5 min time for token

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// storing into the claims struct the user and time
	// expiration time deonotes how long the token lasts

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	// create the token

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// no err set cookie

}

func Home(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	cookie, err := r.Cookie("token")

	// get cookie from previous method

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// makes sure a cookie of correct type is given

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	// get token string if valid

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// pass the token and make sure that the type is correct and valid

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

	// if the token is valid pass data
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}
