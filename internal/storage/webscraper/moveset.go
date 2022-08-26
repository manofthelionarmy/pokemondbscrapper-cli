package webscraper

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

func movesetHandler(moves *[]listing.Moveset) (string, colly.HTMLCallback) {
	return "div.tabset-moves-game", func(div *colly.HTMLElement) {
		tables := div.DOM.Find("table.data-table")

		var movesetTable *goquery.Selection
		tables.Each(func(i int, s *goquery.Selection) {
			// for some reason css selector first-child doesn't work
			if i == 0 {
				movesetTable = s
			}
		})

		// Get all the rows in the body
		tr := movesetTable.Find("tbody tr")

		// The number of moves in a pokemon movesets is equal to the number of rows in tbody
		for range tr.Nodes {
			(*moves) = append(*moves, listing.Moveset{})
		}

		// table cells that hold the data reprensenting the pokemon's data is the first of 3
		levels := tr.Find("td.cell-num:first-child")
		levels.Each(func(i int, s *goquery.Selection) {
			(*moves)[i].Level, _ = strconv.Atoi(s.Text())
		})

		names := tr.Find("td.cell-name")
		names.Each(func(i int, s *goquery.Selection) {
			(*moves)[i].Name = s.Text()
		})

		moveTypes := tr.Find("a.type-icon")
		moveTypes.Each(func(i int, s *goquery.Selection) {
			(*moves)[i].Type = s.Text()
		})

		categories := tr.Find("td[data-sort-value] img")
		categories.Each(func(i int, s *goquery.Selection) {
			value, _ := s.Attr("title")
			(*moves)[i].Category = value
		})

		// power is the 5th column
		power := tr.Find("td.cell-num:nth-child(5)")
		power.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "—" {
				(*moves)[i].Power = 0
			} else {
				(*moves)[i].Power, _ = strconv.Atoi(s.Text())
			}
		})

		// accuracy is 6th column
		accuracy := tr.Find("td.cell-num:nth-child(6)")
		accuracy.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "—" {
				(*moves)[i].Accuracy = 0
			} else {
				(*moves)[i].Accuracy, _ = strconv.Atoi(s.Text())
			}
		})
	}
}
