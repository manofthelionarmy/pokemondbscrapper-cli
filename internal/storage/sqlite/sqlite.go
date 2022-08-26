package sqlite

import (
	"database/sql"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/adding"
	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"
	_ "github.com/mattn/go-sqlite3" // import go-sqlite3 driver
)

var _ adding.Repository = (*Sqlite)(nil)

// Sqlite is our repo implementation
type Sqlite struct {
	db *sql.DB
}

// EggMoves adds the eggmoves for a pokemon
func (s *Sqlite) EggMoves(pokdexNo int, moves []listing.EggMove) {
	panic("not implemented") // TODO: Implement
}

// Moveset adds the moveset for a pokemon
func (s *Sqlite) Moveset(pokdexNo int, moves []listing.Moveset) {
	panic("not implemented") // TODO: Implement
}

// Moves populates the moves table
func (s *Sqlite) Moves(moves []listing.Move) {
	panic("not implemented") // TODO: Implement
}

// PokemonType adds an entry into the pokemon_types table
func (s *Sqlite) PokemonType(pokdexNo int, typeNames ...string) {
	panic("not implemented") // TODO: Implement
}

// TypeEffectiveNess adds all type effectiveness entries and creates table if it doesn't exist
func (s *Sqlite) TypeEffectiveNess() {
	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS type_effectiveness(
		type_effectiveness_id INTEGER PRIMARY KEY,
		type_name VARCHAR(255),
		against_type VARCHAR(255),
		attack_score INTEGER,
		defense_score INTEGER
	);`)
	if err != nil {
		panic(err)
	}
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}

	if !checkTypeEntriesAdded(s.db, adding.Grass) {
		typeEffectivenessChart := make([]adding.TypeEffectiveNess, 0)
		typeEffectivenessChart = appendGrassType(typeEffectivenessChart)

		for _, entry := range typeEffectivenessChart {
			stmt, err := s.db.Prepare(`INSERT INTO type_effectiveness (type_name, against_type, attack_score, defense_score)
			VALUES ($1, $2, $3, $4);
		`)
			if err != nil {
				panic(err)
			}
			if _, err := stmt.Exec(entry.TypeName, entry.AgainstType, entry.AttackScore, entry.DefenseScore); err != nil {
				panic(err)
			}
		}
	}
	if !checkTypeEntriesAdded(s.db, adding.Normal) {
		typeEffectivenessChart := make([]adding.TypeEffectiveNess, 0)
		typeEffectivenessChart = appendNormalType(typeEffectivenessChart)

		for _, entry := range typeEffectivenessChart {
			stmt, err := s.db.Prepare(`INSERT INTO type_effectiveness (type_name, against_type, attack_score, defense_score)
			VALUES ($1, $2, $3, $4);
		`)
			if err != nil {
				panic(err)
			}
			if _, err := stmt.Exec(entry.TypeName, entry.AgainstType, entry.AttackScore, entry.DefenseScore); err != nil {
				panic(err)
			}
		}
	}
	if !checkTypeEntriesAdded(s.db, adding.Fire) {
		typeEffectivenessChart := make([]adding.TypeEffectiveNess, 0)
		typeEffectivenessChart = appendFireType(typeEffectivenessChart)

		for _, entry := range typeEffectivenessChart {
			stmt, err := s.db.Prepare(`INSERT INTO type_effectiveness (type_name, against_type, attack_score, defense_score)
			VALUES ($1, $2, $3, $4);
		`)
			if err != nil {
				panic(err)
			}
			if _, err := stmt.Exec(entry.TypeName, entry.AgainstType, entry.AttackScore, entry.DefenseScore); err != nil {
				panic(err)
			}
		}
	}
}

// Pokemon adds all of the pokemon entries.
func (s *Sqlite) Pokemon(pokemon []listing.Pokemon) {
}
