package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

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

	//encrypts pw
	hashPW, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	credentials.Password = string(hashPW)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
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

	/* expirationTime := time.Now().Add(time.Minute * 5)

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

	// no err set cookie */

}

func Home(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	/* cookie, err := r.Cookie("token")

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

	*/

	// pass the token and make sure that the type is correct and valid

	// w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

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

type game struct {
	name        string
	platform    string
	releaseYear int
	developer   string
	publisher   string
}

type db struct {
	games    []game
	curGames []game
	sortType string
	//data     *sql.DB
}

func (gm *game) print() string {
	temp := ("Name:" + gm.name + "\t\tPlatform: " + gm.platform + "\t\tRelease Year: " + fmt.Sprint(gm.releaseYear) + "\t\tDeveloper: " + gm.developer + "\t\tPublisher: " + gm.publisher)
	return temp
}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
func (g *db) init() {
	g.games = make([]game, 0)
	g.curGames = make([]game, 0)
	g.sortType = "name"
	file, err := os.Open("gamelist.txt") //open userlist
	if err != nil {
		log.Fatal(err)
	}
	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() { //read in userlist
		var temp game
		temp.name = fileReader.Text()
		//fmt.Println(temp.name)
		fileReader.Scan()
		temp.platform = fileReader.Text()
		//fmt.Println(temp.platform)
		fileReader.Scan()
		temp.releaseYear, err = strconv.Atoi(fileReader.Text())
		handleErr(err)
		//fmt.Println(temp.releaseYear)
		fileReader.Scan()
		temp.developer = fileReader.Text()
		//fmt.Println(temp.developer)
		fileReader.Scan()
		temp.publisher = fileReader.Text()
		//fmt.Println(temp.publisher)
		g.games = append(g.games, temp)
	}
	err = file.Close()
	handleErr(err)

}
func (g *db) changeSort(newSort string) {
	g.sortType = newSort
}
func (g *db) addGame(name string, platform string, releaseYear int, developer string, publisher string) {
	file2, err := os.OpenFile("gamelist.txt", os.O_WRONLY|os.O_APPEND, 0644)
	handleErr(err)
	file2.WriteString(name + "\n")
	file2.WriteString(platform + "\n")
	file2.WriteString(fmt.Sprint(releaseYear) + "\n")
	file2.WriteString(developer + "\n")
	file2.WriteString(publisher + "\n")
	g.games = append(g.games, game{name, platform, releaseYear, developer, publisher})
	err = file2.Close()
	handleErr(err)
}
func (g *db) sort() {
	if g.sortType == "name" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].name < g.games[j].name
		})
	} else if g.sortType == "platform" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].platform < g.games[j].platform
		})
	} else if g.sortType == "releaseYear" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].releaseYear < g.games[j].releaseYear
		})
	} else if g.sortType == "developer" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].developer < g.games[j].developer
		})
	} else if g.sortType == "publisher" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].publisher < g.games[j].publisher
		})
	}
}
func (g *db) search(item string) []string {
	g.curGames = nil
	//fmt.Println(len(g.curGames))
	for _, v := range g.games {
		if strings.Contains(v.name, item) /*|| strings.Contains(v.platform, item) || strings.Contains(v.developer, item) || strings.Contains(v.publisher, item) */ {
			g.curGames = append(g.curGames, v)
			//fmt.Println(v.name, "name")
		} else if strings.Contains(v.platform, item) {
			g.curGames = append(g.curGames, v)
			//fmt.Println(v.platform, "platform")
		} else if strings.Contains(v.developer, item) {
			g.curGames = append(g.curGames, v)
			//fmt.Println(v.developer, "dev")
		} else if strings.Contains(v.publisher, item) {
			g.curGames = append(g.curGames, v)
			//fmt.Println(v.publisher, "pub")
		}
		//fmt.Println(len(g.curGames))

	}
	return g.printSearch()
}
func (g db) print() []string {
	output := make([]string, 0)
	for _, v := range g.games {
		output = append(output, v.print())
	}
	return output
}
func (g db) printSearch() []string {
	output := make([]string, 0)
	for _, v := range g.curGames {
		//fmt.Println(v.name)
		output = append(output, v.print())
	}
	//fmt.Println(len(output))
	return output
}
func test() {
	var games db
	fo, err := os.Create("output.txt")

	handleErr(err)
	games.init()
	games.sort()
	//fo.WriteString(games.games[0].name)
	fo.WriteString("test1:\n")
	fo.WriteString(strings.Join(games.print(), "\n"))
	fo.WriteString("\n")
	games.addGame("Test", "test", 2000, "test", "test")
	games.sort()
	fo.WriteString("test2:\n")
	fo.WriteString(strings.Join(games.print(), "\n"))
	fo.WriteString("\n")
	fo.WriteString("test3:\n")
	fo.WriteString(strings.Join(games.search("test"), "\n"))
	fo.WriteString("\n")
	games.changeSort("releaseYear")
	games.sort()
	fo.WriteString("test4:\n")
	fo.WriteString(strings.Join(games.print(), "\n"))
}
func main() {
	test() // test function to test the gamedb api

	http.HandleFunc("/create", Create)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
