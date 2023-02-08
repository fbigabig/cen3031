package main

/*
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	username     string
	password     string
	creationDate int
	email        string
}

func authenticate(inName string, inPassword string, matchUser *User) bool { //authenticate password, this needs to be made secure, rn everything is plaintext
	if matchUser.password == inPassword {
		fmt.Println("success")
		return true
	}
	fmt.Println("fail")
	return false
}
func login(inName string, inPassword string, userlist *[]User) {
	for _, v := range *userlist { //looks through the list of users

		if v.username == inName {
			authenticate(inName, inPassword, &v)
		}
	}
	fmt.Println("done")
}
func getInput(reader *bufio.Scanner) string { //simplifies getting input from the user
	reader.Scan()

	input := reader.Text()
	input = strings.TrimSuffix(input, "\n")
	return input
}


func main() {
	reader := bufio.NewScanner(os.Stdin)
	file, err := os.Open("userlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	userList := make([]User, 0)

	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() { //read in userlist
		name := fileReader.Text()
		fileReader.Scan()
		password := fileReader.Text()
		tempUser := new(User)
		tempUser.username = name
		tempUser.password = password
		userList = append(userList, *tempUser)
	}
	fmt.Println("Username") //get user info
	var username string = getInput(reader)
	fmt.Println("Password")
	var password string = getInput(reader)
	login(username, password, &userList) //try to login with the info

}
*/

import (
	"fmt"
	"log"
	"net/http"
)

var users = map[string]string{
	"testUser": "Pass",
}
// map which holds the user names and pass words will be hashed on the client side

func isAuth(username, password string) bool {
	pass, isAuthorrized := users[username]
	if !isAuthorrized {
		return false
	}

	return password == pass
}
// checks if the user and pass match based on the map returns bool

func startUp(web http.ResponseWriter, rep *http.Request) {
	web.Header().Add("Content-Type", "application/json")
	username, password, isIn := rep.BasicAuth()
	// uses built in Basic Auth func to check if user is on the site

	if !isIn {
		web.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		web.WriteHeader(http.StatusUnauthorized)
		web.Write([]byte(`{"message": "No User Was Given"}`))
		return
	}
	// prompts user to enter credentials pop on the site also on terminal 

	if !isAuth(username, password) {
		web.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		web.WriteHeader(http.StatusUnauthorized)
		web.Write([]byte(`{"message": "Incorrect Info"}`))
		return
	}
	// if the user and pass do not match display message

	web.WriteHeader(http.StatusOK)
	web.Write([]byte(`{"message": "Log in Successful"}`))
	// once logged in display welcome message

	return
}

func main() {
	http.HandleFunc("/", startUp)
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// creating a http server using go
}

// based on the code on https://umesh.dev/blog/how-to-implement-http-basic-auth-in-gogolang/

