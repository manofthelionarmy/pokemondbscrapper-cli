package adding

import "github.com/manofthelionarmy/pokemondbscrapper-cli/internal/listing"

// Service is our adding service interface
type Service interface {
	EggMoves(pokdexNo int, moves []listing.EggMove)
	Moveset(pokdexNo int, moves []listing.Moveset)
	Moves(moves []listing.Move)
	Pokemon(pokemon []listing.Pokemon)
	PokemonType([]listing.Pokemon)
	TypeEffectiveNess([]listing.TypeEffectiveNess)
}

// Repository is our respository interface
type Repository interface {
	EggMoves(pokdexNo int, moves []listing.EggMove)
	Moveset(pokdexNo int, moves []listing.Moveset)
	Moves(moves []listing.Move)
	Pokemon(pokemon []listing.Pokemon)
	PokemonType([]listing.Pokemon)
	TypeEffectiveNess([]listing.TypeEffectiveNess)
}

type service struct {
	r Repository
}

// NewService returns a new adding Service
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) EggMoves(pokdexNo int, moves []listing.EggMove) {
	s.r.EggMoves(pokdexNo, moves)
}

func (s *service) Moveset(pokdexNo int, moves []listing.Moveset) {
	s.r.Moveset(pokdexNo, moves)
}

func (s *service) Moves(moves []listing.Move) {
	s.r.Moves(moves)
}

func (s *service) PokemonType(pokemon []listing.Pokemon) {
	s.r.PokemonType(pokemon)
}

func (s *service) Pokemon(pokemon []listing.Pokemon) {
	s.r.Pokemon(pokemon)
}

func (s *service) TypeEffectiveNess(typeEffectiveness []listing.TypeEffectiveNess) {
	s.r.TypeEffectiveNess(typeEffectiveness)
}
