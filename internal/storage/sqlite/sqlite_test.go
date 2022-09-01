package sqlite

import (
	"testing"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/webscraper"
)

func TestSqlite(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		"adding entries into type effectiveness": testTypeEffectivess,
	} {
		t.Run(scenario, f)
	}
}

func testTypeEffectivess(t *testing.T) {
	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon2.db").Build()
	svc := listing.NewService(scraper)
	db.TypeEffectiveNess(svc.TypeEffectiveNess())
}
