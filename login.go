package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var userList []User

func startUp(web http.ResponseWriter, req *http.Request) {
	//
	tmpl := template.Must(template.ParseFiles("form.html"))
	if req.Method != http.MethodPost {
		tmpl.Execute(web, nil)
		return
	}
	file, err := os.Open("userlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	userList = make([]User, 0)
	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() { //read in userlist
		name := fileReader.Text()
		fileReader.Scan()
		password := fileReader.Text()
		//fmt.Println("UNPW: " + name + " " + password + "\n")
		tempUser := new(User)
		tempUser.username = name
		tempUser.password = password
		userList = append(userList, *tempUser)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	web.Header().Add("Content-Type", "application/json")
	username, password, isIn := req.BasicAuth()
	// uses built in Basic Auth func to check if user is on the site

	if !isIn {
		web.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		web.WriteHeader(http.StatusUnauthorized)
		web.Write([]byte("uh oh, something went wrong!"))
		return
	}
	// prompts user to enter credentials pop on the site also on terminal

	if !login(username, password, &userList) {
		web.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		web.WriteHeader(http.StatusUnauthorized)
		web.Write([]byte("Incorrect Info"))
		return
	}
	// if the user and pass do not match display message
	web.WriteHeader(http.StatusOK)
	web.Write([]byte("Log in Successful"))

	// once logged in display welcome message
	file2, err := os.OpenFile("userlist.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	//get new user's info

	newUsername := req.FormValue("username")
	newPassword := req.FormValue("password")

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file2.WriteString(newUsername + "\n" + string(hashedPW) + "\n") //store new user's info
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = file2.Close()
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(web, struct{ Success bool }{true})

}

type User struct {
	username     string
	password     string
	creationDate int
	email        string
}

func authenticate(inName string, inPassword string, matchUser *User) bool { //authenticate password, this needs to be made secure, rn everything is plaintext
	err := bcrypt.CompareHashAndPassword([]byte(matchUser.password), []byte(inPassword))
	return err == nil
}
func login(inName string, inPassword string, userlist *[]User) bool { //checks if the user exists, then calls authenticate
	for _, v := range *userlist { //looks through the list of users

		if v.username == inName {
			return authenticate(inName, inPassword, &v)
		}
	}
	return false
}
func getInput(reader *bufio.Scanner) string { //simplifies getting input from the user
	reader.Scan()

	input := reader.Text()
	input = strings.TrimSuffix(input, "\n")
	return input
}
func doLogIn(reader *bufio.Scanner, userList *[]User) bool {
	fmt.Println("Username") //get user info
	var username string = getInput(reader)
	fmt.Println("Password")
	var password string = getInput(reader)

	return login(username, password, userList) //try to login with the info
}
func main() {
	//reader := bufio.NewScanner(os.Stdin)
	http.HandleFunc("/", startUp)
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	/* var loggedIn bool = doLogIn(reader, &userList)
	if loggedIn {
		fmt.Println("Logged in")
	} else {
		for !loggedIn {
			fmt.Println("Login failed")
			fmt.Println("Type 1 to try again, or anything else to exit.")
			var choice string = getInput(reader)
			if choice == "1" {
				loggedIn = doLogIn(reader, &userList)
			} else {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
		fmt.Println("Logged in")
	}*/
	/*
		if choice == "1" {
	*/

	/*
		} else {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	*/

}

// checks if the user and pass match based on the map returns bool

/*
func main() {
	http.HandleFunc("/", startUp)
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// creating a http server using go
}
*/
// based on the code on https://umesh.dev/blog/how-to-implement-http-basic-auth-in-gogolang/
