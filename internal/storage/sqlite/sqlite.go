package sqlite

import (
	"database/sql"

	"github.com/manofthelionarmy/pokemondbscrapper-cli/internal/adding"
)

// Sqlite is our repo implementation
type Sqlite struct {
	db *sql.DB
}

// EggMoves adds the eggmoves for a pokemon
func (s *Sqlite) EggMoves(pokdexNo int, moves []adding.EggMove) {
	panic("not implemented") // TODO: Implement
}

// Moveset adds the moveset for a pokemon
func (s *Sqlite) Moveset(pokdexNo int, moves []adding.Moveset) {
	panic("not implemented") // TODO: Implement
}

// Moves populates the moves table
func (s *Sqlite) Moves(moves []adding.Move) {
	panic("not implemented") // TODO: Implement
}

// PokemonType adds an entry into the pokemon_types table
func (s *Sqlite) PokemonType(pokdexNo int, typeName string) {
	panic("not implemented") // TODO: Implement
}

// TypeEffectiveNess adds a type effectiveness entry
func (s *Sqlite) TypeEffectiveNess() {

}
