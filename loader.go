package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stefan93/goiban"
	co "github.com/stefan93/goiban/countries"
)

var (
	db                *sql.DB
	err               error
	bundesbankFile    = "data/bundesbank.txt"
	nbbFile           = "data/nbb.xlsx"
	netherlandsFile   = "data/nl.xlsx"
	luxembourgFile    = "data/lu.xlsx"
	switzerlandFile   = "data/ch.txt"
	austriaFile       = "data/at.csv"
	liechtensteinFile = "data/li.xlsx"

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
		fmt.Println("e.g: goiban-data-loader <bundesbank|nbb|nl|lu|ch> root:root@/goiban?charset=utf8")
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

			bic := strings.Replace(bbEntry.Bic, " ", "", -1)

			if bbEntry.M == 1 {
				_, err := INSERT_BANK_DATA.Exec(
					sourceId,
					bbEntry.Bankcode,
					bbEntry.Name,
					bbEntry.Zip,
					bbEntry.City,
					bic,
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
				// Remove spaces from BIC
				bic := strings.Replace(nbbEntry.Bic, " ", "", -1)

				_, err := INSERT_BANK_DATA.Exec(
					sourceId,
					nbbEntry.Bankcode,
					nbbEntry.Name,
					"",
					"",
					bic,
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

			bic := strings.Replace(nlEntry.Bic, " ", "", -1)

			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				nlEntry.Bankcode,
				nlEntry.Name,
				"",
				"",
				bic,
				"NL",
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, nlEntry)
			}
		}
	case "lu":
		go goiban.ReadFileToEntries(luxembourgFile, &co.LuxembourgFileEntry{}, ch)

		source := "LU"
		sourceId, err := getDataSourceId(source)

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			luEntry := entry.(co.LuxembourgFileEntry)

			if strings.TrimSpace(luEntry.Bankcode) == "" {
				log.Printf("Skipping invalid entry without Bankcode %v", luEntry)
				continue
			}

			if strings.TrimSpace(luEntry.Bic) == "" {
				log.Printf("Skipping invalid entry without BIC %v", luEntry)
				continue
			}

			bic := strings.Replace(luEntry.Bic, " ", "", -1)

			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				luEntry.Bankcode,
				luEntry.Name,
				"",
				"",
				bic,
				"LU",
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, luEntry)
			}
		}

	case "ch":
		go goiban.ReadFileToEntries(switzerlandFile, &co.SwitzerlandBankFileEntry{}, ch)

		source := "CH"
		sourceId, err := getDataSourceId(source)
		err = nil
		uniqueEntries := map[string]co.SwitzerlandBankFileEntry{}

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			chEntry := entry.(*co.SwitzerlandBankFileEntry)

			if strings.TrimSpace(chEntry.BankCode) == "" {
				log.Printf("Skipping invalid entry without Bankcode %v", chEntry)
				continue
			}

			if strings.TrimSpace(chEntry.Bic) == "" {
				log.Printf("Skipping invalid entry without BIC %v", chEntry)
				continue
			}

			chEntry.Bic = strings.Replace(chEntry.Bic, " ", "", -1)

			uniqueEntries[chEntry.BankCode] = *chEntry
		}

		for _, chEntry := range uniqueEntries {
			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				chEntry.BankCode,
				chEntry.BankName,
				chEntry.Zip,
				chEntry.Place,
				chEntry.Bic,
				source,
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, chEntry)
			}
		}
	case "li":
		go goiban.ReadFileToEntries(liechtensteinFile, &co.LiechtensteinFileEntry{}, ch)

		source := "LI"
		sourceId, err := getDataSourceId(source)
		uniqueEntries := map[string]co.LiechtensteinFileEntry{}

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			chEntry := entry.(co.LiechtensteinFileEntry)

			if strings.TrimSpace(chEntry.Bankcode) == "" {
				log.Printf("Skipping invalid entry without Bankcode %v", chEntry)
				continue
			}

			if strings.TrimSpace(chEntry.Bic) == "" {
				log.Printf("Skipping invalid entry without BIC %v", chEntry)
				continue
			}

			chEntry.Bic = strings.Replace(chEntry.Bic, " ", "", -1)

			uniqueEntries[chEntry.Bankcode] = chEntry
		}

		for _, chEntry := range uniqueEntries {
			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				chEntry.Bankcode,
				chEntry.Name,
				"",
				"",
				chEntry.Bic,
				source,
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, chEntry)
			}
		}
	case "at":
		go goiban.ReadFileToEntries(austriaFile, &co.AustriaBankFileEntry{}, ch)

		source := "AT"
		sourceId, err := getDataSourceId(source)
		uniqueEntries := map[string]co.AustriaBankFileEntry{}

		if err != nil {
			log.Fatalf("Aborting: %v", err)
			return
		}

		log.Printf("Removing entries for source '%v'", source)
		db.Exec("DELETE FROM BANK_DATA WHERE source = ?", sourceId)

		for entry := range ch {
			chEntry := entry.(*co.AustriaBankFileEntry)

			if strings.TrimSpace(chEntry.Bankcode) == "" {
				log.Printf("Skipping invalid entry without Bankcode %v", chEntry)
				continue
			}

			if strings.TrimSpace(chEntry.Bic) == "" {
				log.Printf("Skipping invalid entry without BIC %v", chEntry)
				continue
			}

			chEntry.Bic = strings.Replace(chEntry.Bic, " ", "", -1)

			uniqueEntries[chEntry.Bankcode] = *chEntry
		}

		for _, chEntry := range uniqueEntries {
			_, err := INSERT_BANK_DATA.Exec(
				sourceId,
				chEntry.Bankcode,
				chEntry.Name,
				"",
				"",
				chEntry.Bic,
				source,
				"")
			if err == nil {
				rows++
			} else {
				log.Fatal(err, chEntry)
			}
		}

	}

	log.Printf("Loaded %v for source '%v'", rows, target)

}
