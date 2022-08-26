package webscraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

var _ listing.Repository = (*WebScraper)(nil)

// WebScraper is our repository implementation
type WebScraper struct {
	url string
}

// EggMoves uses colly to scrape a webpage
func (w *WebScraper) EggMoves(name string) []listing.EggMove {
	scraper := colly.NewCollector(
		colly.MaxDepth(2),
	)
	moves := make([]listing.EggMove, 0)
	scraper.OnHTML(eggMovesHandler(&moves))
	scraper.Visit(fmt.Sprintf("%s/pokedex/%s", w.url, name))
	scraper.Wait()
	return moves
}

// AllPokemon tells the webscraper to get all pokemon
func (w *WebScraper) AllPokemon() []listing.Pokemon {
	scraper := colly.NewCollector(
		colly.MaxDepth(2),
	)
	pokemon := make([]listing.Pokemon, 0)
	scraper.OnHTML(allPokemonHandler(&pokemon))
	scraper.Visit(fmt.Sprintf("%s/%s", w.url, "pokedex/all"))
	scraper.Wait()
	return pokemon
}

// AllMoves returns  all the moves in-game
func (w *WebScraper) AllMoves() []listing.Move {
	scraper := colly.NewCollector(
		colly.MaxDepth(2),
	)
	moves := make([]listing.Move, 0)
	scraper.OnHTML(allMovesHandler(&moves))
	scraper.Visit(fmt.Sprintf("%s/%s", w.url, "move/all"))
	scraper.Wait()
	return moves
}

// Moveset gets all the moves for a particular pokemon
func (w *WebScraper) Moveset(pokemon string) []listing.Moveset {
	scraper := colly.NewCollector(
		colly.MaxDepth(2),
	)
	moves := make([]listing.Moveset, 0)
	scraper.OnHTML(movesetHandler(&moves))
	scraper.Visit(fmt.Sprintf("%s/pokedex/%s", w.url, pokemon))
	scraper.Wait()
	return moves
}
