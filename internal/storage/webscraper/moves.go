package webscraper

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

const (
	// Infinity means a move never misses
	Infinity = iota - 1
)

func allMovesHandler(moves *[]listing.Move) (string, colly.HTMLCallback) {
	return "table tbody tr", func(tr *colly.HTMLElement) {
		*moves = append(*moves, listing.Move{})

		n := len(*moves) - 1

		name := tr.DOM.Find("td.cell-name").Text()

		(*moves)[n].Name = name

		moveType := tr.DOM.Find("a.type-icon").Text()
		(*moves)[n].Type = moveType

		category, exists := tr.DOM.Find("img.img-fixed").Attr("title")

		if exists {
			(*moves)[n].Category = category
		}

		stats := tr.DOM.Find("td.cell-num")
		stats.Each(func(statsIndex int, s *goquery.Selection) {
			switch statsIndex {
			case 0:
				if s.Text() == "—" {
					(*moves)[n].Power = 0
				} else {
					(*moves)[n].Power, _ = strconv.Atoi(s.Text())
				}
			case 1:
				if s.Text() == "—" {
					(*moves)[n].Accuracy = 0
				} else if s.Text() == "∞" {
					(*moves)[n].Accuracy = Infinity
				} else {
					(*moves)[n].Accuracy, _ = strconv.Atoi(s.Text())
				}
			case 2:
				(*moves)[n].PowerPoints, _ = strconv.Atoi(s.Text())
			}
		})

		effect := tr.DOM.Find("td.cell-long-text").Text()
		if len(effect) == 0 {
			(*moves)[n].Effect = "No Effect"
		} else {
			(*moves)[n].Effect = effect
		}
	}
}
