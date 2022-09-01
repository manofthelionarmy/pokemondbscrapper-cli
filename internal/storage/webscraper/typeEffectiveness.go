package webscraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
)

const (
	normal = iota
	fighting
	rock
	fire
	poison
	ghost
	water
	ground
	dark
	grass
	flying
	dragon
	electric
	psychic
	steel
	ice
	bug
	fairy
)

var typeMap = map[string]int{
	"normal":   normal,
	"fighting": fighting,
	"rock":     rock,
	"fire":     fire,
	"poison":   poison,
	"ghost":    ghost,
	"water":    water,
	"ground":   ground,
	"dark":     dark,
	"grass":    grass,
	"flying":   flying,
	"dragon":   dragon,
	"electric": electric,
	"psychic":  psychic,
	"steel":    steel,
	"ice":      ice,
	"bug":      bug,
	"fairy":    fairy,
}

var invTypeMap = map[int]string{
	normal:   "normal",
	fighting: "fighting",
	rock:     "rock",
	fire:     "fire",
	poison:   "poison",
	ghost:    "ghost",
	water:    "water",
	ground:   "ground",
	dark:     "dark",
	grass:    "grass",
	flying:   "flying",
	dragon:   "dragon",
	electric: "electric",
	psychic:  "psychic",
	steel:    "steel",
	ice:      "ice",
	bug:      "bug",
	fairy:    "fairy",
}

func typeEffectivenessHandler(typeEffectiveness *[]listing.TypeEffectiveNess) (string, colly.HTMLCallback) {
	return ".type-table", func(table *colly.HTMLElement) {
		matrix := [18][18]string{}
		rows := table.DOM.Find("tbody > tr")
		rows.Each(func(i int, s *goquery.Selection) {
			header := s.Find("th > a")
			row := typeMap[strings.ToLower(header.Text())]
			td := s.Find("td")
			td.Each(func(i int, s *goquery.Selection) {
				title, _ := s.Attr("title")
				againstType := strings.Split(title, " ")[2]
				againstType = strings.ToLower(againstType)

				col := typeMap[strings.TrimSpace(againstType)]

				idx := strings.Index(title, "=")
				effectiveNess := title[idx+2:]
				matrix[row][col] = effectiveNess
			})
		})
		for row := range matrix {
			for col := range matrix[row] {
				tE := listing.TypeEffectiveNess{}
				tE.TypeName = invTypeMap[row]
				tE.AgainstType = invTypeMap[col]
				effectiveNess := matrix[row][col]
				tE.AttackScore = getAttackScore(effectiveNess)
				tE.DefenseScore = getDefenseScore(matrix[col][row])
				(*typeEffectiveness) = append((*typeEffectiveness), tE)
			}
		}
	}
}

func getAttackScore(effectiveNess string) int {
	switch effectiveNess {
	case "normal effectiveness":
		return listing.Effective
	case "not very effective":
		return listing.NotVeryEffective
	case "super-effective":
		return listing.SuperEffective
	case "no effect":
		return listing.NoEffect
	default:
		return -100
	}
}

func getDefenseScore(effectiveNess string) int {
	switch effectiveNess {
	case "normal effectiveness":
		return listing.Effective
	case "not very effective":
		return listing.NotVeryEffective
	case "super-effective":
		return listing.SuperWeak
	case "no effect":
		return listing.NoEffect
	default:
		return -100
	}
}
