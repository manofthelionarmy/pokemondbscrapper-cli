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
		FOREIGN KEY(moves_id) REFERENCES moves(moves_id)
	);`)
	if err != nil {
		panic(err)
	}
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	// we'll get more than one pokemon for a pokedex no
	// I still need to populate this, so I need to get the pokemon_id by using the pokdexNo
	// screw it, we'll enter the same stuff for a pokedexNo for now
	pokemonQueryStmt, err := s.db.Prepare(`select pokemon_id from pokemon where pokedexNo = $1;`)
	if err != nil {
		panic(err)
	}

	rows, _ := pokemonQueryStmt.Query(pokdexNo)
	pokemonIDs := make([]int, 0)
	for rows.Next() {
		var pokemonID int
		if err := rows.Scan(&pokemonID); err != nil {
			panic(err)
		}
		pokemonIDs = append(pokemonIDs, pokemonID)
	}

	movesQueryStmt, err := s.db.Prepare(`select moves_id from moves where name = $1;`)
	if err != nil {
		panic(err)
	}

	movesIDs := make([]int, 0)
	for _, move := range moves {
		// get the move id by getting the name of the move
		rows, err := movesQueryStmt.Query(move.Move)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var moveID int
			if err := rows.Scan(&moveID); err != nil {
				panic(err)
			}
			movesIDs = append(movesIDs, moveID)
		}
	}

	insertStmt, err := s.db.Prepare(`insert into eggmoves (
		pokemon_id,
		moves_id
	) values ($1, $2);`)

	if err != nil {
		panic(err)
	}

	for _, moveID := range movesIDs {
		for _, pokemonID := range pokemonIDs {
			if _, err := insertStmt.Exec(pokemonID, moveID); err != nil {
				panic(err)
			}
		}
	}
}

// Moveset adds the moveset for a pokemon
func (s *Sqlite) Moveset(pokdexNo int, moves []listing.Moveset) {
	stmt, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS moveset(
		moveset_id INTEGER PRIMARY KEY,
		level INTEGER,
		pokemon_id INTEGER,
		moves_id INTEGER,
		FOREIGN KEY(pokemon_id) REFERENCES pokemon(pokemon_id),
		FOREIGN KEY(moves_id) REFERENCES moves(moves_id)
	);`)
	if err != nil {
		panic(err)
	}
	if _, err := stmt.Exec(); err != nil {
		panic(err)
	}
	// we'll get more than one pokemon for a pokedex no
	// I still need to populate this, so I need to get the pokemon_id by using the pokdexNo
	// screw it, we'll enter the same stuff for a pokedexNo for now
	pokemonQueryStmt, err := s.db.Prepare(`select pokemon_id from pokemon where pokedexNo = $1;`)
	if err != nil {
		panic(err)
	}

	rows, _ := pokemonQueryStmt.Query(pokdexNo)
	pokemonIDs := make([]int, 0)
	for rows.Next() {
		var pokemonID int
		if err := rows.Scan(&pokemonID); err != nil {
			panic(err)
		}
		pokemonIDs = append(pokemonIDs, pokemonID)
	}

	movesQueryStmt, err := s.db.Prepare(`select moves_id from moves where name = $1;`)
	if err != nil {
		panic(err)
	}

	movesIDs := make([]int, 0)
	for _, move := range moves {
		// get the move id by getting the name of the move
		rows, err := movesQueryStmt.Query(move.Name)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var moveID int
			if err := rows.Scan(&moveID); err != nil {
				panic(err)
			}
			movesIDs = append(movesIDs, moveID)
		}
	}

	insertStmt, err := s.db.Prepare(`insert into moveset (
		level,
		pokemon_id,
		moves_id
	) values ($1, $2, $3);`)

	if err != nil {
		panic(err)
	}

	for i, moveID := range movesIDs {
		for _, pokemonID := range pokemonIDs {
			if _, err := insertStmt.Exec(moves[i].Level, pokemonID, moveID); err != nil {
				panic(err)
			}
		}
	}
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
