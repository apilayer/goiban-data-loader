package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fourcube/goiban"
	co "github.com/fourcube/goiban/countries"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db               *sql.DB
	err              error
	bundesbankFile   = "data/bundesbank.txt"
	nbbFile          = "data/nbb.xlsx"
	netherlandsFile  = "data/nl.xlsx"
	PREP_ERR         error
	INSERT_BANK_DATA *sql.Stmt
	SELECT_SOURCE_ID *sql.Stmt
)

func prepareStatements() error {
	INSERT_BANK_DATA, err = db.Prepare(`INSERT INTO BANK_DATA
		(id, source, bankcode, name, zip, city, bic, country, algorithm, created, last_update)
		VALUES
		(NULL, ?, ?, ?, ?, ?, ?, ?, ?, NULL, NULL);`)

	if err != nil {
		log.Fatalf("Error while preparing statement: %v", err)
		return err
	}

	SELECT_SOURCE_ID, err = db.Prepare(`SELECT id FROM DATA_SOURCE where name = ?;`)

	if err != nil {
		log.Fatalf("Error while preparing statement: %v", err)
		return err
	}

	return nil
}

func getDataSourceId(sourceName string) (int, error) {
	var id int
	result := SELECT_SOURCE_ID.QueryRow(sourceName)

	err := result.Scan(&id)

	if err != nil {
		log.Fatalf("Data source %v not found: %v", sourceName, err)
		return -1, err
	}

	return id, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: goiban-data-loader <src> <dburl>")
		fmt.Println("e.g: goiban-data-loader <bundesbank|nbb|netherlands> root:root@/goiban?charset=utf8")
		return
	}

	target := os.Args[1]
	db, err = sql.Open("mysql", os.Args[2])

	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
		return
	}

	err = prepareStatements()

	if err != nil {
		log.Fatalf("DB Prepare Statement error: %v", err)
		return
	}

	ch := make(chan interface{})
	rows := 0

	switch target {
	default:
		fmt.Println("unknown target")
	case "bundesbank":
		go goiban.ReadFileToEntries(bundesbankFile, &co.BundesbankFileEntry{}, ch)

		source := "German Bundesbank"
		sourceId, err := getDataSourceId(source)

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			bbEntry := entry.(*co.BundesbankFileEntry)
			if bbEntry.M == 1 {
				_, err := INSERT_BANK_DATA.Exec(
					sourceId,
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
					log.Fatal(err, bbEntry)
				}
			}
		}
	case "nbb":
		go goiban.ReadFileToEntries(nbbFile, &co.BelgiumFileEntry{}, ch)

		source := "NBB"
		sourceId, err := getDataSourceId(source)

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			entries := entry.([]co.BelgiumFileEntry)
			for _, nbbEntry := range entries {
				_, err := INSERT_BANK_DATA.Exec(
					sourceId,
					nbbEntry.Bankcode,
					nbbEntry.Name,
					"",
					"",
					nbbEntry.Bic,
					"BE",
					"")
				if err == nil {
					rows++
				} else {
					log.Fatal(err, nbbEntry)
				}
			}
		}
	case "nl":
		go goiban.ReadFileToEntries(netherlandsFile, &co.NetherlandsFileEntry{}, ch)

		source := "NL"
		sourceId, err := getDataSourceId(source)

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			nlEntry := entry.(co.NetherlandsFileEntry)
			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				nlEntry.Bankcode,
				nlEntry.Name,
				"",
				"",
				nlEntry.Bic,
				"NL",
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, nlEntry)
			}
		}
	}

	log.Printf("Loaded %v for source '%v'", rows, target)

}
