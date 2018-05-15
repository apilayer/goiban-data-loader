package loader

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fourcube/goiban"
	"github.com/fourcube/goiban-data"
	co "github.com/fourcube/goiban/countries"
)

// LoadBundesbankData loads data from a text file and stores it in
// a bank data repository
func LoadBundesbankData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.BundesbankFileEntry{}, ch)

	source := "German Bundesbank"

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	for entry := range ch {
		bbEntry := entry.(*co.BundesbankFileEntry)

		bic := strings.Replace(bbEntry.Bic, " ", "", -1)

		if bbEntry.M == 1 {
			bankInfo := data.BankInfo{
				Bankcode:  bbEntry.Bankcode,
				Name:      bbEntry.Name,
				Zip:       bbEntry.Zip,
				City:      bbEntry.City,
				Bic:       bic,
				CheckAlgo: bbEntry.CheckAlgo,
				Country:   "DE",
				Source:    source,
			}
			ok, err := repo.Store(bankInfo)

			if err == nil && ok {
				rows++
			} else {
				log.Fatal(err, bankInfo)
			}
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadAustriaData loads data from a text file and stores it in
// a bank data repository
func LoadAustriaData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.AustriaBankFileEntry{}, ch)

	source := "AT"
	uniqueEntries := map[string]co.AustriaBankFileEntry{}

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

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

	for _, entry := range uniqueEntries {
		bankInfo := data.BankInfo{
			Bankcode:  entry.Bankcode,
			Name:      entry.Name,
			Zip:       "",
			City:      "",
			Bic:       entry.Bic,
			CheckAlgo: "",
			Country:   "AT",
			Source:    source,
		}

		_, err := repo.Store(bankInfo)
		if err == nil {
			rows++
		} else {
			log.Fatal(err, entry)
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadSwitzerlandData loads data from a text file and stores it in
// a bank data repository
func LoadSwitzerlandData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.SwitzerlandFileEntry{}, ch)

	source := "CH"
	uniqueEntries := map[string]co.SwitzerlandFileEntry{}

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	for entry := range ch {
		chEntry := entry.(co.SwitzerlandFileEntry)

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

	for _, entry := range uniqueEntries {
		bankInfo := data.BankInfo{
			Bankcode:  entry.Bankcode,
			Name:      entry.Name,
			Zip:       "",
			City:      "",
			Bic:       entry.Bic,
			CheckAlgo: "",
			Country:   "CH",
			Source:    source,
		}

		_, err := repo.Store(bankInfo)
		if err == nil {
			rows++
		} else {
			log.Fatal(err, entry)
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadLiechtensteinData loads data from a text file and stores it in
// a bank data repository
func LoadLiechtensteinData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.LiechtensteinFileEntry{}, ch)

	source := "LI"
	uniqueEntries := map[string]co.LiechtensteinFileEntry{}
	ok, err := repo.Clear(source)

	if err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	log.Printf("Deleted %v entries.", ok)

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

	for _, entry := range uniqueEntries {
		bankInfo := data.BankInfo{
			Bankcode:  entry.Bankcode,
			Name:      entry.Name,
			Zip:       "",
			City:      "",
			Bic:       entry.Bic,
			CheckAlgo: "",
			Country:   "LI",
			Source:    source,
		}

		_, err := repo.Store(bankInfo)
		if err == nil {
			rows++
		} else {
			log.Fatal(err, entry)
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadLuxembourgData loads data from a text file and stores it in
// a bank data repository
func LoadLuxembourgData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.LuxembourgFileEntry{}, ch)

	source := "LU"
	uniqueEntries := map[string]co.LuxembourgFileEntry{}

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	for entry := range ch {
		chEntry := entry.(co.LuxembourgFileEntry)

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

	for _, entry := range uniqueEntries {
		bankInfo := data.BankInfo{
			Bankcode:  entry.Bankcode,
			Name:      entry.Name,
			Zip:       "",
			City:      "",
			Bic:       entry.Bic,
			CheckAlgo: "",
			Country:   "LU",
			Source:    source,
		}

		_, err := repo.Store(bankInfo)
		if err == nil {
			rows++
		} else {
			log.Fatal(err, entry)
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadNetherlandsData loads data from a text file and stores it in
// a bank data repository
func LoadNetherlandsData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.NetherlandsFileEntry{}, ch)

	source := "NL"

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	for entry := range ch {
		entry := entry.(co.NetherlandsFileEntry)
		entry.Bic = strings.Replace(entry.Bic, " ", "", -1)

		bankInfo := data.BankInfo{
			Bankcode:  entry.Bankcode,
			Name:      entry.Name,
			Zip:       "",
			City:      "",
			Bic:       entry.Bic,
			CheckAlgo: "",
			Country:   "NL",
			Source:    source,
		}

		_, err := repo.Store(bankInfo)
		if err == nil {
			rows++
		} else {
			log.Fatal(err, entry)
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

// LoadBelgiumData loads data from a text file and stores it in
// a bank data repository
func LoadBelgiumData(filePath string, repo data.BankDataRepository) {
	ch := make(chan interface{})
	rows := 0
	go goiban.ReadFileToEntries(filePath, &co.BelgiumFileEntry{}, ch)

	source := "NBB"

	if ok, err := repo.Clear(source); err != nil || ok < 0 {
		log.Fatalf("Failed to clear entries for %s: %v", source, err)
	}

	for entry := range ch {
		entries := entry.([]co.BelgiumFileEntry)
		for _, entry := range entries {
			entry.Bic = strings.Replace(entry.Bic, " ", "", -1)

			bankInfo := data.BankInfo{
				Bankcode:  entry.Bankcode,
				Name:      entry.Name,
				Zip:       "",
				City:      "",
				Bic:       entry.Bic,
				CheckAlgo: "",
				Country:   "BE",
				Source:    source,
			}

			_, err := repo.Store(bankInfo)
			if err == nil {
				rows++
			} else {
				log.Fatal(err, entry)
			}
		}
	}

	log.Printf("Loaded %d rows from %s", rows, source)
}

var defaultBasePath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "fourcube", "goiban-data-loader", "data")

func SetBasePath(path string) {
	defaultBasePath = path
}

func DefaultDataPath(basePath, source string) string {
	return filepath.Join(basePath, source)
}

func DefaultBundesbankPath() string {
	return DefaultDataPath(defaultBasePath, "bundesbank.txt")
}

func DefaultAustriaPath() string {
	return DefaultDataPath(defaultBasePath, "at.csv")
}

func DefaultLiechtensteinPath() string {
	return DefaultDataPath(defaultBasePath, "li.xlsx")
}

func DefaultSwitzerlandPath() string {
	return DefaultDataPath(defaultBasePath, "ch.xlsx")
}

func DefaultLuxembourgPath() string {
	return DefaultDataPath(defaultBasePath, "lu.xlsx")
}

func DefaultNetherlandsPath() string {
	return DefaultDataPath(defaultBasePath, "nl.xlsx")
}

func DefaultBelgiumPath() string {
	return DefaultDataPath(defaultBasePath, "nbb.xlsx")
}
