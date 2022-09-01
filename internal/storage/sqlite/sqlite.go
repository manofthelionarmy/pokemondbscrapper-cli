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
	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS eggmoves(
		eggmoves_id INTEGER PRIMARY KEY,
		pokemon_id INTEGER,
		moves_id INTEGER,
		FOREIGN KEY(pokemon_id) REFERENCES pokemon(pokemon_id),
		FOREIGN KEY(moves_id) REFERENCES moves(moves_id),
	);`)
	if err != nil {
		panic(err)
	}
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	// I still need to populate this, so I need to get the pokemon_id by using the pokdexNo
}

// Moveset adds the moveset for a pokemon
func (s *Sqlite) Moveset(pokdexNo int, moves []listing.Moveset) {
	panic("not implemented") // TODO: Implement
}

// Moves populates the moves table
func (s *Sqlite) Moves(moves []listing.Move) {
	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS moves (
		moves_id INTEGER PRIMARY KEY,
		name VARCHAR(255),
		type VARCHAR(255),
		power INTEGER,
		category VARCHAR(255),
		accuracy INTEGER,
		power_points INTEGER,
		effect TEXT
	);`)
	if err != nil {
		panic(err)
	}
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}

	for _, move := range moves {
		insertStmt, err := s.db.Prepare(`INSERT INTO moves (name, type, power, category, accuracy, power_points, effect)
			VALUES ($1, $2, $3, $4, $5, $6, $7);
		`)
		if err != nil {
			panic(err)
		}
		if _, err := insertStmt.Exec(
			move.Name, move.Type, move.Power,
			move.Category, move.Accuracy,
			move.PowerPoints, move.Effect); err != nil {
			panic(err)
		}
	}
}

// PokemonType adds an entry into the pokemon_types table
func (s *Sqlite) PokemonType(pokdexNo int, typeNames ...string) {
	panic("not implemented") // TODO: Implement
}

// TypeEffectiveNess adds all type effectiveness entries and creates table if it doesn't exist
func (s *Sqlite) TypeEffectiveNess(typeEffectiveness []listing.TypeEffectiveNess) {
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
	for _, entry := range typeEffectiveness {
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

// Pokemon adds all of the pokemon entries.
func (s *Sqlite) Pokemon(pokemon []listing.Pokemon) {
	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS pokemon (
		pokemon_id INTEGER PRIMARY KEY,
		pokedexNo INTEGER,
		name VARCHAR(255),
		hp INTEGER,
		attack INTEGER,
		defense INTEGER,
		spAttack INTEGER,
		spDefense INTEGER,
		speed INTEGER
	);`)
	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}

	insertStmt, err := s.db.Prepare(`INSERT INTO pokemon (pokedexNo, name, hp, attack, defense, spAttack, spDefense, speed)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8);
	`)
	if err != nil {
		panic(err)
	}
	for _, p := range pokemon {
		if _, err := insertStmt.Exec(
			p.PokedexNo, p.Name,
			p.Hp, p.Attack,
			p.Defense, p.SpAttack,
			p.SpDefense, p.Speed,
		); err != nil {
			panic(err)
		}
	}
}
