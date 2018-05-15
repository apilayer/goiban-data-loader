package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fourcube/goiban-data"
	"github.com/fourcube/goiban-data-loader/loader"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db                *sql.DB
	err               error
	bundesbankFile    = "data/bundesbank.txt"
	nbbFile           = "data/nbb.xlsx"
	netherlandsFile   = "data/nl.xlsx"
	luxembourgFile    = "data/lu.xlsx"
	switzerlandFile   = "data/ch.xlsx"
	austriaFile       = "data/at.csv"
	liechtensteinFile = "data/li.xlsx"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: goiban-data-loader <src> <dburl>")
		fmt.Println("e.g: goiban-data-loader <bundesbank|nbb|nl|lu|at|ch|li> root:root@/goiban?charset=utf8")
		return
	}

	target := os.Args[1]
	connURL := os.Args[2]

	repo := data.NewSQLStore("mysql", connURL)

	switch target {
	default:
		fmt.Println("unknown target")
	case "bundesbank":
		loader.LoadBundesbankData(bundesbankFile, repo)
	case "nbb":
		loader.LoadBelgiumData(nbbFile, repo)
	case "nl":
		loader.LoadNetherlandsData(netherlandsFile, repo)
	case "lu":
		loader.LoadLuxembourgData(luxembourgFile, repo)
	case "ch":
		loader.LoadSwitzerlandData(switzerlandFile, repo)
	case "li":
		loader.LoadLiechtensteinData(liechtensteinFile, repo)
	case "at":
		loader.LoadAustriaData(austriaFile, repo)
	}
}
