package main

import (
	"os"
	"database/sql"
	"fmt"
	"github.com/fourcube/goiban"
	co "github.com/fourcube/goiban/countries"
	_ "github.com/go-sql-driver/mysql"

)



var (
	db *sql.DB
	err error
	bundesbankFile = "data/bundesbank.txt"  
	PREP_ERR error
	INSERT_BANK_DATA *sql.Stmt 
	
)


func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: goiban-data-loader <src> <dburl>")
		fmt.Println("e.g: goiban-data-loader bb root:root@/goiban")
		return
	}
	target := os.Args[1]	
	db, err = sql.Open("mysql", os.Args[2])

	INSERT_BANK_DATA, PREP_ERR = db.Prepare(`INSERT INTO BANK_DATA
		(id, bankcode, name, zip, city, bic, country, algorithm, created, last_update)
		VALUES
		(NULL, ?, ?, ?, ?, ?, ?, ?, NULL, NULL);`)
	ch := make(chan interface{})
	rows := 0
	
	switch target {
	default:
		fmt.Println("unknown target")
	case "bb":
		go goiban.ReadFileToEntries(bundesbankFile, &co.BundesbankFileEntry{}, ch)
		
		for entry := range ch {
			bbEntry := entry.(*co.BundesbankFileEntry)
			if bbEntry.M == 1 {
				_, err := INSERT_BANK_DATA.Exec(
					bbEntry.Bankcode,
					bbEntry.Name,
					bbEntry.Zip,
					bbEntry.City,
					bbEntry.Bic,
					"DE",
					bbEntry.CheckAlgo)
				if err == nil {
					rows++
				} else {
					fmt.Println(err, bbEntry)
				}
			}
		}

	}

	fmt.Println("Loaded", rows, "for", target);

}