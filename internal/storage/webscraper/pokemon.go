package webscraper

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

func allPokemonHandler(pokemon *[]listing.Pokemon) (string, colly.HTMLCallback) {
	return "table[id=pokedex] tbody tr", func(tr *colly.HTMLElement) {
		// infocard contains info about pokedex number
		*pokemon = append(*pokemon, listing.Pokemon{})
		n := len(*pokemon) - 1

		infocard := tr.DOM.Find("span.infocard-cell-data")
		(*pokemon)[n].PokedexNo, _ = strconv.Atoi(infocard.Text())

		// are these concurrent?
		cellName := tr.DOM.Find("td.cell-name > a")

		// what does this index refer to?
		(*pokemon)[n].Name = cellName.Text()

		// check if it's an other pokemon form
		pokemonForm := tr.DOM.Find("td.cell-name > small")
		if len(pokemonForm.Text()) > 0 {
			(*pokemon)[n].Name = pokemonForm.Text()
		}

		types := tr.DOM.Find("td > a.type-icon")
		// This isn't split into one type :(
		(*pokemon)[n].Types = make([]string, 0)
		types.Each(func(_ int, s *goquery.Selection) {
			(*pokemon)[n].Types = append((*pokemon)[n].Types, s.Text())
		})

		totalBaseStats := tr.DOM.Find("td.cell-total")
		(*pokemon)[n].TotalBaseStats, _ = strconv.Atoi(totalBaseStats.Text())

		stats := tr.DOM.Find("td.cell-num")

		// 6 stats per pokemon
		stats.Each(func(statIndex int, s *goquery.Selection) {
			switch statIndex {
			case 0:
				(*pokemon)[n].Hp, _ = strconv.Atoi(s.Text())
			case 1:
				(*pokemon)[n].Attack, _ = strconv.Atoi(s.Text())
			case 2:
				(*pokemon)[n].Defense, _ = strconv.Atoi(s.Text())
			case 3:
				(*pokemon)[n].SpAttack, _ = strconv.Atoi(s.Text())
			case 4:
				(*pokemon)[n].SpDefense, _ = strconv.Atoi(s.Text())
			case 5:
				(*pokemon)[n].Speed, _ = strconv.Atoi(s.Text())
			}
		})
	}
}
