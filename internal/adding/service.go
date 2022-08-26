package adding

// Service is our adding service interface
type Service interface {
	EggMoves(pokdexNo int, moves []EggMove)
	Moveset(pokdexNo int, moves []Moveset)
	Moves(moves []Move)
	PokemonType(pokdexNo int, typeName string)
}

// Repository is our respository interface
type Repository interface {
	EggMoves(pokdexNo int, moves []EggMove)
	Moveset(pokdexNo int, moves []Moveset)
	Moves(moves []Move)
	PokemonType(pokdexNo int, typeName string)
}

type service struct {
	r Repository
}

func (s *service) EggMoves(pokdexNo int, moves []EggMove) {
	s.r.EggMoves(pokdexNo, moves)
}

func (s *service) Moveset(pokdexNo int, moves []Moveset) {
	s.r.Moveset(pokdexNo, moves)
}

func (s *service) Moves(moves []Move) {
	s.r.Moves(moves)
}

func (s *service) PokemonType(pokdexNo int, typeName string) {
	s.r.PokemonType(pokdexNo, typeName)
}
