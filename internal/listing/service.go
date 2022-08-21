package listing

// Repository is a storage abstraction
type Repository interface {
	GetEggMoves(name string) []EggMove
	GetMoveset()
	GetPokemon()
	GetPokemonTypes(pokdexNo int)
	GetPokemonTypeEffectiveNess(name string)
}

// Service is a service abstraction that interacts with a Repository
type Service interface {
	GetEggMoves(name string) []EggMove
	GetMoveset()
	GetPokemon()
	GetPokemonTypes(pokdexNo int)
	GetPokemonTypeEffectiveNess(name string)
}

// internal service
type service struct {
	r Repository
}

// NewService takes in a Repository
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetEggMoves(name string) []EggMove {
	// need to find out a way to get the name of the pokemon while getting the eggmoves
	return s.r.GetEggMoves(name)
}

func (s *service) GetMoveset() {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetPokemon() {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetPokemonTypes(pokdexNo int) {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetPokemonTypeEffectiveNess(name string) {
	panic("not implemented") // TODO: Implement
}
