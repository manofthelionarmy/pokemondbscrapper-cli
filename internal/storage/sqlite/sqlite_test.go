package sqlite

import (
	"testing"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/storage/webscraper"
)

func TestSqlite(t *testing.T) {
	for scenario, f := range map[string]func(t *testing.T){
		// "adding entries into type effectiveness": testTypeEffectivess,
		// "adding entries into moves table": testMoves,
		// "adding pokemon entries into pokemon table":  testPokemon,
		// "adding eggmove entries into eggmoves table": testEggmoves,
		// "adding moveset entries into movset table": testMovset,
		"adding pokemon type entries into pokemon_types table": testPokemonTypes,
	} {
		t.Run(scenario, f)
	}
}

func testTypeEffectivess(t *testing.T) {
	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)
	db.TypeEffectiveNess(svc.TypeEffectiveNess())
}

func testMoves(t *testing.T) {
	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)
	db.Moves(svc.AllMoves())
}

func testPokemon(t *testing.T) {
	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)
	db.Pokemon(svc.AllPokemon())
}

func testEggmoves(t *testing.T) {
	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)

	pokemon := svc.AllPokemon()
	for _, p := range pokemon {
		eggmoves := svc.EggMoves(p.Name)
		db.EggMoves(p.PokedexNo, eggmoves)
	}
}

func testMovset(t *testing.T) {

	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)

	pokemon := svc.AllPokemon()
	for _, p := range pokemon {
		moves := svc.Moveset(p.Name)
		db.Moveset(p.PokedexNo, moves)
	}
}

// TODO: make these better testable
func testPokemonTypes(t *testing.T) {

	scraper := webscraper.NewBuilder().WithURL("https://pokemondb.net").Build()
	db := NewBuilder().WithDataSource("pokemon.db").Build()
	svc := listing.NewService(scraper)

	pokemon := svc.AllPokemon()
	db.PokemonType(pokemon)
}
