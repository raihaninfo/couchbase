package main

import (
	"fmt"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

// CRUD

// database connection
func init() {
	db = mcb.Connect(host, username, password, false)
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
}

// royaltypool
func main() {
	// INSERT data
	// data := db.Query(`INSERT INTO royaltypool.raihan.client (KEY, VALUE) VALUES(
	// 	'clid-323',
	// 	{
	// 		"name": "raihan",
	// 		"country": "bangladesh",
	// 		"id":"99",
	// 		"type":"client"
	// 	}
	// )`)
	// fmt.Println(data.Result)

	// READ data
	// data := db.Query(`SELECT * FROM royaltypool.raihan.client WHERE id='98'`)
	// fmt.Println(data.Result)

	// DELETE data
	// data := db.Query(`DELETE FROM royaltypool.raihan.client WHERE id='99'`)
	// fmt.Println(data.Result)

	// UPDATE data
	data := db.Query(`UPDATE royaltypool.raihan.client SET country='US' WHERE id='343'`)

	fmt.Println(data.Result, "\n", data.Errors)

}
