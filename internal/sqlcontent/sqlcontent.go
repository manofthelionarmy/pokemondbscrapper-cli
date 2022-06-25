package sqlcontent

import (
	"fmt"
	"os"
	"reflect"

	"github.com/manofthelionarmy/pokemondbscrapper/pkg/listing"
)

// SQLContent is responsbile for converting our scrapped data to SQLContnet. Wraps bytes.Buffer
type SQLContent struct {
	*SQLFile
}

// NewSQLContent initiates our sql content writer
func NewSQLContent(f *os.File) *SQLContent {
	return &SQLContent{
		SQLFile: NewSQLFILE(f),
	}
}

// CreateTable creates a table
func (sc *SQLContent) CreateTable(name string) error {
	var tableStatement string
	switch name {
	case "pokemon":
		tableStatement = pokemonTableStatement()
	case "moves":
		tableStatement = movesTableStatement()
	default:
		return fmt.Errorf("%s is not a supported table", name)
	}
	_, err := sc.SQLFile.Write([]byte(tableStatement))
	if err != nil {
		return err
	}

	return nil
}

// InsertInto creates InsertInto statement
func (sc *SQLContent) InsertInto(table string, pokemonInfo interface{}, opts ...Options) error {
	var pokemonName string
	if len(opts) != 0 {
		pokemonName = opts[0].PokemonName
	}

	switch pokemonInfo.(type) {
	case []listing.Pokemon:
		pokemon, _ := pokemonInfo.([]listing.Pokemon)
		for _, mon := range pokemon {
			var pokemonType1 string
			var pokemonType2 string
			pokemonType1 = mon.Type[0]
			pokemonType2 = pokemonType1
			if len(mon.Type) == 2 {
				pokemonType2 = mon.Type[1]
			}
			// TODO: handle region variants
			sc.SQLFile.Write([]byte(
				"INSERT INTO " + table +
					" (pokedexNo, name, type1, type2, totalBaseStats, hp, attack, defense, spAtk, spDef, speed) " +
					fmt.Sprintf(
						"VALUES (%s, \"%s\", \"%s\", \"%s\", %s, %s, %s, %s, %s, %s);\n\n",
						mon.PokedexNo, mon.Name, pokemonType1, pokemonType2, mon.TotalBaseStats, mon.HP,
						mon.Attack, mon.SpAtk, mon.SpDef, mon.Speed,
					)),
			)
		}
	case []listing.Move:
		moves, _ := pokemonInfo.([]listing.Move)
		for _, move := range moves {
			sc.SQLFile.Write([]byte(
				"INSERT INTO " + table +
					" (name, type, category, power, accuracy, effect) " +
					fmt.Sprintf("VALUES (\"%s\", \"%s\", \"%s\", \"%s\", \"%s\", \"%s\");\n\n",
						move.Name, move.Type, move.Category, move.Power, move.Accuracy, move.Effect,
					)),
			)
		}
	case []listing.EggMove:
		eggMoves, _ := pokemonInfo.([]listing.EggMove)
		if len(opts) == 0 {
			return fmt.Errorf("expecting a pokemon name to be passed in for a moveset")
		}
		for _, move := range eggMoves {
			sc.SQLFile.Write([]byte(
				//FIXME: category can be empty for gmax moves
				"INSERT INTO " + table +
					" (pokemon, name, type, category, power, accuracy)" +
					fmt.Sprintf("VALUES (%s, %s, %s, %s, %s, %s);\n\n",
						pokemonName, move.Move, move.Type, move.Category, move.Power, move.Accuracy,
					)),
			)
		}
	case []listing.Moveset:
		moveset, _ := pokemonInfo.([]listing.Moveset)
		if len(opts) == 0 {
			return fmt.Errorf("expecting a pokemon name to be passed in for a moveset")
		}

		// NOTE: I can always fix this later
		// FIXME: duplication of data, we need to find the entry/id of this move
		// We can always use the name (they are unique move names anyway)
		for _, move := range moveset {
			sc.SQLFile.Write([]byte(
				"INSERT INTO " + table +
					" (pokemon, name, level, type, category, power, accuracy)" +
					fmt.Sprintf("VALUES (%s, %s, %s, %s, %s, %s, %s);\n\n",
						pokemonName, move.Level, move.Move, move.Type, move.Category, move.Power, move.Accuracy,
					)),
			)
		}
	default:
		return fmt.Errorf(
			"unsupported type, can only support %+v",
			[]reflect.Type{
				reflect.TypeOf([]listing.Pokemon{}),
				reflect.TypeOf([]listing.Move{}),
				reflect.TypeOf([]listing.EggMove{}),
				reflect.TypeOf([]listing.Moveset{}),
			},
		)
	}
	return nil
}

// Flush flushes the content from the buffered stream
func (sc *SQLContent) Flush() error {
	return sc.SQLFile.Flush()
}

func pokemonTableStatement() string {
	return "CREATE TABLE pokemon (\n" +
		"\tpokedexNo INTEGER,\n" +
		"\tname VARCHAR,\n" +
		"\ttype1 VARCHAR,\n" +
		"\ttype2 VARCHAR,\n" +
		"\ttotalBaseStats INTEGER,\n" +
		"\thp INTEGER,\n" +
		"\tattack INTEGER,\n" +
		"\tdefense INTEGER,\n" +
		"\tspAtk INTEGER,\n" +
		"\tspDef INTEGER,\n" +
		"\tspeed INTEGER);\n\n"
}

func movesTableStatement() string {
	return "CREATE TABLE moves (\n" +
		"\tname VARCHAR,\n" +
		"\ttype VARCHAR,\n" +
		"\tcategory VARCHAR,\n" +
		"\tpower VARCHAR,\n" +
		"\taccuracy INTEGER,\n" +
		"\teffect VARCHAR);\n\n"
}
