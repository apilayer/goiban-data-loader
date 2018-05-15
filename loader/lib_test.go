package loader_test

import (
	"testing"

	"github.com/fourcube/goiban-data"

	"github.com/fourcube/goiban-data-loader/loader"
)

func TestLoadBundesbank(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBundesbankData("../data/bundesbank.txt", repo)
}

func TestLoadBundesbankFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBundesbankData(loader.DefaultBundesbankPath(), repo)
}

func TestLoadAustriaFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadAustriaData(loader.DefaultAustriaPath(), repo)
}

func TestLoadSwitzerlandFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadSwitzerlandData(loader.DefaultSwitzerlandPath(), repo)
}

func TestLoadLiechtensteinFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLiechtensteinData(loader.DefaultLiechtensteinPath(), repo)
}

func TestLoadLuxembourgFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadLuxembourgData(loader.DefaultLuxembourgPath(), repo)
}

func TestLoadNetherlandsFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadNetherlandsData(loader.DefaultNetherlandsPath(), repo)
}

func TestLoadBelgiumFromDefaultPath(t *testing.T) {
	repo := data.NewInMemoryStore()
	loader.LoadBelgiumData(loader.DefaultBelgiumPath(), repo)
}
