package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

type master struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Profession string   `json:"profession"`
	Hobbies    []string `json:"hobbies"`
	Type       string   `json:"type"`
}

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

	//How to insert into couchbase bucket
	var myData master

	form := make(url.Values, 0)
	form.Add("bucket", "royaltypool") //bucket and collection-> namespace:bucket.scope.collection
	form.Add("aid", "d006")           //document ID
	form.Add("name", "Mostain Billah")
	form.Add("age", "36")
	form.Add("profession", "Developer")
	form.Add("hobbies", "Programming")
	form.Add("hobbies", "Problem Solving")
	form.Add("type", "participant") //what type of data or table name in general (SQL)

	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

}
