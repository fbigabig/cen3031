package main

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
	file, err := os.Open("C:\\Users\\aaaaaa\\go\\src\\github.com\\cen3031\\userlist.txt")
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
