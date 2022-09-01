package listing

// Repository is a storage abstraction
type Repository interface {
	EggMoves(name string) []EggMove
	Moveset(name string) []Moveset
	AllMoves() []Move
	AllPokemon() []Pokemon
	TypeEffectiveNess() []TypeEffectiveNess
}

// Service is a service abstraction that interacts with a Repository
type Service interface {
	EggMoves(name string) []EggMove
	Moveset(name string) []Moveset
	AllPokemon() []Pokemon
	AllMoves() []Move
	TypeEffectiveNess() []TypeEffectiveNess
}

// internal service
type service struct {
	r Repository
}

// NewService takes in a Repository
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) EggMoves(name string) []EggMove {
	// need to find out a way to get the name of the pokemon while getting the eggmoves
	return s.r.EggMoves(name)
}

func (s *service) Moveset(pokemon string) []Moveset {
	return s.r.Moveset(pokemon)
}

func (s *service) AllPokemon() []Pokemon {
	return s.r.AllPokemon()
}

func (s *service) AllMoves() []Move {
	return s.r.AllMoves()
}

func (s *service) TypeEffectiveNess() []TypeEffectiveNess {
	return s.r.TypeEffectiveNess()
}
