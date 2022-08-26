package adding

// Service is our adding service interface
type Service interface {
	AddEggMoves(pokdexNo int, moves []EggMove)
}

// Repository is our respository interface
type Repository interface {
	AddEggMove(pokdexNo int, moves []EggMove)
}

type service struct {
	r Repository
}

func (s *service) AddEggMoves(pokdexNo int, moves []EggMove) {
	s.r.AddEggMove(pokdexNo, moves)
}
