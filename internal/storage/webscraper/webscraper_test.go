package webscraper

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebScraper(t *testing.T) {
	for scenario, f := range map[string]func(*testing.T){
		"test eggmoves": testEggMoves,
		"all pokemon":   testAllPokemon,
		"all moves":     testAllMoves,
	} {
		t.Run(scenario, f)
	}

}

func testEggMoves(t *testing.T) {
	url := "https://pokemondb.net"
	webScraper := NewBuilder().WithURL(url).Build()

	name := "bulbasaur"
	eggMoves := webScraper.EggMoves(name)
	require.NotZero(t, len(eggMoves))
	log.Println(eggMoves)
}

func testAllPokemon(t *testing.T) {
	url := "https://pokemondb.net"
	webScraper := NewBuilder().WithURL(url).Build()

	allPokemon := webScraper.AllPokemon()
	require.NotZero(t, len(allPokemon))
	log.Println(allPokemon)
}

func testAllMoves(t *testing.T) {
	url := "https://pokemondb.net"
	webScraper := NewBuilder().WithURL(url).Build()

	allMoves := webScraper.AllMoves()
	require.NotZero(t, len(allMoves))
	log.Println(allMoves)
}
