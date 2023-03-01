package main

import (
	//"database/sql"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	//_ "github.com/mattn/go-sqlite3"
)

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

func (gm game) print() string {
	temp := ("Name:" + gm.name + "\tPlatform: " + gm.platform + "\tRelease Year: " + fmt.Sprint(gm.releaseYear) + "\tDeveloper: " + gm.developer + "\tPublisher: " + gm.publisher)
	return temp
}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
func (g db) init() {
	g.games = make([]game, 0)
	//curGames = make([]game, 0)
	g.sortType = "name"
	file, err := os.Open("gamelist.txt") //open userlist
	if err != nil {
		log.Fatal(err)
	}
	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() { //read in userlist
		var temp game
		temp.name = fileReader.Text()
		fileReader.Scan()
		temp.platform = fileReader.Text()
		//fmt.Println("UNPW: " + name + " " + password + "\n")
		fileReader.Scan()
		temp.releaseYear, err = strconv.Atoi(fileReader.Text())
		handleErr(err)
		fileReader.Scan()
		temp.developer = fileReader.Text()
		fileReader.Scan()
		temp.publisher = fileReader.Text()
		g.games = append(g.games, temp)
	}
	err = file.Close()
	handleErr(err)

}
func (g db) changeSort(newSort string) {
	g.sortType = newSort
}
func (g db) addGame(name string, platform string, releaseYear int, developer string, publisher string) {
	file2, err := os.OpenFile("userlist.txt", os.O_WRONLY|os.O_APPEND, 0644)
	handleErr(err)
	file2.WriteString(name + "\n")
	file2.WriteString(platform + "\n")
	file2.WriteString(fmt.Sprint(releaseYear) + "\n")
	file2.WriteString(developer + "\n")
	file2.WriteString(publisher + "\n")
	g.games = append(g.games, game{name, platform, releaseYear, developer, publisher})
}
func (g db) sort() {
	if g.sortType == "name" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].name > g.games[j].name
		})
	} else if g.sortType == "platform" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].platform > g.games[j].platform
		})
	} else if g.sortType == "releaseYear" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].releaseYear > g.games[j].releaseYear
		})
	} else if g.sortType == "developer" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].developer > g.games[j].developer
		})
	} else if g.sortType == "publisher" {
		sort.Slice(g.games, func(i, j int) bool {
			return g.games[i].publisher > g.games[j].publisher
		})
	}
}
func (g db) search(item string) []string {
	g.curGames = nil
	for _, v := range g.games {
		if strings.Contains(v.name, item) || strings.Contains(v.platform, item) || strings.Contains(v.developer, item) || strings.Contains(v.publisher, item) {
			g.curGames = append(g.games, v)
		}
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
		output = append(output, v.print())
	}
	return output
}
