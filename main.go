package main

import (
	"fmt"
	"os"

	"github.com/mateors/mcb"
)

// type client struct {
// 	Callsing string
// 	Country  string
// 	Iata     string
// 	Icao     string
// 	Name     string
// 	Id       string
// 	Type     string
// }

var db *mcb.DB

func init() {

	db = mcb.Connect(host, username, password, false)

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {
	InsertData()
}
