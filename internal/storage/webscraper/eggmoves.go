package webscraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

func eggMovesHandler(moves *[]listing.EggMove) (string, colly.HTMLCallback) {
	return "div.tabset-moves-game", func(div *colly.HTMLElement) {
		tables := div.DOM.Find("table.data-table")
		var movesetTable *goquery.Selection
		tables.Each(func(i int, s *goquery.Selection) {
			// for some reason css selector first-child doesn't work
			// The reason it doesn't work because the tables aren't sibilings
			if i == 1 {
				movesetTable = s
			}
		})

		// Get all the rows in the body
		tr := movesetTable.Find("tbody tr")

		// We are visitng every page serially per pokemon, therefore he next index of the moves
		// is at i + n
		n := len(*moves)

		// The number of moves in a pokemon movesets is equal to the number of rows in tbody
		for range tr.Nodes {
			(*moves) = append(*moves, listing.EggMove{})
		}

		names := tr.Find("td.cell-name")
		names.Each(func(i int, s *goquery.Selection) {
			(*moves)[i+n].Move = s.Text()
		})

		moveTypes := tr.Find("a.type-icon")
		moveTypes.Each(func(i int, s *goquery.Selection) {
			(*moves)[i+n].Type = s.Text()
		})

		categories := tr.Find("td[data-sort-value] img")
		categories.Each(func(i int, s *goquery.Selection) {
			value, _ := s.Attr("title")
			(*moves)[i+n].Category = value
		})

		// power is the 5th column
		power := tr.Find("td.cell-num:nth-child(4)")
		power.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "—" {
				(*moves)[i+n].Power = "0"
			} else {
				(*moves)[i+n].Power = s.Text()
			}
		})

		// accuracy is 6th column
		accuracy := tr.Find("td.cell-num:nth-child(5)")
		accuracy.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "—" {
				(*moves)[i+n].Accuracy = "0"
			} else if s.Text() == "∞" {
				(*moves)[i+n].Accuracy = "Infinity"
			} else {
				(*moves)[i+n].Accuracy = s.Text()
			}
		})
	}
}
